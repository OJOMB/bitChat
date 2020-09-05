package bitdata

import "time"

// UserDocument represents the data model for a bitChat user as it exists in the database
type UserDocument struct {
	ID        string
	Bio       string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}
