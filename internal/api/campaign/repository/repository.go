package repository

import (
	"fmt"

	dr "github.com/husfuu/crowdfunding-octo/internal/api/campaign/model"
	"github.com/husfuu/crowdfunding-octo/pkg/infra/db"
)

type Repository interface {
	GetCampaignById(id int64) (dr.CampaignResponse, error)
	GetCImage(campaignId int64) ([]dr.CampaignImage, error)
	PostCampaign(params ...interface{}) (int64, error)
	PostCImage(params ...interface{}) (int64, error)
	UpdateCampaign(params ...interface{}) (int64, error)
}

type CampaignRepo struct {
	DBList *db.DatabaseList
}

func NewCampaignRepo(dbList *db.DatabaseList) CampaignRepo {
	return CampaignRepo{
		DBList: dbList,
	}
}

func (cr CampaignRepo) GetCampaignById(id int64) (dr.CampaignResponse, error) {
	var response dr.CampaignResponse
	err := cr.DBList.APIDB.Raw(qSelectCampaign+qWhere+qCampaignId, id).Scan(&response).Error

	return response, err
}

func (cr CampaignRepo) GetCImage(campaignId int64) ([]dr.CampaignImage, error) {
	var response []dr.CampaignImage
	err := cr.DBList.APIDB.Raw(qSelectCampaignImage+qWhere+qCampaignId, campaignId).Scan(&response).Error

	return response, err
}

func (cr CampaignRepo) PostCampaign(params ...interface{}) (int64, error) {
	var response int64
	err := cr.DBList.APIDB.Raw(qInsertCampaign+qReturningCampaign, params...).Scan(&response).Error

	return response, err
}

func (cr CampaignRepo) PostCImage(params ...interface{}) (int64, error) {
	var response int64
	err := cr.DBList.APIDB.Raw(qInsertCampaignImage+qReturningCampaignImage, params...).Scan(&response).Error

	return response, err
}

func (cr CampaignRepo) UpdateCampaign(params ...interface{}) (int64, error) {
	var response int64

	fmt.Println(params...)
	err := cr.DBList.APIDB.Raw(qUpdateCampaign+qReturningCampaign, params...).Scan(&response).Error

	return response, err
}
