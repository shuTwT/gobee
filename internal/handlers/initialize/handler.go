package initialize

import (
	"fmt"
	"gobee/internal/database"
	"gobee/internal/services"
	"gobee/internal/services/setting"
	"gobee/pkg/domain/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// InitializeRequest 定义了初始化请求的结构
type InitializeRequest struct {
	DBType          string `json:"dbType" validate:"required,oneof=sqlite mysql postgres"`
	DBHost          string `json:"dbHost"`
	DBPort          int    `json:"dbPort"`
	DBUser          string `json:"dbUser"`
	DBPassword      string `json:"dbPassword"`
	DBName          string `json:"dbName"`
	AdminUsername   string `json:"adminUsername" validate:"required"`
	AdminPassword   string `json:"adminPassword" validate:"required,min=6"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=6"`
	AdminEmail      string `json:"adminEmail" validate:"required,email"`
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
func Initialize(c *fiber.Ctx) error {
	var req InitializeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}

	if req.AdminPassword != req.ConfirmPassword {
		return c.JSON(model.NewError(fiber.StatusBadRequest, "两次输入的密码不一致"))
	}

	// 检查系统是否已初始化
	isInitialized, err := setting.IsSystemInitialized(c.UserContext(), database.DB)
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

	// 创建超级管理员账户
	adminUser, err := services.CreateUser(services.CreateUserRequest{
		Username: req.AdminUsername,
		Password: req.AdminPassword,
		Email:    req.AdminEmail,
	})
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, fmt.Sprintf("创建超级管理员账户失败: %v", err)))
	}
	fmt.Printf("超级管理员账户创建成功: %s\n", adminUser.Name)

	// 标记系统已初始化
	if err := setting.SetSystemInitialized(c.UserContext(), database.DB); err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, fmt.Sprintf("标记系统初始化状态失败: %v", err)))
	}

	return c.JSON(model.NewSuccess("系统初始化成功", nil))
}
