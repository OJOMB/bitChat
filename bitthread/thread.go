package bitthread

import (
	"time"

	"github.com/OJOMB/bitChat/bitdata"
	"github.com/OJOMB/bitChat/bitrepo"
)

// Thread represents a BitChat thread business object
type Thread struct {
	ID        int
	Topic     string
	Title     string
	CreatedBy string
	CreatedAt time.Time
}

// GenerateDummyThread returns a dummy thread for testing purposes
func GenerateDummyThread() *Thread {
	return &Thread{
		ID:        "7bc4215f-91e4-4768-9ac3-05f93fefbcbd",
		Topic:     "topic",
		Title:     "title",
		CreatedBy: "b916bfc4-4a6f-4559-9cfa-213f7a9c3b73",
		CreatedAt: time.Parse(time.iso8601, "2020-08-29T15:32:11+0000"),
	}
}

// CreatedAtUTC returns the CreatedAt datetime in RFC 3339 format
func (t *Thread) CreatedAtUTC() string {
	return t.CreatedAt.UTC().Format("2006-01-02T15:04:05-0700")
}

// get posts to a thread
func (t *Thread) getAllPosts(db bitrepo.BitRepo) ([]Post, error) {
	return db.getAllThreadPosts(t.ID)
}

// ToDocument converts the receiver to a Post Document
func (t *Thread) ToDocument() *bitdata.ThreadDocument {
	return &bitdata.ThreadDocument{
		id:        t.ID,
		topic:     t.Topic,
		title:     t.Title,
		CreatedBy: t.CreatedBy,
		CreatedAt: t.CreatedAtUTC(),
	}
}

// ThreadFromDocument returns a thread business object based on the given Thread document
func ThreadFromDocument(doc *bitdata.ThreadDocument) *Thread {
	return &Thread{
		ID:        doc.id,
		topic:     doc.topic,
		title:     doc.title,
		CreatedBy: doc.createdBy,
		CreatedAt: doc.createdAt,
	}
}
