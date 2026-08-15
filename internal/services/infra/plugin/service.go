package plugin

import (
	"archive/zip"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/shuTwT/hoshikuzu/ent"
	plugin_ent "github.com/shuTwT/hoshikuzu/ent/plugin"
	plugin_infra "github.com/shuTwT/hoshikuzu/internal/infra/plugin"
	"github.com/shuTwT/hoshikuzu/pkg/domain/model"
	plugin_shared "github.com/shuTwT/hoshikuzu/pkg/plugin/shared"
	"gopkg.in/yaml.v3"
)

type PluginService interface {
	ListPluginPage(ctx context.Context, page, size int) (int, []*ent.Plugin, error)
	ListPluginPageWithQuery(ctx context.Context, req model.PluginPageReq) (int, []*ent.Plugin, error)
	QueryPlugin(ctx context.Context, id int) (*ent.Plugin, error)
	CreatePlugin(ctx context.Context, fileHeader *multipart.FileHeader) (*ent.Plugin, error)
	DeletePlugin(ctx context.Context, id int) error
	StartPlugin(ctx context.Context, id int) error
	StopPlugin(ctx context.Context, id int) error
	RestartPlugin(ctx context.Context, id int) error
	AutoStartPlugins(ctx context.Context) error
	RegisterPlugin(ctx context.Context, pluginInfo *model.PluginRegisterReq) error
	HeartbeatPlugin(ctx context.Context, heartbeatInfo *model.PluginHeartbeatReq) error
	CheckPluginTimeout(ctx context.Context) error
	StartHeartbeatChecker(ctx context.Context)
	GetPluginInstance(ctx context.Context, id int) (plugin_shared.PluginStore, error)
	CallPlugin(ctx context.Context, id int, method string, params []byte) ([]byte, error)
}

// PluginServiceImpl 插件服务：负责插件 CRUD 与上传安装，
// 插件进程的加载、启停与心跳管理委托给基础设施层的 PluginManager
type PluginServiceImpl struct {
	client  *ent.Client
	manager *plugin_infra.PluginManager
}

func NewPluginServiceImpl(client *ent.Client, manager *plugin_infra.PluginManager) *PluginServiceImpl {
	return &PluginServiceImpl{
		client:  client,
		manager: manager,
	}
}

func (s *PluginServiceImpl) ListPluginPage(ctx context.Context, page, size int) (int, []*ent.Plugin, error) {
	count, err := s.client.Plugin.Query().Count(ctx)
	if err != nil {
		return 0, nil, err
	}

	plugins, err := s.client.Plugin.Query().
		Order(ent.Desc(plugin_ent.FieldID)).
		Offset((page - 1) * size).
		Limit(size).
		All(ctx)
	if err != nil {
		return 0, nil, err
	}

	return count, plugins, nil
}

func (s *PluginServiceImpl) ListPluginPageWithQuery(ctx context.Context, req model.PluginPageReq) (int, []*ent.Plugin, error) {
	query := s.client.Plugin.Query()

	if req.Name != "" {
		query.Where(plugin_ent.NameContains(req.Name))
	}

	if req.Key != "" {
		query.Where(plugin_ent.KeyContains(req.Key))
	}

	if req.Status != "" {
		query.Where(plugin_ent.StatusEQ(plugin_ent.Status(req.Status)))
	}

	if req.Enabled != nil {
		query.Where(plugin_ent.Enabled(*req.Enabled))
	}

	if req.AutoStart != nil {
		query.Where(plugin_ent.AutoStart(*req.AutoStart))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return 0, nil, err
	}

	plugins, err := query.
		Order(ent.Desc(plugin_ent.FieldID)).
		Offset((req.Page - 1) * req.Size).
		Limit(req.Size).
		All(ctx)
	if err != nil {
		return 0, nil, err
	}

	return count, plugins, nil
}

func (s *PluginServiceImpl) QueryPlugin(ctx context.Context, id int) (*ent.Plugin, error) {
	pluginEntity, err := s.client.Plugin.Query().
		Where(plugin_ent.ID(id)).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return pluginEntity, nil
}

func (s *PluginServiceImpl) CreatePlugin(ctx context.Context, fileHeader *multipart.FileHeader) (*ent.Plugin, error) {
	if fileHeader == nil {
		return nil, errors.New("文件不能为空")
	}

	srcFile, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("打开上传文件失败: %w", err)
	}
	defer srcFile.Close()

	tempFile, err := os.CreateTemp("", "plugin-*.zip")
	if err != nil {
		return nil, fmt.Errorf("创建临时文件失败: %w", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	if _, err := io.Copy(tempFile, srcFile); err != nil {
		return nil, fmt.Errorf("复制文件失败: %w", err)
	}

	zipReader, err := zip.OpenReader(tempFile.Name())
	if err != nil {
		return nil, fmt.Errorf("打开压缩包失败: %w", err)
	}
	defer zipReader.Close()

	var configContent []byte
	var binaryFile *zip.File
	pluginDir := ""

	for _, f := range zipReader.File {
		if f.Name == "plugin-config.yaml" {
			rc, err := f.Open()
			if err != nil {
				return nil, fmt.Errorf("打开配置文件失败: %w", err)
			}
			configContent, err = io.ReadAll(rc)
			rc.Close()
			if err != nil {
				return nil, fmt.Errorf("读取配置文件失败: %w", err)
			}
		} else if strings.HasSuffix(f.Name, "/") {
			if pluginDir == "" {
				pluginDir = strings.TrimSuffix(f.Name, "/")
			}
		} else {
			ext := filepath.Ext(f.Name)
			if ext == "" || strings.Contains(strings.ToLower(f.Name), "bin") {
				binaryFile = f
			}
		}
	}

	if configContent == nil {
		return nil, errors.New("压缩包中未找到 plugin-config.yaml 文件")
	}

	var pluginConfig model.PluginConfig
	if err := yaml.Unmarshal(configContent, &pluginConfig); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	if err := validatePluginConfig(&pluginConfig); err != nil {
		return nil, err
	}

	exists, err := s.client.Plugin.Query().Where(plugin_ent.Key(pluginConfig.Key)).Exist(ctx)
	if err != nil {
		return nil, fmt.Errorf("检查插件是否存在失败: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("插件 key '%s' 已存在", pluginConfig.Key)
	}

	if len(pluginConfig.Dependencies) > 0 {
		for _, depKey := range pluginConfig.Dependencies {
			depExists, err := s.client.Plugin.Query().Where(plugin_ent.Key(depKey)).Exist(ctx)
			if err != nil {
				return nil, fmt.Errorf("检查依赖插件 '%s' 失败: %w", depKey, err)
			}
			if !depExists {
				return nil, fmt.Errorf("依赖插件 '%s' 不存在", depKey)
			}
		}
	}

	pluginsDir := "./data/plugins"
	if err := os.MkdirAll(pluginsDir, 0755); err != nil {
		return nil, fmt.Errorf("创建插件目录失败: %w", err)
	}

	targetDir := filepath.Join(pluginsDir, pluginConfig.Key)
	if err := os.RemoveAll(targetDir); err != nil {
		return nil, fmt.Errorf("清理旧插件目录失败: %w", err)
	}

	for _, f := range zipReader.File {
		targetPath := filepath.Join(targetDir, f.Name)
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(targetPath, f.Mode()); err != nil {
				return nil, fmt.Errorf("创建目录失败: %w", err)
			}
		} else {
			if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
				return nil, fmt.Errorf("创建父目录失败: %w", err)
			}
			rc, err := f.Open()
			if err != nil {
				return nil, fmt.Errorf("打开压缩文件失败: %w", err)
			}
			defer rc.Close()

			outFile, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return nil, fmt.Errorf("创建文件失败: %w", err)
			}
			defer outFile.Close()

			if _, err := io.Copy(outFile, rc); err != nil {
				return nil, fmt.Errorf("解压文件失败: %w", err)
			}
		}
	}

	var binPath string
	if binaryFile != nil {
		binPath = filepath.Join(targetDir, binaryFile.Name)
		if _, err := os.Stat(binPath); os.IsNotExist(err) {
			return nil, errors.New("压缩包中未找到二进制文件")
		}
		if err := os.Chmod(binPath, 0755); err != nil {
			return nil, fmt.Errorf("设置二进制文件权限失败: %w", err)
		}
	} else {
		return nil, errors.New("压缩包中未找到二进制文件")
	}

	builder := s.client.Plugin.Create().
		SetKey(pluginConfig.Key).
		SetName(pluginConfig.Name).
		SetVersion(pluginConfig.Version).
		SetBinPath(binPath).
		SetMagicCookieValue(pluginConfig.MagicCookieValue).
		SetDependencies(pluginConfig.Dependencies).
		SetEnabled(true).
		SetAutoStart(pluginConfig.AutoStart).
		SetStatus("stopped")

	if pluginConfig.Description != "" {
		builder.SetDescription(pluginConfig.Description)
	}

	if pluginConfig.ProtocolVersion != "" {
		parsedVersion, err := strconv.ParseUint(pluginConfig.ProtocolVersion, 10, 32)
		if err == nil {
			builder.SetProtocolVersion(uint(parsedVersion))
		} else {
			builder.SetProtocolVersion(1)
		}
	} else {
		builder.SetProtocolVersion(1)
	}

	if pluginConfig.MagicCookieKey != "" {
		builder.SetMagicCookieKey(pluginConfig.MagicCookieKey)
	} else {
		builder.SetMagicCookieKey("GO_PLUGIN")
	}

	if pluginConfig.Config != "" {
		builder.SetConfig(pluginConfig.Config)
	}

	pluginEntity, err := builder.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("保存插件信息失败: %w", err)
	}

	return pluginEntity, nil
}

func (s *PluginServiceImpl) DeletePlugin(ctx context.Context, id int) error {
	pluginEntity, err := s.client.Plugin.Query().Where(plugin_ent.ID(id)).First(ctx)
	if err != nil {
		return err
	}

	if pluginEntity.Status == "running" {
		if err := s.manager.Stop(ctx, id); err != nil {
			return fmt.Errorf("停止插件失败: %w", err)
		}
	}

	pluginDir := filepath.Join("./data/plugins", pluginEntity.Key)
	if err := os.RemoveAll(pluginDir); err != nil {
		return fmt.Errorf("删除插件目录失败: %w", err)
	}

	return s.client.Plugin.DeleteOneID(id).Exec(ctx)
}

func (s *PluginServiceImpl) StartPlugin(ctx context.Context, id int) error {
	return s.manager.Start(ctx, id)
}

func (s *PluginServiceImpl) StopPlugin(ctx context.Context, id int) error {
	return s.manager.Stop(ctx, id)
}

func (s *PluginServiceImpl) RestartPlugin(ctx context.Context, id int) error {
	return s.manager.Restart(ctx, id)
}

func (s *PluginServiceImpl) AutoStartPlugins(ctx context.Context) error {
	return s.manager.AutoStart(ctx)
}

func (s *PluginServiceImpl) RegisterPlugin(ctx context.Context, pluginInfo *model.PluginRegisterReq) error {
	return s.manager.Register(ctx, pluginInfo)
}

func (s *PluginServiceImpl) HeartbeatPlugin(ctx context.Context, heartbeatInfo *model.PluginHeartbeatReq) error {
	return s.manager.Heartbeat(ctx, heartbeatInfo)
}

func (s *PluginServiceImpl) CheckPluginTimeout(ctx context.Context) error {
	return s.manager.CheckTimeout(ctx)
}

func (s *PluginServiceImpl) StartHeartbeatChecker(ctx context.Context) {
	s.manager.StartHeartbeatChecker(ctx)
}

func (s *PluginServiceImpl) GetPluginInstance(ctx context.Context, id int) (plugin_shared.PluginStore, error) {
	return s.manager.GetInstance(ctx, id)
}

func (s *PluginServiceImpl) CallPlugin(ctx context.Context, id int, method string, params []byte) ([]byte, error) {
	return s.manager.Call(ctx, id, method, params)
}

func validatePluginConfig(config *model.PluginConfig) error {
	if config.Key == "" {
		return errors.New("插件 key 不能为空")
	}
	if config.Name == "" {
		return errors.New("插件名称不能为空")
	}
	if config.Version == "" {
		return errors.New("插件版本不能为空")
	}
	if config.MagicCookieValue == "" {
		return errors.New("Magic Cookie Value 不能为空")
	}
	return nil
}
