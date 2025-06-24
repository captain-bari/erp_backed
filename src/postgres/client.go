package postgres

import (
	"database/sql"
	log "erp/log"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5438
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

var (
	DBclient *sql.DB
)

func Init() error {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Errorf("Open: ERROR[%s]", err.Error())
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Errorf("Ping: ERROR[%s]", err.Error())
		return err
	}

	DBclient = db

	log.Info("Successfully connected!")
	return nil
}

// Helper function -> applyArgs applies args to the query with args and forms the complete query.
func applyArgs(tmpl string, args ...interface{}) string {
	for i := len(args) - 1; i >= 0; i-- {
		tmpl = strings.ReplaceAll(tmpl, fmt.Sprintf("$%d", i+1), fmt.Sprintf("%v", args[i]))
	}

	return tmpl
}
