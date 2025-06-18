package auth

import (
	log "erp/log"
	"erp/postgres"
)

func AuthenticateUser(userID, userHash string) (notFound bool, name, role string, err error) {
	role, name, notFound, err = postgres.GetUserRoleAndNameDB(userID, userHash)
	if err != nil {
		log.Errorf("postgres.GetUserRoleAndNameDB: %s", err.Error())
	}

	if notFound {
		log.Debugf("No user found with ID:[%s] HASH:[%s]", userID, userHash)
	}
	return
}
