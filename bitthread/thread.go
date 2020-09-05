package bitthread

import (
	"time"

	"github.com/OJOMB/bitChat/bitdata"
)

// Thread represents a BitChat thread business object
type Thread struct {
	ID        string
	Topic     string
	Title     string
	CreatedBy string
	CreatedAt time.Time
}

// GenerateDummyThread returns a dummy thread for testing purposes
func GenerateDummyThread() *Thread {
	t, _ := time.Parse(time.RFC3339, "2020-08-29T15:32:11+0000")
	return &Thread{
		ID:        "7bc4215f-91e4-4768-9ac3-05f93fefbcbd",
		Topic:     "topic",
		Title:     "title",
		CreatedBy: "b916bfc4-4a6f-4559-9cfa-213f7a9c3b73",
		CreatedAt: t,
	}
}

// CreatedAtUTC returns the CreatedAt datetime in RFC 3339 format
func (t *Thread) CreatedAtUTC() string {
	return t.CreatedAt.UTC().Format("2006-01-02T15:04:05-0700")
}

// ToDocument converts the receiver to a Post Document
func (t *Thread) ToDocument() *bitdata.ThreadDocument {
	return &bitdata.ThreadDocument{
		ID:        t.ID,
		Topic:     t.Topic,
		Title:     t.Title,
		CreatedBy: t.CreatedBy,
		CreatedAt: t.CreatedAt,
	}
}

// ThreadFromDocument returns a thread business object based on the given Thread document
func ThreadFromDocument(doc *bitdata.ThreadDocument) *Thread {
	return &Thread{
		ID:        doc.ID,
		Topic:     doc.Topic,
		Title:     doc.Title,
		CreatedBy: doc.CreatedBy,
		CreatedAt: doc.CreatedAt,
	}
}
