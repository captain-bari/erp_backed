package postgres

const (
	GetUserRoleAndName = `SELECT role,name FROM users where id=$1 and user_hash=$2;`
	GetMaterials       = `SELECT id, e_type, name, properties, linked_processes, linked_purchase, linked_sell FROM materials;`
	InsertMaterials    = `INSERT INTO materials(e_type, name, properties, linked_processes, linked_purchase, linked_sell)
	VALUES ($1,$2,$3,$4,$5,$6);`
)
