package plugin

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"sync"
	"time"

	go_plugin "github.com/hashicorp/go-plugin"
	"github.com/shuTwT/hoshikuzu/ent"
	plugin_ent "github.com/shuTwT/hoshikuzu/ent/plugin"
	"github.com/shuTwT/hoshikuzu/pkg/domain/model"
	capability_plugin "github.com/shuTwT/hoshikuzu/pkg/plugin/capability"
	plugin_shared "github.com/shuTwT/hoshikuzu/pkg/plugin/shared"
)

// go-plugin 分发键，插件侧 Serve 注册时必须使用相同的键
const dispatchKey = "plugin_store"

// 心跳超时阈值，同时也是超时检查的轮询周期
const heartbeatTimeout = 30 * time.Second

// PluginManager 插件管理器（基础设施层）：负责插件的加载、运行实例管理、生命周期与心跳管理。
// 服务层只负责 CRUD，通过本管理器完成插件进程的启停与协议通信。
type PluginManager struct {
	entClient      *ent.Client
	pluginMap      map[string]go_plugin.Plugin // 宿主与插件共享的 gRPC 协议插件映射
	runningPlugins map[int]*go_plugin.Client   // 运行中的插件客户端（key：插件ID）
	registry       map[string]map[string]any   // 插件注册信息（key：插件名，由插件主动上报）
	heartbeats     map[string]time.Time        // 插件心跳时间（key：插件名）
	mu             sync.RWMutex
}

// NewPluginManager 创建插件管理器实例
func NewPluginManager(entClient *ent.Client) *PluginManager {
	return &PluginManager{
		entClient: entClient,
		// 宿主内置的通用插件协议：
		// plugin_store - 生命周期协议（进程管理/健康检查/静态资源）
		// capability  - 通用能力协议（业务方法通过 Call 动态调用，新增插件无需改宿主代码）
		pluginMap: map[string]go_plugin.Plugin{
			dispatchKey:  &plugin_shared.StoreGRPCPlugin{},
			"capability": &capability_plugin.CapabilityPlugin{},
		},
		runningPlugins: make(map[int]*go_plugin.Client),
		registry:       make(map[string]map[string]any),
		heartbeats:     make(map[string]time.Time),
	}
}

// Start 启动指定ID的插件：校验启用状态与依赖，状态置 loading 后异步拉起进程并完成 gRPC 握手
func (m *PluginManager) Start(ctx context.Context, id int) error {
	pluginEntity, err := m.entClient.Plugin.Query().Where(plugin_ent.ID(id)).First(ctx)
	if err != nil {
		return err
	}

	if !pluginEntity.Enabled {
		return errors.New("插件未启用")
	}

	if _, err := os.Stat(pluginEntity.BinPath); os.IsNotExist(err) {
		return fmt.Errorf("插件二进制文件不存在: %s", pluginEntity.BinPath)
	}

	if len(pluginEntity.Dependencies) > 0 {
		for _, depKey := range pluginEntity.Dependencies {
			depPlugin, err := m.entClient.Plugin.Query().Where(plugin_ent.Key(depKey)).First(ctx)
			if err != nil {
				return fmt.Errorf("获取依赖插件 '%s' 失败: %w", depKey, err)
			}
			if depPlugin.Status != "running" {
				return fmt.Errorf("依赖插件 '%s' 未运行", depKey)
			}
		}
	}

	m.mu.Lock()
	if client, exists := m.runningPlugins[id]; exists {
		if !client.Exited() {
			m.mu.Unlock()
			return errors.New("插件已在运行中")
		}
		client.Kill()
		delete(m.runningPlugins, id)
	}
	m.mu.Unlock()

	now := time.Now()
	if err := m.entClient.Plugin.UpdateOneID(id).
		SetStatus("loading").
		SetLastStartedAt(now).
		SetLastError("").
		Exec(ctx); err != nil {
		return err
	}

	go m.startProcess(context.Background(), pluginEntity)
	return nil
}

// startProcess 异步拉起插件进程：握手、分发协议实例、注入配置
func (m *PluginManager) startProcess(ctx context.Context, pluginEntity *ent.Plugin) {
	client := go_plugin.NewClient(&go_plugin.ClientConfig{
		HandshakeConfig: go_plugin.HandshakeConfig{
			ProtocolVersion:  pluginEntity.ProtocolVersion,
			MagicCookieKey:   pluginEntity.MagicCookieKey,
			MagicCookieValue: pluginEntity.MagicCookieValue,
		},
		Plugins:          m.pluginMap,
		Cmd:              exec.Command(pluginEntity.BinPath),
		AllowedProtocols: []go_plugin.Protocol{go_plugin.ProtocolGRPC},
		Managed:          true,
	})

	rpcClient, err := client.Client()
	if err != nil {
		m.failStart(pluginEntity.ID, fmt.Sprintf("插件连接失败: %s", err.Error()))
		client.Kill()
		return
	}

	raw, err := rpcClient.Dispense(dispatchKey)
	if err != nil {
		m.failStart(pluginEntity.ID, fmt.Sprintf("插件分发失败: %s", err.Error()))
		client.Kill()
		return
	}

	instance, ok := raw.(plugin_shared.PluginStore)
	if !ok {
		m.failStart(pluginEntity.ID, "插件实例类型错误")
		client.Kill()
		return
	}

	if err := instance.Init(ctx, map[string]string{
		"plugin_name": pluginEntity.Name,
		"version":     pluginEntity.Version,
	}); err != nil {
		m.failStart(pluginEntity.ID, fmt.Sprintf("插件初始化失败: %s", err.Error()))
		client.Kill()
		return
	}

	m.mu.Lock()
	m.runningPlugins[pluginEntity.ID] = client
	m.mu.Unlock()

	if err := m.entClient.Plugin.UpdateOneID(pluginEntity.ID).
		SetStatus("running").
		Exec(ctx); err != nil {
		slog.Error("更新插件运行状态失败", "plugin_id", pluginEntity.ID, "error", err.Error())
	}
}

// failStart 记录插件启动失败状态
func (m *PluginManager) failStart(id int, message string) {
	slog.Error("插件启动失败", "plugin_id", id, "error", message)
	if err := m.entClient.Plugin.UpdateOneID(id).
		SetStatus("error").
		SetLastError(message).
		Exec(context.Background()); err != nil {
		slog.Error("更新插件错误状态失败", "plugin_id", id, "error", err.Error())
	}
}

// Stop 停止指定ID的插件：调用插件销毁方法并终止进程
func (m *PluginManager) Stop(ctx context.Context, id int) error {
	m.mu.RLock()
	client, exists := m.runningPlugins[id]
	m.mu.RUnlock()

	if !exists {
		return errors.New("插件未运行")
	}

	if instance, err := m.dispense(client, dispatchKey); err == nil {
		if lifecycleInstance, ok := instance.(plugin_shared.PluginStore); ok {
			_ = lifecycleInstance.Destroy(ctx)
		}
	}

	client.Kill()

	m.mu.Lock()
	delete(m.runningPlugins, id)
	m.mu.Unlock()

	now := time.Now()
	return m.entClient.Plugin.UpdateOneID(id).
		SetStatus("stopped").
		SetLastStoppedAt(now).
		Exec(ctx)
}

// Restart 重启指定ID的插件
func (m *PluginManager) Restart(ctx context.Context, id int) error {
	pluginEntity, err := m.entClient.Plugin.Query().Where(plugin_ent.ID(id)).First(ctx)
	if err != nil {
		return err
	}

	if pluginEntity.Status == "running" {
		if err := m.Stop(ctx, id); err != nil {
			return fmt.Errorf("停止插件失败: %w", err)
		}
		time.Sleep(1 * time.Second)
	}

	return m.Start(ctx, id)
}

// AutoStart 自动启动所有启用且标记了自动启动的插件
func (m *PluginManager) AutoStart(ctx context.Context) error {
	plugins, err := m.entClient.Plugin.Query().
		Where(plugin_ent.AutoStart(true)).
		Where(plugin_ent.Enabled(true)).
		All(ctx)
	if err != nil {
		return fmt.Errorf("获取自动启动插件失败: %w", err)
	}

	for _, p := range plugins {
		if p.Status != "running" {
			go func(pluginID int) {
				if err := m.Start(context.Background(), pluginID); err != nil {
					slog.Error("自动启动插件失败", "plugin_id", pluginID, "error", err.Error())
				}
			}(p.ID)
		}
	}

	return nil
}

// GetInstance 获取运行中插件的生命周期实例（PluginStore 协议）
func (m *PluginManager) GetInstance(ctx context.Context, id int) (plugin_shared.PluginStore, error) {
	raw, err := m.Dispense(ctx, id, dispatchKey)
	if err != nil {
		return nil, err
	}
	instance, ok := raw.(plugin_shared.PluginStore)
	if !ok {
		return nil, errors.New("插件实例类型错误")
	}
	return instance, nil
}

// GetInstanceByKey 按插件 key 查找运行中的插件，并分发指定分发键的协议实例（类型由调用方断言）
func (m *PluginManager) GetInstanceByKey(ctx context.Context, key, dispatchKey string) (interface{}, error) {
	pluginEntity, err := m.entClient.Plugin.Query().Where(plugin_ent.Key(key)).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("插件 '%s' 不存在: %w", key, err)
	}
	if pluginEntity.Status != "running" {
		return nil, fmt.Errorf("插件 '%s' 未运行", key)
	}
	return m.Dispense(ctx, pluginEntity.ID, dispatchKey)
}

// Dispense 获取运行中插件指定分发键的协议实例，类型由调用方断言
func (m *PluginManager) Dispense(ctx context.Context, id int, dispatchKey string) (interface{}, error) {
	m.mu.RLock()
	client, exists := m.runningPlugins[id]
	m.mu.RUnlock()

	if !exists {
		return nil, errors.New("插件未运行")
	}
	return m.dispense(client, dispatchKey)
}

// dispense 从已连接的 go-plugin 客户端分发协议实例
func (m *PluginManager) dispense(client *go_plugin.Client, dispatchKey string) (interface{}, error) {
	if client.Exited() {
		return nil, errors.New("插件进程已退出")
	}

	rpcClient, err := client.Client()
	if err != nil {
		return nil, err
	}

	return rpcClient.Dispense(dispatchKey)
}

// Call 调用运行中插件的通用能力方法（capability 协议）。
// 宿主不感知插件具体业务，method 与 JSON 参数由插件自行定义与解析。
func (m *PluginManager) Call(ctx context.Context, id int, method string, params []byte) ([]byte, error) {
	raw, err := m.Dispense(ctx, id, "capability")
	if err != nil {
		return nil, err
	}
	instance, ok := raw.(capability_plugin.Capability)
	if !ok {
		return nil, errors.New("插件不支持通用能力协议")
	}
	return instance.Call(ctx, method, params)
}

// Register 记录插件注册信息（插件进程启动后主动上报）
func (m *PluginManager) Register(ctx context.Context, info *model.PluginRegisterReq) error {
	pluginKey := info.Name

	m.mu.Lock()
	m.registry[pluginKey] = map[string]any{
		"name":         info.Name,
		"version":      info.Version,
		"grpc_address": info.GrpcAddress,
		"status":       info.Status,
		"start_time":   info.StartTime,
		"metadata":     info.Metadata,
	}
	m.heartbeats[pluginKey] = time.Now()
	m.mu.Unlock()

	return nil
}

// Heartbeat 更新插件心跳时间
func (m *PluginManager) Heartbeat(ctx context.Context, info *model.PluginHeartbeatReq) error {
	pluginKey := info.Name

	m.mu.Lock()
	m.heartbeats[pluginKey] = time.Now()
	if info.Status != "" {
		if pluginInfo, exists := m.registry[pluginKey]; exists {
			pluginInfo["status"] = info.Status
			m.registry[pluginKey] = pluginInfo
		}
	}
	m.mu.Unlock()

	return nil
}

// CheckTimeout 检查插件心跳是否超时，超时则标记为已停止
func (m *PluginManager) CheckTimeout(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	now := time.Now()
	for pluginKey, heartbeatTime := range m.heartbeats {
		if now.Sub(heartbeatTime) > heartbeatTimeout {
			if pluginInfo, exists := m.registry[pluginKey]; exists {
				pluginInfo["status"] = "stopped"
				m.registry[pluginKey] = pluginInfo
			}
		}
	}

	return nil
}

// StartHeartbeatChecker 启动心跳超时检查协程（30秒周期），随宿主进程常驻
func (m *PluginManager) StartHeartbeatChecker(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(heartbeatTimeout)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if err := m.CheckTimeout(ctx); err != nil {
					slog.Error("插件心跳超时检查失败", "error", err.Error())
				}
			}
		}
	}()
}
