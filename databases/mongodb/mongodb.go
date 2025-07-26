package mongodb

import (
	"context"
	"fmt"
	cfg "github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewMongoDB Return new MongoDB client
func NewMongoDB(ctx context.Context) (db *mongo.Client, err error) {

	connStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority&authSource=admin", cfg.GetString("MongoDB.USER"), cfg.GetString("MongoDB.PASS"), cfg.GetString("MongoDB.HOST"), cfg.GetString("MongoDB.DEFAULT_DB"))

	db, err = mongo.Connect(ctx, options.Client().ApplyURI(connStr))
	if err != nil {
		return nil, err
	}

	return
}
