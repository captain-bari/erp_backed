package postgres

const (
	GetUserRoleAndName = `SELECT role,name FROM users where id=$1 and user_hash=$2;`
)
