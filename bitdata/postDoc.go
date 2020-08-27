package bitdata

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// PostDocument represents the data model for a bitChat post as it exists in the database
type PostDocument struct {
	id           string
	body         string
	threadId     string
	createdBy    string
	createdAt    time.Time
	inResponseTo string // post id
}

// ToBson converts the thread document to a BSON document
func (pd *PostDocument) ToBson() bson.D {
	return bson.D{
		{Key: "id", Value: pd.id},
		{Key: "body", Value: pd.body},
		{Key: "threadId", Value: pd.threadId},
		{Key: "createdBy", Value: pd.createdBy},
		{Key: "createdAt", Value: pd.createdAt},
		{Key: "inResponseTo", Value: pd.inResponseTo},
	}
}
