package user_handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"gobee/ent"
	"gobee/ent/user"
	"gobee/internal/database"
	"gobee/pkg/domain/model"
)

// @Summary 获取用户列表
// @Description 获取所有用户的列表
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} ent.User
// @Failure 500 {object} model.HttpError
// @Router /api/v1/users [get]
func ListUser(c *fiber.Ctx) error {
	client := database.DB
	users, err := client.User.Query().All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(-1, err.Error()))
	}
	return c.JSON(model.NewSuccess("success", users))
}

// @Summary 创建用户
// @Description 创建一个新用户
// @Tags users
// @Accept json
// @Produce json
// @Param user body ent.User true "用户信息"
// @Success 201 {object} ent.User
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/users [post]
func CreateUser(c *fiber.Ctx) error {
	client := database.DB
	var userData struct {
		Email       string `json:"email"`
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
		Password    string `json:"password"`
	}
	if err := c.BodyParser(&userData); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			err.Error(),
		))
	}

	// 检查邮箱是否已存在
	exists, err := client.User.Query().
		Where(user.EmailEQ(userData.Email)).
		Exist(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	if exists {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Email already exists",
		))
	}

	// 如果提供了手机号，检查手机号是否已存在
	if userData.PhoneNumber != "" {
		exists, err = client.User.Query().
			Where(user.PhoneNumberEQ(userData.PhoneNumber)).
			Exist(c.Context())
		if err != nil {
			return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
		}
		if exists {
			return c.JSON(model.NewError(fiber.StatusBadRequest,
				"Phone number already exists",
			))
		}
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			"Failed to hash password",
		))
	}

	// 创建用户
	newUser, err := client.User.Create().
		SetEmail(userData.Email).
		SetName(userData.Name).
		SetPassword(string(hashedPassword)).
		SetPhoneNumber(userData.PhoneNumber).
		Save(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	return c.JSON(model.NewSuccess("success", newUser))
}

// @Summary 更新用户
// @Description 更新指定用户的信息
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "用户ID"
// @Param user body ent.User true "用户信息"
// @Success 200 {object} ent.User
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/users/{id} [put]
func UpdateUser(c *fiber.Ctx) error {
	client := database.DB
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format",
		))
	}

	var userData struct {
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
		Password    string `json:"password"`
	}
	if err := c.BodyParser(&userData); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			err.Error(),
		))
	}

	// 开始构建更新
	update := client.User.UpdateOneID(id)

	// 如果提供了新名称
	if userData.Name != "" {
		update.SetName(userData.Name)
	}

	// 如果提供了新手机号
	if userData.PhoneNumber != "" {
		// 检查手机号是否已被其他用户使用
		exists, err := client.User.Query().
			Where(
				user.And(
					user.PhoneNumberEQ(userData.PhoneNumber),
					user.IDNEQ(id),
				),
			).
			Exist(c.Context())
		if err != nil {
			return c.JSON(model.NewError(fiber.StatusInternalServerError,
				err.Error(),
			))
		}
		if exists {
			return c.JSON(model.NewError(fiber.StatusBadRequest,
				"Phone number already exists",
			))
		}
		update.SetPhoneNumber(userData.PhoneNumber)
	}

	// 如果提供了新密码
	if userData.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.JSON(model.NewError(fiber.StatusInternalServerError,
				"Failed to hash password",
			))
		}
		update.SetPassword(string(hashedPassword))
	}

	// 执行更新
	updatedUser, err := update.Save(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	return c.JSON(model.NewSuccess("success", updatedUser))
}

// @Summary 查询用户
// @Description 查询指定用户的详细信息
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "用户ID"
// @Success 200 {object} ent.User
// @Failure 400 {object} model.HttpError
// @Failure 404 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/users/{id} [get]
func QueryUser(c *fiber.Ctx) error {
	client := database.DB
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format",
		))
	}

	user, err := client.User.Query().
		Where(user.ID(id)).
		Only(c.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(model.NewError(fiber.StatusNotFound,
				"User not found",
			))
		}
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}

	return c.JSON(model.NewSuccess("success", user))
}

// @Summary 删除用户
// @Description 删除指定用户
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "用户ID"
// @Success 200 {object} model.HttpSuccess
// @Failure 400 {object} model.HttpError
// @Failure 404 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/users/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	client := database.DB
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format",
		))
	}

	err = client.User.DeleteOneID(id).Exec(c.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(model.NewError(fiber.StatusNotFound,
				"User not found",
			))
		}
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}

	return c.JSON(model.NewSuccess("success", nil))
}
