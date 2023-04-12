package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Clientdb *sql.DB
	username = "freedb_salim"
	password = "N&bBHhh8S!!ZW8@"
	host     = "sql.freedb.tech"
	schema   = "freedb_transactions"
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, schema)
	var err error
	Clientdb, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
}
