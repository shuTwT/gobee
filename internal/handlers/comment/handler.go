package comment

import (
	commentService "gobee/internal/services/comment"

	"github.com/gofiber/fiber/v2"
)

func ListCommentPage(c *fiber.Ctx) error {
	return commentService.ListCommentPage(c)
}
