package app

import (
	"database/sql"
	"fmt"
	"integrasi_apex_ems/helper"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {

	db1, err := sql.Open("mysql", "root@tcp(localhost:3317)/integrasi_apex_ems")
	helper.FatalIfError(err)
	defer db1.Close()

	db2, err := sql.Open("mysql", "root@tcp(localhost:3317)/integrasi_apex_ems_sys")
	helper.FatalIfError(err)
	defer db2.Close()

	fmt.Println("Connected to databases..")
}
