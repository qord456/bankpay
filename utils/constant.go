package utils

const (
	// FOR USER REPO
	SELECT_ALL_USER     = "SELECT id, username, password, created_at, updated_at FROM user_cred"
	INSERT_USER         = "INSERT INTO user_cred (username, password, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	SELECT_USER_BY_NAME = "SELECT id, username, password, created_at, updated_at FROM user_cred WHERE username = $1"
	SELECT_USER_BY_ID   = "SELECT id, username, password, created_at, updated_at FROM user_cred WHERE id = $1"
	DELETE_USER         = "DELETE FROM user_cred WHERE id = $1"
	UPDATE_USER         = "UPDATE user_cred SET username = $2, password = $3, updated_at = $4 WHERE id = $1"

	// FOR CUSTOMER REPO
	SELECT_ALL_CUSTOMER     = "SELECT id, user_id,nik, name, email, phone, address, birthdate, balance, status FROM customer;"
	SELECT_CUSTOMER_BY_NAME = "SELECT id, user_id,nik, name, email, phone, address, birthdate, balance, status FROM customer WHERE name = $1;"
	SELECT_CUSTOMER_BY_ID   = "SELECT id, user_id,nik, name, email, phone, address, birthdate, balance, status FROM customer WHERE id = $1;"
	REGISTER_CUSTOMER       = "INSERT INTO customer (user_id,nik, name, email, phone, address, birthdate, balance, status) VALUES ($1, $2, $3, $4,$5, $6, $7, $8)"
	DELETE_CUSTOMER         = "DELETE FROM customer WHERE id = $1"
	UPDATE_CUSTOMER         = "UPDATE customer SET user_id=$2, nik=$3, name=$4, email=$5, phone=$6, address=$7, birthdate = $8 WHERE id = $1"
	STATUS_CUSTOMER         = "UPDATE customer SET status = $2 WHERE id=$1"
	GET_BALANCE             = "SELECT balance FROM customer WHERE id = $1 "
	UPDATE_BALANCE          = "UPDATE customer SET balance = $1"

	//FOR PAYMENT REPO
	INSERT_PAYMENT                = "INSERT INTO payment (customer_id, paid, destination_id, created_by)"
	SELECT_PAYMENT_BY_CUSTOMER_ID = "SELECT id, paid, destination_id, created_at, created_by FROM payment WHERE customer_id = $1"
	SELECT_ALL_PAYMENT            = "SELECT id, customer_id, paid, destination_id, created_at, created_by FROM payment"
	SELECT_PAYMENT_BY_ID          = "SELECT id, customer_id, paid, destination_id, created_at, created_by FROM payment WHERE id = $1 && customer_id = $2"
)
