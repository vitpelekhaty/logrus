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
	timeout *time.Duration
}

// MongoHookOption mongo hook option
type MongoHookOption func(*MongoHook)

func WithTimeout(timeout *time.Duration) MongoHookOption {
	return func(hook *MongoHook) {
		hook.timeout = timeout
	}
}

// NewMongoHook returns new hook
func NewMongoHook(collection *mongo.Collection, opts ...MongoHookOption) (*MongoHook, error) {
	if collection == nil {
		return nil, errors.New("undefined mongo collection")
	}

	hook := &MongoHook{collection: collection}

	for _, opt := range opts {
		opt(hook)
	}

	return hook, nil
}

// Fire hook
func (hook *MongoHook) Fire(entry *logrus.Entry) error {
	ctx, cancel := hook.context(entry)

	defer func() {
		if cancel != nil {
			cancel()
		}
	}()

	s, _ := entry.String()

	document := bson.M{
		"time": entry.Time,
		"level": entry.Level.String(),
		"message": entry.Message,
		"entry": s,
	}

	if len(entry.Data) > 0 {
		for key, value := range entry.Data {
			document[key] = value
		}
	}

	_, err := hook.collection.InsertOne(ctx, document)

	return err
}

func (hook *MongoHook) context(entry *logrus.Entry) (context.Context, context.CancelFunc) {
	parentContext := entry.Context

	if parentContext == nil {
		parentContext = context.Background()
	}

	if hook.timeout != nil {
		return context.WithTimeout(parentContext, *hook.timeout)
	}

	return parentContext, nil
}

func (hook *MongoHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
