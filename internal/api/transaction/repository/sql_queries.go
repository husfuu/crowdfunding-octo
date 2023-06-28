package repository

const (
	qSelectTransaction = `
	SELECT t.trx_id, t.campaign_id, t.user_id, t.amount, t.status, t.code, t.payment_url,
	u.user_id, u.email, c.name
	FROM transaction t
	JOIN "user" u ON t.user_id = u.user_id
	JOIN campaign c ON t.campaign_id = c.campaign_id
	`
	qWhere = `
	where
	`

	qUserId = `
	t.user_id = ?
	`
	qTransactionId = `
	t.trx_id = ?
	`
	qCampaignId = `
	t.campaign_id = ?
	`
	qInsertTransaction = `
	insert INTO transaction(campaign_id, user_id, amount, status)
	values(?, ?, ?, ?, ?) returning trx_id, campaign_id, user_id, amount, status
	`
)
