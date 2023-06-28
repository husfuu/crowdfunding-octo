package repository

import (
	dr "github.com/husfuu/crowdfunding-octo/internal/api/transaction/model"
	"github.com/husfuu/crowdfunding-octo/pkg/infra/db"
)

type Repository interface {
	GetTransactionByCampaignId(campaignId int) ([]dr.Transaction, error)
	GetTransactionByUserId(userId int) ([]dr.Transaction, error)
	GetTransactionById(trxId int) (dr.Transaction, error)
	SaveTransaction(request dr.Transaction) (dr.Transaction, error)
	// Update(request int) (int, error)
	// FindAll() ([]dr.Transaction, error)
}

type TransactionRepo struct {
	DBList *db.DatabaseList
}

func NewTransactionRepo(dbList *db.DatabaseList) TransactionRepo {
	return TransactionRepo{
		DBList: dbList,
	}
}

func (tr TransactionRepo) GetTransactionByCampaignId(campaignId int) ([]dr.Transaction, error) {
	var response []dr.Transaction
	err := tr.DBList.APIDB.Raw(qSelectTransaction+qWhere+qCampaignId, campaignId).Scan(&response).Error
	return response, err
}

func (tr TransactionRepo) GetTransactionByUserId(userId int) ([]dr.Transaction, error) {
	var response []dr.Transaction
	err := tr.DBList.APIDB.Raw(qSelectTransaction+qWhere+qUserId, userId).Scan(&response).Error
	return response, err
}

func (tr TransactionRepo) GetTransactionById(trxId int) (dr.Transaction, error) {
	var response dr.Transaction
	err := tr.DBList.APIDB.Raw(qSelectTransaction+qWhere+qTransactionId, trxId).Scan(&response).Error
	return response, err
}

func (tr TransactionRepo) SaveTransaction(request dr.Transaction) (dr.Transaction, error) {
	var response dr.Transaction
	params := make([]interface{}, 0)
	params = append(params,
		request.CampaignId,
		request.UserId,
		request.Amount,
		request.Status,
	)
	err := tr.DBList.APIDB.Raw(qInsertTransaction, params...).Scan(response).Error

	return response, err
}

// func (tr TransactionRepo) Update(request int) (int, error) {

// }
// func (tr TransactionRepo) FindAll() ([]dr.Transaction, error) {

// }
