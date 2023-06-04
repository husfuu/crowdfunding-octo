package exception

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateResponse_Log(exc InitialExceptionCreateResponse, code int, message string, data interface{}) error {
	// requestId [SOON]

	RespData := ResponseData{
		Data: data,
		Status: StatusResponseData{
			Code:    code,
			Message: message,
		},
		TimeStamp: time.Now(),
	}

	if code == fiber.StatusOK || code == fiber.StatusCreated {
		RespData.Status.Type = "SUCCESS"
	} else {
		RespData.Status.Type = "ERROR"
	}

	switch code {
	case fiber.StatusOK:
		RespData.Status.Name = "OK"
	case fiber.StatusCreated:
		RespData.Status.Name = "CREATED"
	case fiber.StatusBadRequest:
		RespData.Status.Name = "BAD REQUEST"
	case fiber.StatusUnauthorized:
		RespData.Status.Name = "UNAUTHORIZED"
	case fiber.StatusNotFound:
		RespData.Status.Name = "NOT FOUND"
	case fiber.StatusInternalServerError:
		RespData.Status.Name = "INTERNAL SERVER ERROR"
	}

	return exc.Ctx.Status(code).JSON(RespData)
}
