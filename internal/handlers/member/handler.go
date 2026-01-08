package member

import (
	"strconv"

	"gobee/ent"
	member_service "gobee/internal/services/member"
	"gobee/pkg/domain/model"

	"github.com/gofiber/fiber/v2"
)

type MemberHandler interface {
	QueryMember(c *fiber.Ctx) error
	QueryMemberPage(c *fiber.Ctx) error
	CreateMember(c *fiber.Ctx) error
	UpdateMember(c *fiber.Ctx) error
	DeleteMember(c *fiber.Ctx) error
}

type MemberHandlerImpl struct {
	memberService member_service.MemberService
}

func NewMemberHandlerImpl(memberService member_service.MemberService) *MemberHandlerImpl {
	return &MemberHandlerImpl{
		memberService: memberService,
	}
}

func (h *MemberHandlerImpl) QueryMember(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("user_id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid user ID format",
		))
	}

	m, err := h.memberService.QueryMember(c, userId)
	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(model.NewError(fiber.StatusNotFound,
				"Member not found",
			))
		}
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}

	return c.JSON(model.NewSuccess("success", m))
}

func (h *MemberHandlerImpl) QueryMemberPage(c *fiber.Ctx) error {
	pageQuery := model.PageQuery{}
	err := c.QueryParser(&pageQuery)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			err.Error(),
		))
	}

	count, members, err := h.memberService.QueryMemberPage(c, pageQuery)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}

	pageResult := model.PageResult[*ent.Member]{
		Total:   int64(count),
		Records: members,
	}
	return c.JSON(model.NewSuccess("success", pageResult))
}

func (h *MemberHandlerImpl) CreateMember(c *fiber.Ctx) error {
	var createReq *model.MemberCreateReq
	if err := c.BodyParser(&createReq); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			err.Error(),
		))
	}

	m, err := h.memberService.CreateMember(c.Context(), createReq.UserID, createReq.MemberLevel, createReq.MemberNo)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}

	return c.JSON(model.NewSuccess("success", m))
}

func (h *MemberHandlerImpl) UpdateMember(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format",
		))
	}

	var updateReq *model.MemberUpdateReq
	if err = c.BodyParser(&updateReq); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			err.Error(),
		))
	}

	updatedMember, err := h.memberService.UpdateMember(c, id, updateReq)
	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(model.NewError(fiber.StatusNotFound,
				"Member not found",
			))
		}
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}

	return c.JSON(model.NewSuccess("success", updatedMember))
}

func (h *MemberHandlerImpl) DeleteMember(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest,
			"Invalid ID format",
		))
	}

	err = h.memberService.DeleteMember(c, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(model.NewError(fiber.StatusNotFound,
				"Member not found",
			))
		}
		return c.JSON(model.NewError(fiber.StatusInternalServerError,
			err.Error(),
		))
	}

	return c.JSON(model.NewSuccess("success", nil))
}
