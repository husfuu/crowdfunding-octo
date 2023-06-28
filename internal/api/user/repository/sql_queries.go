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
	INSERT INTO "user" (name, email, occupation, password_hash)
	VALUES(?, ?, ?, ?)	
	`
	qUpdateUser = `
	UPDATE "user" 
	SET name = ?,
	email = ?,
	occupation = ?
	WHERE user_id = ?
	`
	qReturningUserId = `
	RETURNING user_id
	`
)
