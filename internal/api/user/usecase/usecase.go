package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/husfuu/crowdfunding-octo/config"
	du "github.com/husfuu/crowdfunding-octo/internal/api/user/model"
	repo "github.com/husfuu/crowdfunding-octo/internal/wrapper/repository"
	"github.com/husfuu/crowdfunding-octo/pkg/exception"
	"github.com/husfuu/crowdfunding-octo/pkg/infra/db"
	"github.com/husfuu/crowdfunding-octo/pkg/utils"
	"github.com/sirupsen/logrus"
)

type Usecase interface {
	Login(request du.UserLoginRequest) (*du.User, *exception.Error)
	Register(request du.UserReqisterRequest) (*int64, *exception.Error)
	IsEmailAvailable(request du.EmailAvailableRequest) (bool, error)
	GetAllUser() (*[]du.User, *exception.Error)
	GetUserById(userId int64) (*du.User, *exception.Error)
	SaveAvatar(request du.AvatarRequest) (string, error)
	UpdateUser(request du.UserUpdateRequest) (*int64, *exception.Error)
}

type UserUsecase struct {
	Repo   repo.Repository
	Conf   *config.Config
	DBList *db.DatabaseList
	Log    *logrus.Logger
}

func NewUserUsecase(repo repo.Repository, conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) UserUsecase {
	return UserUsecase{
		Repo:   repo,
		Conf:   conf,
		DBList: dbList,
		Log:    log,
	}
}

func (uu UserUsecase) Login(request du.UserLoginRequest) (*du.User, *exception.Error) {
	user, err := uu.Repo.Api.User.GetUserByEmail(request.Email)
	if err != nil {
		uu.Log.Error(err)
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if user.UserId == 0 {
		uu.Log.Info("user is not found")
		return nil, exception.NewError(fiber.StatusNotFound, "user is not registered")
	}

	// check valid password
	if !utils.CheckPasswordHash(request.Password, user.PasswordHash) {
		return nil, exception.NewError(fiber.StatusUnauthorized, "fail to login, email or password is not correct. ")
	}

	return &user, nil
}

func (uu UserUsecase) GetAllUser() (*[]du.User, *exception.Error) {
	user, err := uu.Repo.Api.User.GetAll()
	if err != nil {
		uu.Log.Info(err)
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return &user, nil
}

func (uu UserUsecase) GetUserById(userId int64) (*du.User, *exception.Error) {
	user, err := uu.Repo.Api.User.GetUserById(userId)
	if err != nil {
		uu.Log.Info(err)
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if user.UserId == 0 {
		return nil, exception.NewError(fiber.StatusNotFound, "user not found.")
	}

	return &user, nil
}

func (uu UserUsecase) Register(request du.UserReqisterRequest) (*int64, *exception.Error) {
	hashedPass, err := utils.HashPassword(request.Password)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, "fail to hash password. "+err.Error())
	}
	request.Password = hashedPass
	userId, err := uu.Repo.Api.User.PostUser(request)
	if err != nil {
		uu.Log.Info(err)
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return &userId, nil
}

func (uu UserUsecase) SaveAvatar(request du.AvatarRequest) (string, error) {
	return "success upload image", nil
}

func (uu UserUsecase) IsEmailAvailable(request du.EmailAvailableRequest) (bool, *exception.Error) {
	user, err := uu.Repo.Api.User.GetUserByEmail(request.Email)
	if err != nil {
		uu.Log.Error(err)
		return false, exception.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// email is not available because there's used before
	if user.UserId != 0 {
		return false, nil
	}

	// email is available because there's no used before
	return true, nil
}

func (uu UserUsecase) UpdateUser(request du.UserUpdateRequest) (*int64, *exception.Error) {
	user, err := uu.Repo.Api.User.GetUserById(request.UserID)
	if err != nil {
		uu.Log.Info(err)
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if user.UserId == 0 {
		return nil, exception.NewError(fiber.StatusNotFound, "user not found. ")
	}

	userId, err := uu.Repo.Api.User.UpdateUser(request)
	if err != nil {
		uu.Log.Info(err)
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return &userId, nil
}

// func (uu UserUsecase) GetUserByEmail(email string) (*du.User, *exception.Error) {
// 	user, err := uu.Repo.Api.User.GetUserByEmail(email)
// 	if err != nil {
// 		uu.Log.Info(err)
// 		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error())
// 	}
// 	return &user, nil
// }
