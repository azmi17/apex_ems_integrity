package app

import (
	"database/sql"
	"integrasi_apex_ems/helper"
	"time"
)

func NewDB() (*sql.DB, *sql.DB) {

	// Db Config to apex
	db1, err := sql.Open("mysql", "root@tcp(localhost:3317)/integrasi_apex_ems?parseTime=true")
	helper.FatalIfError(err)

	// Db Pooling apex
	db1.SetMaxIdleConns(5)
	db1.SetMaxOpenConns(20)
	db1.SetConnMaxLifetime(60 * time.Minute)
	db1.SetConnMaxIdleTime(10 * time.Minute)

	// Db Config to apex_sys
	db2, err := sql.Open("mysql", "root@tcp(localhost:3317)/integrasi_apex_ems_sys?parseTime=true")
	helper.FatalIfError(err)

	// Db Pooling apex_sys
	db2.SetMaxIdleConns(5)
	db2.SetMaxOpenConns(20)
	db2.SetConnMaxLifetime(60 * time.Minute)
	db2.SetConnMaxIdleTime(10 * time.Minute)

	return db1, db2
}
