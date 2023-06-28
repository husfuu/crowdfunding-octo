package repository

import (
	dr "github.com/husfuu/crowdfunding-octo/internal/api/user/model"
	"github.com/husfuu/crowdfunding-octo/pkg/infra/db"
)

type Repository interface {
	GetAll() ([]dr.User, error)
	GetUserByEmail(email string) (dr.User, error)
	GetUserById(id int64) (dr.User, error)
	PostUser(request dr.UserReqisterRequest) (int64, error)
	UpdateUser(request dr.UserUpdateRequest) (int64, error)
}

type UserRepo struct {
	DBList *db.DatabaseList
}

func NewUserRepo(dbList *db.DatabaseList) UserRepo {
	return UserRepo{
		DBList: dbList,
	}
}

func (ur UserRepo) GetAll() ([]dr.User, error) {
	var response []dr.User
	err := ur.DBList.APIDB.Raw(qSelectUser).Scan(response).Error
	return response, err
}

func (ur UserRepo) GetUserByEmail(email string) (dr.User, error) {
	var response dr.User
	err := ur.DBList.APIDB.Raw(qSelectUser+qWhere+qEmail, email).Scan(&response).Error
	return response, err
}

func (ur UserRepo) GetUserById(userId int64) (dr.User, error) {
	var response dr.User
	err := ur.DBList.APIDB.Raw(qSelectUser+qWhere+qUserId, userId).Scan(&response).Error

	return response, err
}

func (ur UserRepo) UpdateUser(request dr.UserUpdateRequest) (int64, error) {
	var response int64
	err := ur.DBList.APIDB.Raw(qUpdateUser+qReturningUserId, request).Scan(&response).Error
	return response, err
}

func (ur UserRepo) PostUser(request dr.UserReqisterRequest) (int64, error) {
	var response int64
	err := ur.DBList.APIDB.Raw(qInsertUser+qReturningUserId, request).Scan(&response).Error
	return response, err
}
