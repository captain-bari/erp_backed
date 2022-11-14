package postgres

import (
	log "erp/log"
	"fmt"
)

func GetUserRoleAndNameDB(id, hash string) (role, name string, noRow bool, err error) {

	log.Logger.Debug("GetUserRoleAndNameDB Query:[%s]", applyArgs(id, hash))

	rows, err := DBclient.Query(GetUserRoleAndName, id, hash)
	if err != nil {
		err = fmt.Errorf("db: GetUserRoleAndNameDB row query ERROR[%s]", err.Error())
		log.Logger.Error(err.Error())
		return
	}

	for rows.Next() {
		err = rows.Scan(&role, &name)
		if err != nil {
			err = fmt.Errorf("db: GetUserRoleAndNameDB row scan ERROR[%s]", err.Error())
			log.Logger.Error(err.Error())
			return
		}
	}

	if role == "" && name == "" {
		noRow = true
	}

	return
}
