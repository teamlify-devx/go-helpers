package clickhouse

import (
	"database/sql"
	"fmt"
	"github.com/ClickHouse/clickhouse-go"
	cfg "github.com/spf13/viper"
)

// NewClickHouseDB Return new Click House client
func NewClickHouseDB() (db *sql.DB, err error) {
	connStr := fmt.Sprintf("tcp://%s:%s?username=%s&password=%s&database=%s&read_timeout=10&write_timeout=20&debug=true", cfg.GetString("Clickhouse.HOST"), cfg.GetString("Clickhouse.PORT"), cfg.GetString("Clickhouse.USER"), cfg.GetString("Clickhouse.PASS"), cfg.GetString("Clickhouse.DEFAULT_DB"))

	db, err = sql.Open("clickhouse", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		return nil, err
	}

	return
}
