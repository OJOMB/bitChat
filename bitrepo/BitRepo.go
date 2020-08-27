package bitrepo

import (
	"github.com/OJOMB/bitChat/bitdata"
)

// BitRepo declares the method set required by a bitChat database (repo) object
type BitRepo interface {
	CreateThread(thread *bitdata.ThreadDocument) error
	UpdateThread(thread *bitdata.ThreadDocument) error
	DeleteThread(threadID string) error
	GetThread(threadID string) (*bitdata.ThreadDocument, error)
	GetAllThreadPosts(postIds []string) ([]bitdata.PostDocument, error)

	CreatePost(post *bitdata.PostDocument) error
	UpdatePost(post *bitdata.PostDocument) error
	DeletePost(postID string) error
	GetPost(postID string) (*bitdata.PostDocument, error)

	CreateUser(user *bitdata.UserDocument) error
	UpdateUser(user *bitdata.UserDocument) error
	DeleteUser(userID string) error
	GetUser(userID string) (*bitdata.UserDocument, error)
}
