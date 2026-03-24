package valkey

import (
	"crypto/tls"
	cfg "github.com/spf13/viper"
	"github.com/valkey-io/valkey-go"
)

// NewValkeyClient Returns new valkey client
func NewValkeyClient(dbNum int) (db valkey.Client, err error) {

	connStr := cfg.GetStringSlice("Valkey.HOSTS")

	options := valkey.ClientOption{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		InitAddress:           connStr,
		SelectDB:              dbNum,
		DisableCache:          cfg.GetBool("Valkey.DISABLE_CACHE"),
		DisableAutoPipelining: cfg.GetBool("Valkey.DISABLE_AUTO_PIPELINING"),
		ReplicaOnly:           cfg.GetBool("Valkey.DISABLE_REPLICA"),
	}

	if cfg.GetString("Valkey.USERNAME") != "" {
		options.Username = cfg.GetString("Valkey.USERNAME")
	}

	if cfg.GetString("Valkey.PASSWORD") != "" {
		options.Password = cfg.GetString("Valkey.PASSWORD")
	}

	db, err = valkey.NewClient(options)

	db, err = valkey.NewClient(valkey.ClientOption{InitAddress: connStr})
	if err != nil {
		return nil, err
	}

	return
}
