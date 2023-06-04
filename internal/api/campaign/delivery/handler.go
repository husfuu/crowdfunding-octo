package delivery

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/husfuu/crowdfunding-octo/config"
	dh "github.com/husfuu/crowdfunding-octo/internal/api/campaign/model"
	"github.com/husfuu/crowdfunding-octo/internal/wrapper/usecase"
	"github.com/husfuu/crowdfunding-octo/pkg/exception"
	"github.com/sirupsen/logrus"
)

type CampaignHandler struct {
	Usecase usecase.Usecase
	Conf    *config.Config
	Log     *logrus.Logger
}

func NewCampaignHandler(usecase usecase.Usecase, conf *config.Config, log *logrus.Logger) CampaignHandler {
	return CampaignHandler{
		Usecase: usecase,
		Conf:    conf,
		Log:     log,
	}
}

// func (ch CampaignHandler) GetListCampaigns(c *fiber.Ctx) error {
// 	userId, _ := strconv.Atoi(c.Query("user_id"))

// 	initException := exception.InitException(c, ch.Conf, ch.Log)

// 	// usecase process
// 	var respData []model.CampaignResponse
// 	var campaignImages []model.CampaignImage

// 	campaignImage := model.CampaignImage{
// 		ID:         1,
// 		CampaignID: 1,
// 		ImgName:    "image_name.jpg",
// 	}

// 	campaignImages = append(campaignImages, campaignImage)

// 	campaign1 := model.CampaignResponse{
// 		Id:               1,
// 		UserID:           11,
// 		Name:             "husni",
// 		ShortDescription: "short_desc",
// 		Description:      "desc",
// 		Perks:            "perks",
// 		BackerCount:      1,
// 		GoalAmount:       1,
// 		CurrentAmount:    1,
// 		Slug:             "campaign",
// 		CreatedAt:        time.Now(),
// 		UpdatedAt:        time.Now(),
// 		CampaignImages:   campaignImages,
// 		User:             "user 1",
// 	}

// 	respData = append(respData, campaign1)

// 	message := fmt.Sprintf("Success to get all campaigns by userId %d", userId)
// 	return exception.CreateResponse_Log(initException, fiber.StatusOK, message, respData)
// }

func (ch CampaignHandler) GetCampaign(c *fiber.Ctx) error {
	ch.Log.Info("=========== Get Campaign By id ===========")

	initException := exception.InitException(c, ch.Conf, ch.Log)

	payload, err := c.ParamsInt("id") // 	// payload, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return exception.CreateResponse_Log(initException, fiber.StatusBadRequest, "Error parse body", nil)
	}
	// usecase process
	respData, errResp := ch.Usecase.Api.Campaign.GetCampaignById(int64(payload))

	if errResp != nil {
		return exception.CreateResponse_Log(initException, fiber.StatusBadRequest, "Error usecase", nil)
	}

	message := fmt.Sprintf("Success to get campaign by id %d", respData.CampaignId)
	return exception.CreateResponse_Log(initException, fiber.StatusOK, message, respData)
}

func (ch CampaignHandler) PostCampaign(c *fiber.Ctx) error {
	initException := exception.InitException(c, ch.Conf, ch.Log)

	payload := new(dh.PostCampaignRequest)
	err := c.BodyParser(payload)

	if err != nil {
		return exception.CreateResponse_Log(initException, fiber.StatusBadRequest, "", nil)
	}

	// usecase process
	respData, errResp := ch.Usecase.Api.Campaign.PostCampaign(*payload)
	if errResp != nil {
		return exception.CreateResponse_Log(initException, fiber.StatusBadRequest, "Error usecase", nil)
	}

	message := "Success to post a new campaign"
	return exception.CreateResponse_Log(initException, fiber.StatusOK, message, respData)
}

func (ch CampaignHandler) PostCImage(c *fiber.Ctx) error {
	initException := exception.InitException(c, ch.Conf, ch.Log)

	payload := new(dh.PostCampaignImgRequest)
	err := c.BodyParser(payload)

	if err != nil {
		return exception.CreateResponse_Log(initException, fiber.StatusBadRequest, err.Error(), nil)
	}

	// usecase
	respData, errResp := ch.Usecase.Api.Campaign.PostCImage(*payload)
	if errResp != nil {
		return exception.CreateResponse_Log(initException, fiber.StatusBadRequest, errResp.Message, nil)
	}

	message := "Success to post a new campaign image"
	return exception.CreateResponse_Log(initException, fiber.StatusOK, message, respData)
}

func (ch CampaignHandler) UpdateCampaign(c *fiber.Ctx) error {
	initException := exception.InitException(c, ch.Conf, ch.Log)

	campaignId, err := c.ParamsInt("id") // campaignId, err := strconv.ParseInt(c.Params("campaign_id"), 10, 64)
	if err != nil {
		return exception.CreateResponse_Log(initException, fiber.StatusBadRequest, "Error parse body", nil)
	}

	payload := new(dh.PostCampaignRequest)
	err = c.BodyParser(payload)
	if err != nil {
		return exception.CreateResponse_Log(initException, fiber.StatusBadRequest, "", nil)
	}

	// usecase process
	respData, errResp := ch.Usecase.Api.Campaign.UpdateCampaign(*payload, int64(campaignId))
	if errResp != nil {
		return exception.CreateResponse_Log(initException, fiber.StatusBadRequest, errResp.Message, nil)
	}

	message := fmt.Sprintf("Success to update campaign wit id = %d", campaignId)
	return exception.CreateResponse_Log(initException, fiber.StatusOK, message, respData)

}

// func (ch CampaignHandler) CreateCampaign(c *fiber.Ctx) error {
// 	initException := exception.InitException(c, ch.Conf, ch.Log)

// 	payload := new(model.CreateCampaignRequest)
// 	err := c.BodyParser(payload)

// 	if err != nil {
// 		return exception.CreateResponse_Log(initException, fiber.StatusBadRequest, "", nil)
// 	}

// 	//usecase process
// 	var campaignImages []model.CampaignImage
// 	campaignImage := model.CampaignImage{
// 		ID:         1,
// 		CampaignID: 1,
// 		ImgName:    "gambar",
// 	}

// 	campaignImages = append(campaignImages, campaignImage)
// 	respData := model.CampaignResponse{
// 		Id:               1,
// 		UserID:           11,
// 		Name:             "husni",
// 		ShortDescription: "short_desc",
// 		Description:      "desc",
// 		Perks:            "perks",
// 		BackerCount:      1,
// 		GoalAmount:       1,
// 		CurrentAmount:    1,
// 		Slug:             "campaign",
// 		CreatedAt:        time.Now(),
// 		UpdatedAt:        time.Now(),
// 		CampaignImages:   campaignImages,
// 		User:             "user 1",
// 	}

// 	message := "Success to create a new campaign"
// 	return exception.CreateResponse_Log(initException, fiber.StatusOK, message, respData)
// }
