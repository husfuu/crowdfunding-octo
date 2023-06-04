package model

type GetCampaignDetailRequest struct {
	Id int
}

type PostCampaignRequest struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             string
}

type PostCampaignImgRequest struct {
	CampaignId int64  `json:"campaign_id"`
	FileName   string `json:"file_name"`
}
