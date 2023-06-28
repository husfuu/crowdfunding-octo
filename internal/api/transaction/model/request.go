package model

type CreateTransactionRequest struct {
	Amount     int `json:"amount" binding:"required"`
	CampaignId int `json:"campaign_id" binding:"required"`
	UserId     int `json:"user_id" binding:"required"`
}
