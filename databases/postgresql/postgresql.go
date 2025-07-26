package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	cfg "github.com/spf13/viper"
)

// NewPsqlDB Return new Postgresql db instance
func NewPsqlDB(ctx context.Context) (db *pgxpool.Pool, err error) {

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s pool_max_conns=%d search_path=%s", cfg.GetString("Postgresql.HOST"), cfg.GetInt("Postgresql.PORT"), cfg.GetString("Postgresql.USER"), cfg.GetString("Postgresql.PASS"), cfg.GetString("Postgresql.DEFAULT_DB"), cfg.GetInt("Postgresql.MAX_CONN"), cfg.GetString("Postgresql.DEFAULT_SCHEMA"))

	db, err = pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(ctx); err != nil {
		return nil, err
	}

	return
}
