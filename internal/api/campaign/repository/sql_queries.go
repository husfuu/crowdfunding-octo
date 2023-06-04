package repository

const (
	qSelectCampaign = `
	SELECT campaign_id, user_id, name, 
	"desc", short_desc, perks, backer_count, goal_amount, 
	current_amount, slug 
	FROM campaign
	`
	qWhere = `
	WHERE
	`
	qCampaignId = `
	campaign_id = ?
	`

	qSelectCampaignImage = `
	SELECT campaign_image_id, campaign_id, 
	file_name FROM campaign_image
	`

	qInsertCampaign = `
	INSERT INTO campaign (user_id, name, short_desc, "desc", perks, goal_amount, slug)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	qReturningCampaign = `
	RETURNING campaign_id
	`

	qInsertCampaignImage = `
	INSERT INTO "campaign_image" (campaign_id, file_name)
	VALUES (?, ?)
	`

	qReturningCampaignImage = `
	RETURNING campaign_image_id
	`

	qUpdateCampaign = `
	UPDATE campaign
	SET name = ?,
	short_desc = ?, 
	"desc" = ?,
	perks = ?,
	goal_amount = ?
	WHERE campaign_id=?
	`
)
