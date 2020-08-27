package bitrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/OJOMB/bitChat/bitconfig"
	"github.com/OJOMB/bitChat/bitdata"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDB implementation of bitChat repo
type MongoDB struct {
	client  *mongo.Client
	users   *mongo.Collection
	threads *mongo.Collection
	posts   *mongo.Collection
}

func getMongoDB(client *mongo.Client, config bitconfig.Config) *MongoDB {
	dbName := fmt.Sprintf("%s-%s", config.App, config.Env)
	db := client.Database(dbName)
	return &MongoDB{
		client:  client,
		users:   db.Collection("users"),
		threads: db.Collection("threads"),
		posts:   db.Collection("posts"),
	}
}

func connectToMongoDB(ip string, port int) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%d", ip, port)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	if err != nil {
		return nil, err
	}

	// Check the connection
	if checkMongoConnection(client) != nil {
		return nil, err
	}

	return client, nil
}

func checkMongoConnection(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return client.Ping(ctx, readpref.Primary())
}

////////////////////
// THREAD METHODS //
////////////////////

// CreateThread puts a new thread document in the threads collection
func (db *MongoDB) CreateThread(td *bitdata.ThreadDocument) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := db.threads.InsertOne(ctx, td.ToBson())
	if err != nil {
		return err
	}
	return nil
}

// GetAllThreads returns all threads documents in the threads collection
func (db *MongoDB) GetAllThreads() ([]bitdata.ThreadDocument, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	cursor, err := db.threads.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var threads []bitdata.ThreadDocument
	for cursor.Next(ctx) {
		var thread bson.M
		if err = cursor.Decode(&thread); err != nil {
			return nil, err
		}
		threads = append(
			threads,
			bitdata.ThreadDocumentFromBson(thread),
		)
	}
	return threads, nil
}

// DeleteThread removes the given thread document from the database
func (db *MongoDB) DeleteThread(tid string) error {}

func (db *MongoDB) createPost(pd *bitdata.PostDocument) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := db.posts.InsertOne(ctx, pd.ToBson())
	if err != nil {
		return err
	}
}
