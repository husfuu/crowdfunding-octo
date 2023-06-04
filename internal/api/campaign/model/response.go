package model

import (
	"time"
)

type CampaignResponse struct {
	CampaignId       int64           `json:"campaign_id" gorm:"column:campaign_id"`
	UserId           int             `json:"user_id" gorm:"column:user_id"`
	Name             string          `json:"name" gorm:"column:name"`
	ShortDescription string          `json:"short_description" gorm:"column:short_desc"`
	Description      string          `json:"description" gorm:"column:desc"`
	Perks            string          `json:"perks" gorm:"column:perks"`
	BackerCount      int             `json:"backer_count" gorm:"column:backer_count"`
	GoalAmount       int             `json:"goal_amount" gorm:"column:goal_amount"`
	CurrentAmount    int             `json:"current_amount" gorm:"column:current_amount"`
	Slug             string          `json:"slug" gorm:"column:slug"`
	CreatedAt        time.Time       `json:"created_at" gorm:"column:created_at"`
	UpdatedAt        time.Time       `json:"updated_at" gorm:"column:updated_at"`
	CampaignImages   []CampaignImage `json:"campaign_images" gorm:"-"`
	User             string          `json:"user"`
}

type CampaignImage struct {
	CampaignImageId int64  `json:"campaign_image_id" gorm:"column:campaign_image_id"`
	CampaignID      int64  `json:"campaign_id" gorm:"column:campaign_id"`
	FileName        string `json:"image_name" gorm:"column:file_name"`
	IsPrimary       int    `json:"is_primary" gorm:"column:is_primary"`
	// CreatedAt  time.Time
	// UpdatedAt  time.Time
}
