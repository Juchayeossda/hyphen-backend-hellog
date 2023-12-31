package controller

import (
	"hyphen-backend-hellog/cerrors/exception"
	"hyphen-backend-hellog/model"
	"hyphen-backend-hellog/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func NewPostController(postSer service.PostService) *PostController {
	return &PostController{postSer}
}

type PostController struct {
	service.PostService
}

func (ctr *PostController) Route(app *fiber.App) {
	app.Post("/api/hellog/posts/post", ctr.create)
	app.Get("/api/hellog/posts/:post_id", ctr.selectByID)
	app.Patch("/api/hellog/posts/:post_id", ctr.update)
	app.Delete("/api/hellog/posts/:post_id", ctr.deleteByID)
}

func (ctr *PostController) create(c *fiber.Ctx) error {
	var clientRequest model.PostCreate

	// title, content parsing
	err := c.BodyParser(&clientRequest)
	exception.Sniff(err)

	// is_private parsing
	clientRequest.IsPrivate = c.FormValue("is_private") == "true"
	exception.Sniff(err)

	// preview_iamge parsing
	clientRequest.PreviewImage, err = c.FormFile("preview_image")
	exception.Sniff(err)

	ctr.PostService.Create(c.Context(), &clientRequest)

	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    fiber.StatusCreated,
		Message: "Success",
		Data:    nil,
	})

}

func (ctr *PostController) selectByID(c *fiber.Ctx) (err error) {
	var clientRequest model.PostSelecByID

	// id parsing
	clientRequest.PostID, err = strconv.Atoi(c.Params("post_id"))
	exception.Sniff(err)

	response := ctr.PostService.SelectByID(c.Context(), &clientRequest)

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success Created",
		Data:    response,
	})
}

func (ctr *PostController) update(c *fiber.Ctx) (err error) {
	var clientRequest model.PostUpdate

	// post id parsing
	clientRequest.PostID, err = strconv.Atoi(c.Params("post_id"))
	exception.Sniff(err)

	// title, content parsing
	err = c.BodyParser(&clientRequest)
	exception.Sniff(err)

	// is_private parsing
	clientRequest.IsPrivate = c.FormValue("is_private") == "true"
	exception.Sniff(err)

	// preview_iamge parsing
	clientRequest.PreviewImage, err = c.FormFile("preview_image")
	exception.Sniff(err)

	ctr.PostService.Update(c.Context(), &clientRequest)

	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    fiber.StatusCreated,
		Message: "Success",
		Data:    nil,
	})
}

func (ctr *PostController) deleteByID(c *fiber.Ctx) (err error) {
	var clientRequest model.PostDeleteByID

	// id parsing
	clientRequest.PostID, err = strconv.Atoi(c.Params("post_id"))
	exception.Sniff(err)

	ctr.PostService.DeleteByID(c.Context(), &clientRequest)

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success Deleted",
		Data:    nil,
	})
}
