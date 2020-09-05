package bitdata

import (
	"time"
)

// PostDocument represents the data model for a bitChat post as it exists in the database
type PostDocument struct {
	ID           string
	Body         string
	ThreadID     string
	CreatedBy    string
	CreatedAt    time.Time
	InResponseTo string // post id
}
