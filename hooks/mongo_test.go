package hooks

import (
	"context"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LogFormatter byte

const (
	LogFormatterText LogFormatter = iota
	LogFormatterJSON
)

type test struct{
	formatter LogFormatter
	level logrus.Level
	message string
	data logrus.Fields
	hookOptions []MongoHookOption
}

const (
	mongoURI = "mongodb://localhost:27017"
	database = "test_hooks"
	collection = "logs"
)

var timeout = time.Second * 5

var cases = []test{
	{
		formatter: LogFormatterText,
		level: logrus.InfoLevel,
		message: "Info",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{},
	},
	{
		formatter: LogFormatterText,
		level: logrus.DebugLevel,
		message: "Debug",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{},
	},
	{
		formatter: LogFormatterText,
		level: logrus.ErrorLevel,
		message: "Error",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{},
	},
	{
		formatter: LogFormatterText,
		level: logrus.TraceLevel,
		message: "Trace",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{},
	},
	{
		formatter: LogFormatterText,
		level: logrus.WarnLevel,
		message: "Warn",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.InfoLevel,
		message: "Info",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.DebugLevel,
		message: "Debug",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.ErrorLevel,
		message: "Error",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.TraceLevel,
		message: "Trace",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.WarnLevel,
		message: "Warn",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{},
	},
	{
		formatter: LogFormatterText,
		level: logrus.InfoLevel,
		message: "Info",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{WithTimeout(nil)},
	},
	{
		formatter: LogFormatterText,
		level: logrus.DebugLevel,
		message: "Debug",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{WithTimeout(nil)},
	},
	{
		formatter: LogFormatterText,
		level: logrus.ErrorLevel,
		message: "Error",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{WithTimeout(nil)},
	},
	{
		formatter: LogFormatterText,
		level: logrus.TraceLevel,
		message: "Trace",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{WithTimeout(nil)},
	},
	{
		formatter: LogFormatterText,
		level: logrus.WarnLevel,
		message: "Warn",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{WithTimeout(nil)},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.InfoLevel,
		message: "Info",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{WithTimeout(nil)},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.DebugLevel,
		message: "Debug",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{WithTimeout(nil)},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.ErrorLevel,
		message: "Error",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{WithTimeout(nil)},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.TraceLevel,
		message: "Trace",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{WithTimeout(nil)},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.WarnLevel,
		message: "Warn",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{WithTimeout(nil)},
	},
	{
		formatter: LogFormatterText,
		level: logrus.InfoLevel,
		message: "Info",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{WithTimeout(&timeout)},
	},
	{
		formatter: LogFormatterText,
		level: logrus.DebugLevel,
		message: "Debug",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{WithTimeout(&timeout)},
	},
	{
		formatter: LogFormatterText,
		level: logrus.ErrorLevel,
		message: "Error",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{WithTimeout(&timeout)},
	},
	{
		formatter: LogFormatterText,
		level: logrus.TraceLevel,
		message: "Trace",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{WithTimeout(&timeout)},
	},
	{
		formatter: LogFormatterText,
		level: logrus.WarnLevel,
		message: "Warn",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{WithTimeout(&timeout)},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.InfoLevel,
		message: "Info",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{WithTimeout(&timeout)},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.DebugLevel,
		message: "Debug",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{WithTimeout(&timeout)},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.ErrorLevel,
		message: "Error",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{WithTimeout(&timeout)},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.TraceLevel,
		message: "Trace",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{WithTimeout(&timeout)},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.WarnLevel,
		message: "Warn",
		data: logrus.Fields{},
		hookOptions: []MongoHookOption{WithTimeout(&timeout)},
	},
	{
		formatter: LogFormatterText,
		level: logrus.InfoLevel,
		message: "Info",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{},
	},
	{
		formatter: LogFormatterText,
		level: logrus.DebugLevel,
		message: "Debug",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{},
	},
	{
		formatter: LogFormatterText,
		level: logrus.ErrorLevel,
		message: "Error",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{},
	},
	{
		formatter: LogFormatterText,
		level: logrus.TraceLevel,
		message: "Trace",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{},
	},
	{
		formatter: LogFormatterText,
		level: logrus.WarnLevel,
		message: "Warn",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.InfoLevel,
		message: "Info",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.DebugLevel,
		message: "Debug",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.ErrorLevel,
		message: "Error",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.TraceLevel,
		message: "Trace",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.WarnLevel,
		message: "Warn",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{},
	},
	{
		formatter: LogFormatterText,
		level: logrus.InfoLevel,
		message: "Info",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{WithTimeout(nil)},
	},
	{
		formatter: LogFormatterText,
		level: logrus.DebugLevel,
		message: "Debug",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{WithTimeout(nil)},
	},
	{
		formatter: LogFormatterText,
		level: logrus.ErrorLevel,
		message: "Error",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{WithTimeout(nil)},
	},
	{
		formatter: LogFormatterText,
		level: logrus.TraceLevel,
		message: "Trace",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{WithTimeout(nil)},
	},
	{
		formatter: LogFormatterText,
		level: logrus.WarnLevel,
		message: "Warn",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{WithTimeout(nil)},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.InfoLevel,
		message: "Info",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{WithTimeout(nil)},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.DebugLevel,
		message: "Debug",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{WithTimeout(nil)},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.ErrorLevel,
		message: "Error",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{WithTimeout(nil)},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.TraceLevel,
		message: "Trace",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{WithTimeout(nil)},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.WarnLevel,
		message: "Warn",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{WithTimeout(nil)},
	},
	{
		formatter: LogFormatterText,
		level: logrus.InfoLevel,
		message: "Info",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{WithTimeout(&timeout)},
	},
	{
		formatter: LogFormatterText,
		level: logrus.DebugLevel,
		message: "Debug",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{WithTimeout(&timeout)},
	},
	{
		formatter: LogFormatterText,
		level: logrus.ErrorLevel,
		message: "Error",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{WithTimeout(&timeout)},
	},
	{
		formatter: LogFormatterText,
		level: logrus.TraceLevel,
		message: "Trace",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{WithTimeout(&timeout)},
	},
	{
		formatter: LogFormatterText,
		level: logrus.WarnLevel,
		message: "Warn",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{WithTimeout(&timeout)},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.InfoLevel,
		message: "Info",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{WithTimeout(&timeout)},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.DebugLevel,
		message: "Debug",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{WithTimeout(&timeout)},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.ErrorLevel,
		message: "Error",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{WithTimeout(&timeout)},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.TraceLevel,
		message: "Trace",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{WithTimeout(&timeout)},
	},
	{
		formatter: LogFormatterJSON,
		level: logrus.WarnLevel,
		message: "Warn",
		data: logrus.Fields{"database": database, "collection": collection},
		hookOptions: []MongoHookOption{WithTimeout(&timeout)},
	},
}

func TestMongoLogHook(t *testing.T) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))

	if err != nil {
		t.Fatal(err)
	}

	defer func(){
		client.Disconnect(context.Background())
	}()

	col := client.Database(database).Collection(collection)

	for _, test := range cases {
		log := logrus.New()
		log.SetLevel(logrus.TraceLevel)

		if test.formatter == LogFormatterJSON {
			log.SetFormatter(&logrus.JSONFormatter{})
		} else {
			log.SetFormatter(&logrus.TextFormatter{})
		}

		hook, err := NewMongoHook(col, test.hookOptions...)

		if err != nil {
			t.Fatal(err)
		}

		log.AddHook(hook)

		switch test.level {
		case logrus.InfoLevel:
			if len(test.data) > 0 {
				log.WithFields(test.data).Info(test.message)
			} else {
				log.Info(test.message)
			}

		case logrus.DebugLevel:
			if len(test.data) > 0 {
				log.WithFields(test.data).Debug(test.message)
			} else {
				log.Debug(test.message)
			}

		case logrus.ErrorLevel:
			if len(test.data) > 0 {
				log.WithFields(test.data).Error(test.message)
			} else {
				log.Error(test.message)
			}

		case logrus.TraceLevel:
			if len(test.data) > 0 {
				log.WithFields(test.data).Trace(test.message)
			} else {
				log.Trace(test.message)
			}

		case logrus.WarnLevel:
			if len(test.data) > 0 {
				log.WithFields(test.data).Warn(test.message)
			} else {
				log.Warn(test.message)
			}

		default:
			if len(test.data) > 0 {
				log.WithFields(test.data).Info(test.message)
			} else {
				log.Info(test.message)
			}
		}
	}
}
