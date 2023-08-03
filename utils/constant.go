package utils

const (
	// FOR USER REPO
	SELECT_ALL_USER     = "SELECT id, username, password, created_at, updated_at FROM user_cred"
	INSERT_USER         = "INSERT INTO user_cred (username, password, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	SELECT_USER_BY_NAME = "SELECT id, username, password, created_at, updated_at FROM user_cred WHERE username = $1"
	SELECT_USER_BY_ID   = "SELECT id, username, password, created_at, updated_at FROM user_cred WHERE id = $1"
	DELETE_USER         = "DELETE FROM user_cred WHERE id = $1"
	UPDATE_USER         = "UPDATE user_cred SET username = $2, password = $3, updated_at = $4 WHERE id = $1"
)
