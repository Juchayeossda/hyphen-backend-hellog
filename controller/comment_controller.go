package controller

import (
	"hyphen-backend-hellog/cerrors/exception"
	"hyphen-backend-hellog/model"
	"hyphen-backend-hellog/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func NewCommentController(commentServ service.CommentService) *CommentController {
	return &CommentController{commentServ}
}

type CommentController struct {
	service.CommentService
}

func (ctr *CommentController) Route(app *fiber.App) {
	app.Post("/api/hellog/posts/:post_id/comments/comment", ctr.create)
	// app.Get("/api/hellog/comments/:id", ctr.selectByID)
}

func (ctr *CommentController) create(c *fiber.Ctx) error {
	var clientRequest model.CommentCreate

	//content, PaerntID parsing
	err := c.BodyParser(&clientRequest)
	exception.Sniff(err)

	// is_private parsing
	clientRequest.PostID, err = strconv.Atoi(c.Params("post_id"))
	exception.Sniff(err)

	ctr.CommentService.Create(c.Context(), &clientRequest)

	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    fiber.StatusCreated,
		Message: "Success",
		Data:    nil,
	})
}
