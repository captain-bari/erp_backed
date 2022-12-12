package postgres

import (
	log "erp/log"
	types "erp/types"
	"fmt"
)

func GetUserRoleAndNameDB(id, hash string) (role, name string, noRow bool, err error) {

	log.Logger.Debugf("GetUserRoleAndNameDB Query:[%s]", applyArgs(GetUserRoleAndName, id, hash))

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

func GetMaterialsDB() (resp []types.Material, err error) {

	log.Logger.Debugf("GetUserRoleAndNameDB Query:[%s]", applyArgs(GetMaterials))

	rows, err := DBclient.Query(GetMaterials)
	if err != nil {
		err = fmt.Errorf("db: GetUserRoleAndNameDB row query ERROR[%s]", err.Error())
		log.Logger.Error(err.Error())
		return
	}

	for rows.Next() {
		var m types.Material
		err = rows.Scan(&m.ID, &m.Type, &m.Name, &m.Properties, &m.Processes, &m.Purchage, &m.Sell)
		if err != nil {
			err = fmt.Errorf("db: GetUserRoleAndNameDB row scan ERROR[%s]", err.Error())
			log.Logger.Error(err.Error())
			return
		}
		resp = append(resp, m)
	}

	return
}
