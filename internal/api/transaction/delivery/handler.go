package delivery

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/husfuu/crowdfunding-octo/config"
	"github.com/husfuu/crowdfunding-octo/internal/wrapper/usecase"
	"github.com/husfuu/crowdfunding-octo/pkg/exception"
	"github.com/sirupsen/logrus"
)

type TransactionHandler struct {
	Usecase usecase.Usecase
	Conf    *config.Config
	Log     *logrus.Logger
}

func NewTransactionHandler(usecase usecase.Usecase, conf *config.Config, log *logrus.Logger) TransactionHandler {
	return TransactionHandler{
		Usecase: usecase,
		Conf:    conf,
		Log:     log,
	}
}

func (th TransactionHandler) GetTransactionById(c *fiber.Ctx) error {
	initException := exception.InitException(c, th.Conf, th.Log)
	trxIdStr := c.Query("trx_id", "0")
	trx_id, err := strconv.Atoi(trxIdStr)
	if err != nil {
		return exception.CreateResponse_Log(initException, fiber.StatusBadRequest, "UserId must integer!", nil)
	}
	respData, errResp := th.Usecase.Api.Transaction.GetTransactionById(trx_id)

	if errResp != nil {
		return exception.CreateResponse_Log(initException, errResp.Code, errResp.Message, nil)
	}
	message := fmt.Sprintf("success to fetch transaction with trx_id = %s", trxIdStr)
	return exception.CreateResponse_Log(initException, fiber.StatusOK, message, respData)

}

func (th TransactionHandler) GetTransactionByCampaignId(c *fiber.Ctx) error {
	initException := exception.InitException(c, th.Conf, th.Log)

	campaignIdStr := c.Query("campaign_id", "0")
	campaign_id, err := strconv.Atoi(campaignIdStr)
	if err != nil {
		return exception.CreateResponse_Log(initException, fiber.StatusBadRequest, "UserId must integer!", nil)
	}
	respData, errResp := th.Usecase.Api.Transaction.GetTransactionByCampaignId(campaign_id)

	if errResp != nil {
		return exception.CreateResponse_Log(initException, errResp.Code, errResp.Message, nil)
	}
	message := fmt.Sprintf("success to fetch transaction with campaign_id = %s", campaignIdStr)
	return exception.CreateResponse_Log(initException, fiber.StatusOK, message, respData)
}

func (th TransactionHandler) GetTransactionByUserId(c *fiber.Ctx) error {
	initException := exception.InitException(c, th.Conf, th.Log)

	userIdStr := c.Query("user_id", "0")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return exception.CreateResponse_Log(initException, fiber.StatusBadRequest, "UserId must integer!", nil)
	}
	respData, errResp := th.Usecase.Api.Transaction.GetTransactionByUserId(userId)

	if errResp != nil {
		return exception.CreateResponse_Log(initException, errResp.Code, errResp.Message, nil)
	}
	message := fmt.Sprintf("success to fetch transaction with user_id = %s", userIdStr)
	return exception.CreateResponse_Log(initException, fiber.StatusOK, message, respData)
}
