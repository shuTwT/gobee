package user_handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"gobee/ent"
	"gobee/ent/personalaccesstoken"
	"gobee/ent/user"
	"gobee/internal/database"
	role_service "gobee/internal/services/role"
	user_service "gobee/internal/services/user"
	"gobee/pkg/domain/model"
)

type UserHandler interface {
	ListUser(c *fiber.Ctx) error
	ListUserPage(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	QueryUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	GetPersonalAccessTokenList(c *fiber.Ctx) error
	GetPersonalAccessToken(c *fiber.Ctx) error
	CreatePat(c *fiber.Ctx) error
	GetUserProfile(c *fiber.Ctx) error
}

type UserHandlerImpl struct {
	userService user_service.UserService
	roleService role_service.RoleService
}

func NewUserHandlerImpl(userService user_service.UserService, roleService role_service.RoleService) *UserHandlerImpl {
	return &UserHandlerImpl{
		userService: userService,
		roleService: roleService,
	}
}

// @Summary 获取用户列表
// @Description 获取所有用户的列表
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} ent.User
// @Failure 500 {object} model.HttpError
// @Router /api/v1/users [get]
func (h *UserHandlerImpl) ListUser(c *fiber.Ctx) error {
	users, err := h.userService.ListUser(c)
	if err != nil {
		return c.JSON(model.NewError(-1, err.Error()))
	}
	userRespList := []model.UserResp{}
	for _, user := range users {
		userRespList = append(userRespList, model.UserResp{
			ID:     user.ID,
			Name:   user.Name,
			Email:  user.Email,
			RoleID: &user.RoleID,
		})
	}
	return c.JSON(model.NewSuccess("success", userRespList))
}

func (h *UserHandlerImpl) ListUserPage(c *fiber.Ctx) error {
	pageQuery := model.PageQuery{}
	err := c.QueryParser(&pageQuery)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			err.Error(),
		))
	}

	count, users, err := h.userService.ListUserPage(c, pageQuery)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}
	userRespList := []model.UserResp{}
	for _, user := range users {
		userResp := model.UserResp{
			ID:     user.ID,
			Name:   user.Name,
			Email:  user.Email,
			RoleID: &user.RoleID,
		}
		role, _ := h.roleService.QueryRole(c, user.RoleID)
		userResp.Role = role
		userRespList = append(userRespList, userResp)
	}
	pageResult := model.PageResult[model.UserResp]{
		Total:   int64(count),
		Records: userRespList,
	}
	return c.JSON(model.NewSuccess("success", pageResult))
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
func (h *UserHandlerImpl) CreateUser(c *fiber.Ctx) error {
	client := database.DB
	var userData *model.UserCreateReq
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
func (h *UserHandlerImpl) UpdateUser(c *fiber.Ctx) error {
	client := database.DB
	var err error
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format",
		))
	}

	var userData *model.UserUpdateReq
	if err = c.BodyParser(&userData); err != nil {
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
		var exists bool
		exists, err = client.User.Query().
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
		var hashedPassword []byte
		hashedPassword, err = bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
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
func (h *UserHandlerImpl) QueryUser(c *fiber.Ctx) error {
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
func (h *UserHandlerImpl) DeleteUser(c *fiber.Ctx) error {
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

func (h *UserHandlerImpl) GetPersonalAccessTokenList(c *fiber.Ctx) error {
	userId := int(c.Locals("userId").(float64))
	client := database.DB
	tokens, err := client.PersonalAccessToken.Query().Where(personalaccesstoken.UserID(userId)).All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(
			fiber.StatusBadRequest, err.Error(),
		))
	}
	result := []model.PersonalAccessTokenListResp{}

	for _, token := range tokens {
		result = append(result, model.PersonalAccessTokenListResp{
			ID:          token.ID,
			Name:        token.Name,
			Expires:     model.ParseTime(token.Expires),
			Description: token.Name,
		})
	}

	return c.JSON(model.NewSuccess("success", result))
}

// @Summary 查询个人令牌
// @Description 查询指定个人令牌的详细信息
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "个人令牌ID"
// @Success 200 {object} model.PersonalAccessTokenResp
// @Failure 400 {object} model.HttpError
// @Failure 404 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/users/personal-access-tokens/{id} [get]
func (h *UserHandlerImpl) GetPersonalAccessToken(c *fiber.Ctx) error {
	userId := int(c.Locals("userId").(float64))
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(
			fiber.StatusBadRequest, err.Error(),
		))
	}
	client := database.DB
	token, err := client.PersonalAccessToken.Query().Where(personalaccesstoken.UserIDEQ(userId), personalaccesstoken.IDEQ(id)).Only(c.Context())
	if err != nil {
		return c.JSON(model.NewError(
			fiber.StatusBadRequest, err.Error(),
		))
	}
	result := model.PersonalAccessTokenResp{
		ID:          token.ID,
		Name:        token.Name,
		Expires:     model.ParseTime(token.Expires),
		Description: token.Description,
		Token:       token.Token,
	}
	return c.JSON(model.NewSuccess("success", result))
}

// @Summary 创建 personalAccessToken 个人令牌
// @Description 创建一个新的个人令牌
// @Tags users
// @Accept json
// @Produce json
// @Param createReq body model.PersonalAccessTokenCreateReq true "个人令牌创建请求"
// @Success 200 {object} model.HttpSuccess
// @Failure 400 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/users/personal-access-tokens [post]
func (h *UserHandlerImpl) CreatePat(c *fiber.Ctx) error {
	userId := int(c.Locals("userId").(float64))
	var createReq *model.PersonalAccessTokenCreateReq
	err := c.BodyParser(&createReq)
	if err != nil {
		return c.JSON(model.NewError(
			fiber.StatusBadRequest, err.Error(),
		))
	}
	client := database.DB
	// 查找用户
	u, _ := client.User.Query().
		Where(user.IDEQ(userId)).
		Only(c.Context())

		// 生成JWT令牌
	claims := jwt.MapClaims{
		"id":    u.ID,
		"email": u.Email,
		"name":  u.Name,
		"exp":   createReq.Expires, // 24小时过期
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名令牌
	t, err := token.SignedString([]byte("your-secret-key")) // 注意：在生产环境中应该使用环境变量存储密钥
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.NewError(
			fiber.StatusBadRequest, "Could not generate token",
		))
	}

	client.PersonalAccessToken.Create().
		SetName(createReq.Name).
		SetDescription(createReq.Description).
		SetExpires(createReq.Expires.Time()).
		SetToken(t).
		SetUserID(userId).Save(c.Context())

	return c.JSON(model.NewSuccess("success", nil))
}

// @Summary 查询用户个人信息
// @Description 查询指定用户的个人信息
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "用户ID"
// @Success 200 {object} model.UserProfileResp
// @Failure 400 {object} model.HttpError
// @Failure 404 {object} model.HttpError
// @Failure 500 {object} model.HttpError
// @Router /api/v1/users/{id}/profile [get]
func (h *UserHandlerImpl) GetUserProfile(c *fiber.Ctx) error {
	userId := int(c.Locals("userId").(float64))
	client := database.DB

	user, err := client.User.Query().
		Where(user.ID(userId)).
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
	role, err := client.User.QueryRole(user).Only(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}

	result := model.UserProfileResp{
		UserID:              userId,
		Email:               user.Email,
		EmailVerified:       user.EmailVerified,
		Name:                user.Name,
		PhoneNumber:         &user.PhoneNumber,
		PhoneNumberVerified: user.PhoneNumberVerified,
		Role:                role,
	}

	return c.JSON(model.NewSuccess("success", result))
}
