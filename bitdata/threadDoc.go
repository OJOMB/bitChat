package bitdata

import (
	"go.mongodb.org/mongo-driver/bson"
)

// ThreadDocument represents the data model for a bitChat thread as it exists in the database
type ThreadDocument struct {
	Id        string
	Topic     string
	Title     string
	CreatedBy string
	CreatedAt string
}

// ToBson converts the thread document to a BSON document
func (td ThreadDocument) ToBson() bson.D {
	return bson.D{
		{Key: "id", Value: td.Id},
		{Key: "topic", Value: td.Topic},
		{Key: "title", Value: td.Title},
		{Key: "createdBy", Value: td.CreatedBy},
		{Key: "createdAt", Value: td.CreatedAt},
	}
}

// ThreadDocumentFromBson returns a ThreadDocument constructed from a BSON document
func ThreadDocumentFromBson(bsonMap bson.M) ThreadDocument {
	return ThreadDocument{
		Id:        bsonMap["id"].(string),
		Topic:     bsonMap["topic"].(string),
		Title:     bsonMap["title"].(string),
		CreatedBy: bsonMap["createdBy"].(string),
		CreatedAt: bsonMap["createdAt"].(string),
	}
}
