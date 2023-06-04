package usecase

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/husfuu/crowdfunding-octo/config"
	du "github.com/husfuu/crowdfunding-octo/internal/api/campaign/model"
	repo "github.com/husfuu/crowdfunding-octo/internal/wrapper/repository"
	"github.com/husfuu/crowdfunding-octo/pkg/exception"
	"github.com/husfuu/crowdfunding-octo/pkg/infra/db"
	"github.com/sirupsen/logrus"
)

type Usecase interface {
	GetCampaignById(id int64) (*du.CampaignResponse, *exception.Error)
	PostCampaign(dataReq du.PostCampaignRequest) (*du.CampaignResponse, *exception.Error)
	PostCImage(dataReq du.PostCampaignImgRequest) (*du.CampaignImage, *exception.Error)
	UpdateCampaign(dataReq du.PostCampaignRequest) (*du.CampaignResponse, *exception.Error)
}

type CampaignUsecase struct {
	Repo   repo.Repository
	Conf   *config.Config
	DBList *db.DatabaseList
	Log    *logrus.Logger
}

func NewCampaignUsecase(repo repo.Repository, conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) CampaignUsecase {
	return CampaignUsecase{
		Repo:   repo,
		Conf:   conf,
		DBList: dbList,
		Log:    log,
	}
}

func (cu CampaignUsecase) GetCampaignById(id int64) (*du.CampaignResponse, *exception.Error) {
	campaign, err := cu.Repo.Api.Campaign.GetCampaignById(id)

	if err != nil {
		cu.Log.Info(err)
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error())
	}

	campaignImgs, err := cu.Repo.Api.Campaign.GetCImage(campaign.CampaignId)

	if err != nil {
		cu.Log.Info(err)
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error())
	}

	campaign.CampaignImages = campaignImgs

	return &campaign, nil
}

func (cu CampaignUsecase) PostCampaign(dataReq du.PostCampaignRequest) (*du.CampaignResponse, *exception.Error) {
	slugCandidate := fmt.Sprintf("%s %d", dataReq.Name, 1)
	slug_data := slug.Make(slugCandidate)
	user_id := 1

	params := make([]interface{}, 0)
	params = append(params,
		user_id,
		dataReq.Name,
		dataReq.ShortDescription,
		dataReq.Description,
		dataReq.Perks,
		dataReq.GoalAmount,
		slug_data,
	)

	respId, err := cu.Repo.Api.Campaign.PostCampaign(params...)
	if err != nil {
		cu.Log.Info(err)
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// repo get campaign
	campaign, _ := cu.Repo.Api.Campaign.GetCampaignById(respId)
	if err != nil {
		cu.Log.Info(err)
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &campaign, nil
}

func (cu CampaignUsecase) PostCImage(dataReq du.PostCampaignImgRequest) (*int64, *exception.Error) {
	_, err := cu.Repo.Api.Campaign.GetCampaignById(dataReq.CampaignId)
	if err != nil {
		cu.Log.Info(err)
		return nil, exception.NewError(fiber.StatusNotFound, err.Error())
	}

	//TODO! validate the owner(user) campaign

	params := make([]interface{}, 0)
	params = append(params,
		dataReq.CampaignId,
		dataReq.FileName,
	)

	respId, err := cu.Repo.Api.Campaign.PostCImage(params...)
	if err != nil {
		cu.Log.Info(err)
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &respId, nil

}

func (cu CampaignUsecase) UpdateCampaign(dataReq du.PostCampaignRequest, campaignId int64) (*du.CampaignResponse, *exception.Error) {
	_, err := cu.Repo.Api.Campaign.GetCampaignById(campaignId)
	if err != nil {
		cu.Log.Info(err)
		return nil, exception.NewError(fiber.StatusNotFound, err.Error())
	}

	params := make([]interface{}, 0)
	params = append(params,
		dataReq.Name,
		dataReq.ShortDescription,
		dataReq.Description,
		dataReq.Perks,
		dataReq.GoalAmount,
		campaignId,
	)

	respId, errResp := cu.Repo.Api.Campaign.UpdateCampaign(params...)
	if errResp != nil {
		cu.Log.Info(err)
		return nil, exception.NewError(fiber.StatusInternalServerError, errResp.Error())
	}

	//* repo get campaign
	respData, errResp := cu.Repo.Api.Campaign.GetCampaignById(respId)
	if errResp != nil {
		cu.Log.Info(err)
		return nil, exception.NewError(fiber.StatusInternalServerError, errResp.Error())
	}

	return &respData, nil
}
