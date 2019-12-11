# logrus\hooks

Custom hooks for Sirupsen/logrus

## Usage

```go
import "github.com/vitpelekhaty/logrus/hooks"
```

### Example

```go
import (
    "context"
    "testing"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/sirupsen/logrus"
    "github.com/vitpelekhaty/logrus/hooks"
)

func main() {
    logger := logrus.New()
    logger.SetFormatter(&logrus.JSONFormatter{})

    client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

    if err != nil {
        logger.Fatal(err)
    }

    defer func() {
        ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
        defer cancelFunc()

        client.Disconnect(ctx)
    }()

    collection := client.Database("my_logs").Collection("logs")
    hook, err := hooks.NewMongoHook(collection)

    if err != nil {
        logger.Fatal(err)
    }

    logger.AddHook(hook)
    log.Info("Congratulations!")
}
```
