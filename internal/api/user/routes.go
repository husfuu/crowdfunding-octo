package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/husfuu/crowdfunding-octo/internal/wrapper/handler"
)

func NewRoutes(api fiber.Router, handler handler.Handler) {
	api.Get("users/:id", handler.Api.User.GetUser)
	api.Post("users/email_checkers", handler.Api.User.CheckEmailAvailability)
	api.Post("login", handler.Api.User.Login)
	api.Post("users", handler.Api.User.Register)
	api.Post("users/avatars", handler.Api.User.UploadAvatar)
	api.Put("users", handler.Api.User.UpdateUser)
}
