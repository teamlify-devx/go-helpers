package mssql

import (
	"database/sql"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
	cfg "github.com/spf13/viper"
)

// NewMssqlDB Return new Microsoft SQL client
func NewMssqlDB() (db *sql.DB, err error) {

	connStr := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;port=%d", cfg.GetString("MsSql.HOST"), cfg.GetString("MsSql.DEFAULT_DB"), cfg.GetString("MsSql.USER"), cfg.GetString("MsSql.PASS"), cfg.GetInt("MsSql.PORT"))

	db, err = sql.Open("sqlserver", connStr)
	if err != nil {
		return nil, err
	}

	return
}
