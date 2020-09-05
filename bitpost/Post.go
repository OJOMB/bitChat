package bitpost

import (
	"time"

	"github.com/OJOMB/bitChat/bitdata"
)

// Post is the data model for User posts on bitChat
type Post struct {
	ID           string
	Body         string
	ThreadID     string
	CreatedBy    string
	CreatedAt    time.Time
	InResponseTo string // post id
}

// GenerateDummyPost returns a dummy Post for testing purposes
func GenerateDummyPost() *Post {
	t, _ := time.Parse(time.RFC3339, "2020-08-29T15:32:11+0000")
	return &Post{
		ID:           "f8b6c091-d37c-4bb3-9418-805569216358",
		Body:         "this is the content of my dummy post",
		ThreadID:     "7bc4215f-91e4-4768-9ac3-05f93fefbcbd",
		CreatedBy:    "b916bfc4-4a6f-4559-9cfa-213f7a9c3b73",
		CreatedAt:    t,
		InResponseTo: "",
	}
}

// CreatedAtUTC returns formatted string of created at datetime
func (p *Post) CreatedAtUTC() string {
	return p.CreatedAt.UTC().Format("2006-01-02T15:04:05-0700")
}

func (p *Post) ToDocument() *bitdata.PostDocument {
	return &bitdata.PostDocument{
		ID:           p.ID,
		Body:         p.Body,
		ThreadID:     p.ThreadID,
		CreatedBy:    p.CreatedBy,
		CreatedAt:    p.CreatedAt,
		InResponseTo: p.InResponseTo,
	}
}
