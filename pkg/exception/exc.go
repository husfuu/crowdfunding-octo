package exception

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/husfuu/crowdfunding-octo/config"
	"github.com/sirupsen/logrus"
)

type ResponseData struct {
	// RequestId string             `json:"request_id"`
	Data      interface{}        `json:"data"`
	Status    StatusResponseData `json:"status"`
	TimeStamp time.Time          `json:"timeStamp"`
	Page      interface{}        `json:"page,omitempty"`
}

type StatusResponseData struct {
	Code    int    `json:"code"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

type InitialExceptionCreateResponse struct {
	Ctx  *fiber.Ctx
	Conf *config.Config
	Log  *logrus.Logger
}

func InitException(c *fiber.Ctx, conf *config.Config, log *logrus.Logger) InitialExceptionCreateResponse {
	init := InitialExceptionCreateResponse{
		Ctx:  c,
		Log:  log,
		Conf: conf,
	}

	return init
}
