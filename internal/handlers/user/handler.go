package user_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"gobee/ent"
	"gobee/ent/user"
)

// ListUser 获取用户列表
func ListUser(c *fiber.Ctx) error {
	client := c.Locals("client").(*ent.Client)
	users, err := client.User.Query().All(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": users,
	})
}

// CreateUser 创建用户
func CreateUser(c *fiber.Ctx) error {
	client := c.Locals("client").(*ent.Client)
	var userData struct {
		Email       string `json:"email"`
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
		Password    string `json:"password"`
	}
	if err := c.BodyParser(&userData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// 检查邮箱是否已存在
	exists, err := client.User.Query().
		Where(user.EmailEQ(userData.Email)).
		Exist(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if exists {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email already exists",
		})
	}

	// 如果提供了手机号，检查手机号是否已存在
	if userData.PhoneNumber != "" {
		exists, err = client.User.Query().
			Where(user.PhoneNumberEQ(userData.PhoneNumber)).
			Exist(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if exists {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Phone number already exists",
			})
		}
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	// 创建用户
	newUser, err := client.User.Create().
		SetEmail(userData.Email).
		SetName(userData.Name).
		SetPassword(string(hashedPassword)).
		SetPhoneNumber(userData.PhoneNumber).
		Save(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": newUser,
	})
}

// UpdateUser 更新用户
func UpdateUser(c *fiber.Ctx) error {
	client := c.Locals("client").(*ent.Client)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	var userData struct {
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
		Password    string `json:"password"`
	}
	if err := c.BodyParser(&userData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
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
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if exists {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Phone number already exists",
			})
		}
		update.SetPhoneNumber(userData.PhoneNumber)
	}

	// 如果提供了新密码
	if userData.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to hash password",
			})
		}
		update.SetPassword(string(hashedPassword))
	}

	// 执行更新
	updatedUser, err := update.Save(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": updatedUser,
	})
}

// QueryUser 查询用户
func QueryUser(c *fiber.Ctx) error {
	client := c.Locals("client").(*ent.Client)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	user, err := client.User.Query().
		Where(user.ID(id)).
		Only(c.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": user,
	})
}

// DeleteUser 删除用户
func DeleteUser(c *fiber.Ctx) error {
	client := c.Locals("client").(*ent.Client)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	err = client.User.DeleteOneID(id).Exec(c.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
