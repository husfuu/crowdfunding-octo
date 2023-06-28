package usecase

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/husfuu/crowdfunding-octo/config"
	du "github.com/husfuu/crowdfunding-octo/internal/api/transaction/model"
	repo "github.com/husfuu/crowdfunding-octo/internal/wrapper/repository"
	"github.com/husfuu/crowdfunding-octo/pkg/exception"
	"github.com/husfuu/crowdfunding-octo/pkg/infra/db"
	"github.com/sirupsen/logrus"
)

type Usecase interface {
	GetTransactionByCampaignId(campaignId int) ([]du.Transaction, *exception.Error)
	GetTransactionByUserId(userId int) ([]du.Transaction, *exception.Error)
	GetTransactionById(trx int) (*du.Transaction, *exception.Error)
	SaveTransaction(du.CreateTransactionRequest) (*du.Transaction, *exception.Error)
}

type TransactionUsecase struct {
	Repo   repo.Repository
	Conf   *config.Config
	DBList *db.DatabaseList
	Log    *logrus.Logger
}

func NewTransactionUsecas(repo repo.Repository, conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) TransactionUsecase {
	return TransactionUsecase{
		Repo:   repo,
		Conf:   conf,
		DBList: dbList,
		Log:    log,
	}
}

func (tu TransactionUsecase) GetTransactionById(trxId int) (*du.Transaction, *exception.Error) {
	transaction, err := tu.Repo.Api.Transaction.GetTransactionById(trxId)
	if err != nil {
		tu.Log.Error(err)
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if transaction.TransactionId == 0 {
		return nil, exception.NewError(fiber.StatusNotFound, fmt.Sprintf("transaction with trx_id = %d is not found", trxId))
	}
	return &transaction, nil
}

func (tu TransactionUsecase) GetTransactionByCampaignId(campaignId int) ([]du.Transaction, *exception.Error) {
	transactions, err := tu.Repo.Api.Transaction.GetTransactionByCampaignId(campaignId)
	if err != nil {
		tu.Log.Error(err)
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return transactions, nil
}

func (tu TransactionUsecase) GetTransactionByUserId(userId int) ([]du.Transaction, *exception.Error) {
	transaction, err := tu.Repo.Api.Transaction.GetTransactionByUserId(userId)
	if err != nil {
		tu.Log.Error(err)
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return transaction, nil
}

func (tu TransactionUsecase) SaveTransaction(request du.CreateTransactionRequest) (*du.Transaction, *exception.Error) {
	const status = "pending"
	transaction := du.Transaction{}
	transaction.CampaignId = request.CampaignId
	transaction.Amount = request.Amount
	transaction.UserId = request.UserId
	transaction.Status = status

	// newTransaction, err := tu.Repo.Api.Transaction.SaveTransaction(transaction)
	// if err != nil {
	// 	tu.Log.Error(err)
	// 	return nil, exception.NewError(fiber.StatusInternalServerError, err.Error())
	// }

	// paymentTransaction, err := tu.Repo.Api.Transaction

	return nil, nil
}
