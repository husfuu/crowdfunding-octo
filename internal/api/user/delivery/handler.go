package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/husfuu/crowdfunding-octo/config"
	dh "github.com/husfuu/crowdfunding-octo/internal/api/user/model"

	"github.com/husfuu/crowdfunding-octo/internal/wrapper/usecase"
	"github.com/husfuu/crowdfunding-octo/pkg/exception"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	Usecase usecase.Usecase
	Conf    *config.Config
	Log     *logrus.Logger
}

func NewUserHandler(usecase usecase.Usecase, conf *config.Config, log *logrus.Logger) UserHandler {
	return UserHandler{
		Usecase: usecase,
		Conf:    conf,
		Log:     log,
	}
}

func (uh UserHandler) Login(c *fiber.Ctx) error {
	initException := exception.InitException(c, uh.Conf, uh.Log)
	var payload dh.UserLoginRequest
	if err := c.BodyParser(&payload); err != nil {
		return exception.CreateResponse_Log(initException, fiber.StatusInternalServerError, "Error body parser", nil)
	}
	respData, errResp := uh.Usecase.Api.User.Login(payload)
	if errResp != nil {
		return exception.CreateResponse_Log(initException, errResp.Code, errResp.Message, nil)
	}
	message := "success to login"
	return exception.CreateResponse_Log(initException, fiber.StatusOK, message, respData)
}

func (uh UserHandler) Register(c *fiber.Ctx) error {
	initException := exception.InitException(c, uh.Conf, uh.Log)
	var payload dh.UserReqisterRequest

	if err := c.BodyParser(&payload); err != nil {
		return exception.CreateResponse_Log(initException, fiber.StatusInternalServerError, "Error body parser", nil)
	}
	respData, errResp := uh.Usecase.Api.User.Register(payload)
	if errResp != nil {
		return exception.CreateResponse_Log(initException, fiber.StatusBadRequest, "", nil)
	}
	message := "success to register"
	return exception.CreateResponse_Log(initException, fiber.StatusOK, message, respData)
}

func (uh UserHandler) CheckEmailAvailability(c *fiber.Ctx) error {
	initException := exception.InitException(c, uh.Conf, uh.Log)
	var payload dh.EmailAvailableRequest
	if err := c.BodyParser(&payload); err != nil {
		return exception.CreateResponse_Log(initException, fiber.StatusInternalServerError, "Error body parser", nil)
	}
	respData, errResp := uh.Usecase.Api.User.IsEmailAvailable(payload)
	if errResp != nil {
		return exception.CreateResponse_Log(initException, fiber.StatusBadRequest, "", nil)
	}
	message := "success to check email availability"
	return exception.CreateResponse_Log(initException, fiber.StatusOK, message, respData)
}

func (uh UserHandler) UploadAvatar(c *fiber.Ctx) error {
	initException := exception.InitException(c, uh.Conf, uh.Log)
	var payload dh.AvatarRequest
	if err := c.BodyParser(&payload); err != nil {
		return exception.CreateResponse_Log(initException, fiber.StatusInternalServerError, "Error body parser", nil)
	}
	respData, errResp := uh.Usecase.Api.User.SaveAvatar(payload)
	if errResp != nil {
		return exception.CreateResponse_Log(initException, fiber.StatusBadRequest, "", nil)
	}

	message := "success to upload the avatar"
	return exception.CreateResponse_Log(initException, fiber.StatusOK, message, respData)
}

func (uh UserHandler) UpdateUser(c *fiber.Ctx) error {
	initException := exception.InitException(c, uh.Conf, uh.Log)
	var payload dh.UserUpdateRequest
	if err := c.BodyParser(&payload); err != nil {
		return exception.CreateResponse_Log(initException, fiber.StatusInternalServerError, "Error body parser", nil)
	}

	// usecase process
	respData, errResp := uh.Usecase.Api.User.UpdateUser(payload)
	if errResp != nil {
		return exception.CreateResponse_Log(initException, fiber.StatusBadRequest, "Error usecase", nil)
	}

	message := "success to update user"
	return exception.CreateResponse_Log(initException, fiber.StatusOK, message, respData)
}

func (uh UserHandler) GetUser(c *fiber.Ctx) error {
	initException := exception.InitException(c, uh.Conf, uh.Log)
	payload, err := c.ParamsInt("id")
	if err != nil {
		return exception.CreateResponse_Log(initException, fiber.StatusBadRequest, "Error parse body", nil)
	}
	// usecase process
	respData, errResp := uh.Usecase.Api.User.GetUserById(int64(payload))
	if errResp != nil {
		return exception.CreateResponse_Log(initException, fiber.StatusBadRequest, "Error usecase", nil)
	}
	message := "success to get user"
	return exception.CreateResponse_Log(initException, fiber.StatusOK, message, respData)
}
