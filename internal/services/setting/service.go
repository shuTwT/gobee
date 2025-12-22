package setting

import (
	"context"
	"gobee/ent"
	"gobee/ent/setting"
)

type SettingService interface {
	GetAllSettings(ctx context.Context, client *ent.Client) ([]*ent.Setting, error)
	ExistSettingByKey(ctx context.Context, client *ent.Client, key string) (bool, error)
	GetSettingByKey(ctx context.Context, client *ent.Client, key string) (*ent.Setting, error)
	UpdateSettingByKey(ctx context.Context, client *ent.Client, key string, value string) error
	CreateSettingIfNotExist(ctx context.Context, client *ent.Client, key string, value string) error
	IsSystemInitialized(ctx context.Context, client *ent.Client) (bool, error)
	SetSystemInitialized(ctx context.Context, client *ent.Client) error
}

type SettingServiceImpl struct {
	client *ent.Client
}

func NewSettingServiceImpl(client *ent.Client) *SettingServiceImpl {
	return &SettingServiceImpl{client: client}
}

// GetAllSettings 获取所有系统设置
func (s *SettingServiceImpl) GetAllSettings(ctx context.Context, client *ent.Client) ([]*ent.Setting, error) {
	return client.Setting.Query().All(ctx)
}

func (s *SettingServiceImpl) ExistSettingByKey(ctx context.Context, client *ent.Client, key string) (bool, error) {
	return client.Setting.Query().Where(setting.KeyEQ(key)).Exist(ctx)
}

// GetSettingByKey 根据键获取设置
func (s *SettingServiceImpl) GetSettingByKey(ctx context.Context, client *ent.Client, key string) (*ent.Setting, error) {
	return client.Setting.Query().Where(setting.KeyEQ(key)).Only(ctx)
}

func (s *SettingServiceImpl) UpdateSettingByKey(ctx context.Context, client *ent.Client, key string, value string) error {
	_, err := client.Setting.Update().Where(setting.KeyEQ(key)).SetValue(value).Save(ctx)
	return err
}

func (s *SettingServiceImpl) CreateSettingIfNotExist(ctx context.Context, client *ent.Client, key string, value string) error {

	_, err := client.Setting.Create().SetKey(key).SetValue(value).Save(ctx)

	return err
}

// IsSystemInitialized 检查系统是否已初始化
// 通过检查是否存在特定的初始化标记设置来判断
func (s *SettingServiceImpl) IsSystemInitialized(ctx context.Context, client *ent.Client) (bool, error) {
	exists, err := client.Setting.Query().
		Where(setting.KeyEQ("system_initialized")).
		Exist(ctx)

	return exists, err
}

// SetSystemInitialized 标记系统已初始化
func (s *SettingServiceImpl) SetSystemInitialized(ctx context.Context, client *ent.Client) error {
	// 检查是否存在 system_initialized 设置
	exists, err := client.Setting.Query().Where(setting.KeyEQ("system_initialized")).Exist(ctx)
	if err != nil {
		return err
	}

	if exists {
		// 如果存在，则更新其值为 true
		_, err = client.Setting.Update().Where(setting.KeyEQ("system_initialized")).SetValue("true").Save(ctx)
	} else {
		// 如果不存在，则创建新的设置
		_, err = client.Setting.Create().SetKey("system_initialized").SetValue("true").Save(ctx)
	}
	return err
}
