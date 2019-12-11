package hooks

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/sirupsen/logrus"
)

func TestMongoLogHook(t *testing.T) {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
		defer cancelFunc()

		client.Disconnect(ctx)
	}()

	collection := client.Database("my_logs").Collection("logs")
	hook, err := NewMongoHook(collection)

	if err != nil {
		t.Fatal(err)
	}

	log.AddHook(hook)

	log.Info("Congratulations!")
}
