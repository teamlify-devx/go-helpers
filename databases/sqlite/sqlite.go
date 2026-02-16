package sqlite

import (
	"database/sql"
	cfg "github.com/spf13/viper"
	_ "modernc.org/sqlite"
	"os"
	"path/filepath"
)

// NewSqliteDB Return new Sqlite db instance
func NewSqliteDB() (db *sql.DB, err error) {
	dir, err := os.MkdirTemp("", cfg.GetString("sqlite.DB_DIR"))
	if err != nil {
		return nil, err
	}

	defer os.RemoveAll(dir)

	fn := filepath.Join(dir, cfg.GetString("sqlite.DB_NAME"))

	db, err = sql.Open("sqlite", fn)
	if err != nil {
		return nil, err
	}

	return
}
