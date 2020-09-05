package bitrepo

import (
	"context"
	"errors"
	"fmt"
	"strings"
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
	client       *mongo.Client
	users        *mongo.Collection
	threads      *mongo.Collection
	posts        *mongo.Collection
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func getMongoDB(client *mongo.Client, config bitconfig.Config) *MongoDB {
	dbName := fmt.Sprintf("%s-%s", config.App, config.Env)
	db := client.Database(dbName)
	return &MongoDB{
		client:       client,
		users:        db.Collection("users"),
		threads:      db.Collection("threads"),
		posts:        db.Collection("posts"),
		ReadTimeout:  config.DBReadTimeout,
		WriteTimeout: config.DBWriteTimeout,
	}
}

// ConnectToMongoDB creates a mongo client and establishes a checked cnxn to the DB
func ConnectToMongoDB(host string, port int) (client *mongo.Client, err error) {
	client, err = getMongoDBClient(host, port)
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return
	}
	err = checkMongoConnection(client)

	return
}

func getMongoDBClient(host string, port int) (client *mongo.Client, err error) {
	uri := fmt.Sprintf("mongodb://%s:%d", host, port)
	client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	return
}

func checkMongoConnection(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return client.Ping(ctx, readpref.Primary())
}

//////////////////
//              //
// USER METHODS //
//              //
//////////////////

// GetUser returns the user associated with the given ID
func (db *MongoDB) GetUser(id string) (*bitdata.UserDocument, error) {
	var u *bitdata.UserDocument
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := db.users.FindOne(ctx, bson.M{"id": id}).Decode(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// GetUserByEmail returns a user document by email address
func (db *MongoDB) GetUserByEmail(email string) (*bitdata.UserDocument, error) {
	// TODO: implement
	return nil, errors.New("")
}

// CreateUser puts a new User document in the users collection
func (db *MongoDB) CreateUser(user *bitdata.UserDocument) (*mongo.InsertOneResult, error) {
	// first check that no record with the same ID already exists
	checkRes, err := db.GetUser(user.ID)
	if checkRes != nil {
		return nil, errors.New("CreateUser: Cannot create user, user with ID: " + user.ID + " already exists")
	} else if !strings.Contains(err.Error(), "no documents in result") {
		return nil, err
	}
	// now create new user in the database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := db.users.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UpdateUser updates User document in the users collection
// attributes that are updateable: Bio, Name, Email, Password
func (db *MongoDB) UpdateUser(user *bitdata.UserDocument) (*mongo.UpdateResult, error) {
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				bson.E{Key: "bio", Value: user.Bio},
				bson.E{Key: "name", Value: user.Name},
				bson.E{Key: "email", Value: user.Email},
				bson.E{Key: "password", Value: user.Password},
			},
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := db.users.UpdateOne(ctx, bson.M{"id": user.ID}, update)
	if err != nil {
		return nil, err
	} else if res.MatchedCount == 0 {
		return nil, errors.New("UpdateUser: no document was found with id: " + user.ID)
	}
	return res, nil
}

// DeleteUser deletes a User document from the users collection
func (db *MongoDB) DeleteUser(user *bitdata.UserDocument) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := db.users.DeleteOne(ctx, bson.M{"id": user.ID})
	if err != nil {
		return nil, err
	}
	return res, nil
}

////////////////////
//                //
// THREAD METHODS //
//                //
////////////////////

// GetThread returns the thread associated with the given ID
func (db *MongoDB) GetThread(id string) (*bitdata.ThreadDocument, error) {
	var t *bitdata.ThreadDocument
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := db.threads.FindOne(ctx, bson.M{"id": id}).Decode(t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// CreateThread puts a new Thread document in the threads collection
func (db *MongoDB) CreateThread(thread *bitdata.ThreadDocument) (*mongo.InsertOneResult, error) {
	// first check that no record with the same ID already exists
	checkRes, err := db.GetThread(thread.ID)
	if checkRes != nil {
		return nil, errors.New("CreateUser: Cannot create thread, thread with ID: " + thread.ID + " already exists")
	} else if !strings.Contains(err.Error(), "no documents in result") {
		return nil, err
	}
	// now create new user in the database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := db.threads.InsertOne(ctx, thread)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UpdateThread updates Thread document in the threads collection
// attributes that are updateable: Title, Topic
func (db *MongoDB) UpdateThread(thread *bitdata.ThreadDocument) (*mongo.UpdateResult, error) {
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				bson.E{Key: "topic", Value: thread.Topic},
				bson.E{Key: "title", Value: thread.Title},
			},
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := db.threads.UpdateOne(ctx, bson.M{"id": thread.ID}, update)
	if err != nil {
		return nil, err
	} else if res.MatchedCount == 0 {
		return nil, errors.New("UpdateThread: no document was found with id: " + thread.ID)
	}
	return res, nil
}

// DeleteThread deletes a Thread document from the thread collection
func (db *MongoDB) DeleteThread(thread *bitdata.ThreadDocument) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := db.threads.DeleteOne(ctx, bson.M{"id": thread.ID})
	if err != nil {
		return nil, err
	}
	return res, nil
}
