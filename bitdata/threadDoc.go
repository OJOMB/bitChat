package bitdata

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// ThreadDocument represents the data model for a bitChat thread as it exists in the database
type ThreadDocument struct {
	ID        string
	Topic     string
	Title     string
	CreatedBy string
	CreatedAt time.Time
}

// ToBson converts the thread document to a BSON document
func (td ThreadDocument) ToBson() bson.D {
	return bson.D{
		{Key: "id", Value: td.ID},
		{Key: "topic", Value: td.Topic},
		{Key: "title", Value: td.Title},
		{Key: "createdBy", Value: td.CreatedBy},
		{Key: "createdAt", Value: td.CreatedAt},
	}
}
