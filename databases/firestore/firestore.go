package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	cfg "github.com/spf13/viper"
	"google.golang.org/api/option"
)

// NewFireStoreDB Return new FirestoreDB client
func NewFireStoreDB() (db *firestore.Client, err error) {

	connStr := fmt.Sprintf("%s", cfg.GetString("Firestore.CREDENTIALS_PATH"))

	opt := option.WithCredentialsFile(connStr)
	db, err = firestore.NewClient(context.Background(), cfg.GetString("Firestore.PROJECT_ID"), opt)
	if err != nil {
		return nil, err
	}

	return
}
