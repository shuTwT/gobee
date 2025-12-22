package initialize

import (
	"fmt"
	"gobee/ent"
	"gobee/internal/database"
	"gobee/internal/services/setting"
	services "gobee/internal/services/user"
	"gobee/pkg/domain/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type InitializeHandler interface {
	Initialize(c *fiber.Ctx) error
}

type InitializeHandlerImpl struct {
	userService    services.UserService
	settingService setting.SettingService
}

func NewInitializeHandlerImpl(userService services.UserService, settingService setting.SettingService) *InitializeHandlerImpl {
	return &InitializeHandlerImpl{userService: userService, settingService: settingService}
}

// Initialize 处理系统初始化请求
// @Summary 系统初始化
// @Description 首次运行系统时进行初始化设置，包括数据库配置和创建管理员账户
// @Tags Initialize
// @Accept json
// @Produce json
// @Param request body InitializeRequest true "初始化请求"
// @Success 200 {object} map[string]string "初始化成功"
// @Failure 400 {object} model.ErrorResponse "请求参数错误"
// @Failure 409 {object} model.ErrorResponse "系统已初始化"
// @Failure 500 {object} model.ErrorResponse "服务器内部错误"
// @Router /initialize [post]
func (h *InitializeHandlerImpl) Initialize(c *fiber.Ctx) error {
	var req *model.InitializeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}

	if req.AdminPassword != req.ConfirmPassword {
		return c.JSON(model.NewError(fiber.StatusBadRequest, "两次输入的密码不一致"))
	}

	// 检查系统是否已初始化
	isInitialized, err := h.settingService.IsSystemInitialized(c.UserContext(), database.DB)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, fmt.Sprintf("检查系统初始化状态失败: %v", err)))
	}
	if isInitialized {
		return c.JSON(model.NewError(fiber.StatusConflict, "系统已初始化，请勿重复操作"))
	}

	log.Info(fmt.Sprintf("初始化请求参数: %s", req.DBType))
	// 根据数据库类型进行参数验证
	switch req.DBType {
	case "mysql", "postgres":
		if req.DBHost == "" || req.DBPort == 0 || req.DBUser == "" || req.DBPassword == "" || req.DBName == "" {
			return c.JSON(model.NewError(fiber.StatusBadRequest, "MySQL/PostgreSQL数据库连接信息不完整"))
		}
	case "sqlite":
		// SQLite不需要额外的连接信息，但可以接受DBName作为文件路径
	default:
		return c.JSON(model.NewError(fiber.StatusBadRequest, "不支持的数据库类型"))
	}

	// 数据库初始化
	dbConfig := database.DBConfig{
		DBType: req.DBType,
	}

	switch req.DBType {
	case "mysql":
		dbConfig.MysqlHost = req.DBHost
		dbConfig.MysqlPort = fmt.Sprintf("%d", req.DBPort)
		dbConfig.MysqlUser = req.DBUser
		dbConfig.MysqlPassword = req.DBPassword
		dbConfig.MysqlDatabase = req.DBName
	case "postgres":
		dbConfig.PgHost = req.DBHost
		dbConfig.PgPort = fmt.Sprintf("%d", req.DBPort)
		dbConfig.PgUser = req.DBUser
		dbConfig.PgPassword = req.DBPassword
		dbConfig.PgDatabase = req.DBName
	case "sqlite":
		dbConfig.SqliteFile = req.DBName
	}

	// _, err = database.InitializeDB(dbConfig, true)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(model.NewError(fiber.StatusInternalServerError, fmt.Sprintf("数据库初始化失败: %v", err)))
	// }

	// 创建角色e
	initRole(c)
	// 创建超级管理员账户
	adminUser, err := h.userService.CreateUser(services.CreateUserRequest{
		Username: req.AdminUsername,
		Password: req.AdminPassword,
		Email:    req.AdminEmail,
		RoleId:   1,
	})
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, fmt.Sprintf("创建超级管理员账户失败: %v", err)))
	}
	fmt.Printf("超级管理员账户创建成功: %s\n", adminUser.Name)

	// 标记系统已初始化
	if err := h.settingService.SetSystemInitialized(c.UserContext(), database.DB); err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, fmt.Sprintf("标记系统初始化状态失败: %v", err)))
	}

	return c.JSON(model.NewSuccess("系统初始化成功", nil))
}

func initRole(c *fiber.Ctx) *ent.Role {
	client := database.DB
	// 初始化角色
	role := client.Role.Create().
		SetID(1).
		SetName("超级管理员").
		SetCode("superAdmin").
		SetIsDefault(true).
		SaveX(c.UserContext())
	client.Role.Create().
		SetID(2).
		SetName("访客").
		SetCode("guest").
		SetIsDefault(true).
		SaveX(c.UserContext())
	client.Role.Create().
		SetID(3).
		SetName("普通用户").
		SetCode("common").
		SetIsDefault(true).
		SaveX(c.UserContext())

	return role
}
