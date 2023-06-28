package model

import (
	// user "github.com/husfuu/crowdfunding-octo/internal/api/user/model"
	"github.com/leekchan/accounting"
)

type Transaction struct {
	TransactionId int    `json:"trx_id" gorm:"column:trx_id"`
	CampaignId    int    `json:"campaign_id" gorm:"column:campaign_id"`
	Amount        int    `json:"amount" gorm:"column:amount"`
	Status        string `json:"status" gorm:"column:status"`
	Code          string `json:"code" gorm:"column:code"`
	PaymentURL    string `json:"payment_url" gorm:"column:payment_url"`
	UserId        int    `json:"user_id" gorm:"column:user_id"`
	UserEmail     string `json:"user_email" gorm:"column:email"`
	CampaignName  string `json:"campaign_name" gorm:"column:name"`
}

func (t Transaction) AmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(t.Amount)
}
