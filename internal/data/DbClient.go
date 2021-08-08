package data

import (
	"HMThird/internal/conf"
	"database/sql"
	"log"
)

func GetDbClient(dbName string, logger log.Logger) (*sql.DB, error) {
	driver, datasource, err := conf.GetConnectionString(dbName)
	if err != nil {
		return nil, err
	}
	db, _ := sql.Open(driver, datasource)
	if err := db.Ping(); err != nil {
		logger.Printf("error occured during opening db, dbName : %s", dbName)
		return nil, err
	}
	return db, nil
}
