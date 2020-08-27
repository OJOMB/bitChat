package bitdata

import "time"

// UserDocument represents the data model for a bitChat user as it exists in the database
type UserDocument struct {
	id        string
	bio       string
	name      string
	email     string
	password  string
	createdAt time.Time
}
