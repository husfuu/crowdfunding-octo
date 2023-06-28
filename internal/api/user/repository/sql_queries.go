package repository

const (
	qSelectUser = `
	SELECT user_id, name, occupation, email, 
	avatar_file_name, role FROM "user"
	`
	qWhere = `
	WHERE
	`
	qUserId = `
	user_id = ?
	`
	qEmail = `
	email = ?
	`
	qInsertUser = `
	INSERT INTO "user" (user_id, name, occupation, email, avatar_file_name, role, created_at, updated_at)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?)	
	`
	qUpdateUser = `
	UPDATE user 
	SET name = ?,
	occupation = ?,
	email = ?,
	avatar_file_name = ?,
	role = ?,
	created_at = ?,
	updated_at = ?
	WHERE user_id = ?
	`
	qReturningUserId = `
	RETURNING user_id
	`
)
