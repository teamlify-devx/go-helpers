package scylla

import (
	"strings"
	"time"

	"github.com/gocql/gocql"
	cfg "github.com/spf13/viper"
)

// ScylaDb consistency level constants
const (
	ConsistencyAny         = gocql.Any
	ConsistencyOne         = gocql.One
	ConsistencyTwo         = gocql.Two
	ConsistencyThree       = gocql.Three
	ConsistencyQuorum      = gocql.Quorum
	ConsistencyAll         = gocql.All
	ConsistencyLocalQuorum = gocql.LocalQuorum
	ConsistencyEachQuorum  = gocql.EachQuorum
	ConsistencyLocalOne    = gocql.LocalOne
)

// NewScyllaDB Return new NewScyllaDB instance
func NewScyllaDB() (db *gocql.Session, err error) {
	retryPolicy := &gocql.ExponentialBackoffRetryPolicy{
		Min:        time.Second,
		Max:        10 * time.Second,
		NumRetries: cfg.GetInt("Scylla.RETRIES"),
	}

	connStr := strings.Split(cfg.GetString("Scylla.HOSTS"), ",")

	cluster := gocql.NewCluster(connStr...)
	cluster.Consistency = gocql.ParseConsistency(cfg.GetString("Scylla.CONSISTENCY"))
	cluster.Keyspace = cfg.GetString("Scylla.KEYSPACE")
	cluster.Timeout = time.Duration(cfg.GetInt("Scylla.TIMEOUT")) * time.Second
	cluster.RetryPolicy = retryPolicy
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())

	db, err = cluster.CreateSession()

	return
}
