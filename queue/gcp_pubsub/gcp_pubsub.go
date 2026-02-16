package gcp_pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"errors"
	cfg "github.com/spf13/viper"
	"google.golang.org/api/option"
)

func NewPubSubClient(ctx context.Context) (client *pubsub.Client, err error) {

	opt := option.WithCredentialsFile(cfg.GetString("PubSub.CONFIG_PATH"))

	client, err = pubsub.NewClient(ctx, cfg.GetString("PubSub.PROJECT_ID"), opt)
	if err != nil {
		return nil, errors.New("Failed to init GCP Pub/Sub client")
	}

	return
}
