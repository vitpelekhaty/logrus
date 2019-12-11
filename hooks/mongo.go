package hooks

import (
	"context"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoHook logrus hook for writing logs into mongo collection
type MongoHook struct {
	collection *mongo.Collection
}

// NewMongoHook returns new hook
func NewMongoHook(collection *mongo.Collection) (*MongoHook, error) {
	if collection == nil {
		return nil, errors.New("undefined mongo collection")
	}

	return &MongoHook{collection: collection}, nil
}

// Fire hook
func (hook *MongoHook) Fire(entry *logrus.Entry) error {
	data, err := entry.String()

	if err != nil {
		return err
	}

	var parentContext = entry.Context

	if entry.Context == nil {
		parentContext = context.Background()
	}

	ctx, cancelFunc := context.WithTimeout(parentContext, time.Second*10)
	defer cancelFunc()

	_, err = hook.collection.InsertOne(ctx, bson.M{"entry": data})

	return err
}

func (hook *MongoHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
