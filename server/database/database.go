package database


import (
	"context"
    "time"
 
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)


func ConnectToMongoDb (url string, timeOut int) (*mongo.Client, context.Context,
    context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
                                       time.Duration(timeOut) * time.Second)

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
    return client, ctx, cancel, err
}

func Ping(ctx context.Context, client *mongo.Client) error{
    if err := client.Ping(ctx, readpref.Primary()); err != nil {
        return err
    }
    return nil
}