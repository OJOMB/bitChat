package bituser

import (
	"errors"
	"fmt"
	"time"

	"github.com/rivo/uniseg"

	"github.com/OJOMB/bitChat/bitdata"
	"github.com/OJOMB/bitChat/bitpost"
	"github.com/OJOMB/bitChat/bitrepo"
	"github.com/OJOMB/bitChat/bitthread"
	"github.com/OJOMB/bitChat/bitutils"
)

// User models a bitChat user
type User struct {
	ID        string
	Bio       string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

// GenerateDummyUser generates a dummy user for testing purposes
func GenerateDummyUser() *User {
	t, _ := time.Parse(time.RFC3339, "2020-08-29T15:32:11+0000")
	return &User{
		ID:        "b916bfc4-4a6f-4559-9cfa-213f7a9c3b73",
		Bio:       "Hey there I'm a user",
		Name:      "Oscar",
		Email:     "oscarm-b@tutamail.com",
		Password:  "password",
		CreatedAt: t,
	}
}

// ToDocument returns a User Document from the user data
func (u *User) ToDocument() *bitdata.UserDocument {
	return &bitdata.UserDocument{
		ID:        u.ID,
		Bio:       u.Bio,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
	}
}

// CreateThread  Creates a new thread in in the given db
func (u *User) CreateThread(db bitrepo.BitRepo, topic, title string) (t *bitthread.Thread, err error) {
	tid := bitutils.CreateUUID()
	t = &bitthread.Thread{
		ID:        tid,
		Topic:     topic,
		Title:     title,
		CreatedBy: u.ID,
		CreatedAt: time.Now(),
	}
	err = db.CreateThread(t.ToDocument())
	return
}

// CreatePost Creates a new post in an existing thread in the given db
// param pid refers to the post id of the post to which the new post is a response
func (u *User) CreatePost(db bitrepo.BitRepo, body string, tid string, pid string) (p *bitpost.Post, err error) {
	pid = bitutils.CreateUUID()
	p = &bitpost.Post{
		ID:           pid,
		Body:         body,
		ThreadID:     tid,
		CreatedBy:    u.ID,
		CreatedAt:    time.Now(),
		InResponseTo: pid,
	}
	err = db.CreatePost(p.ToDocument())
	return
}

// CreateThreadPost Creates a thread and the initial post in the given db
func (u *User) CreateThreadPost(db bitrepo.BitRepo, body string, topic string, title string) (t *bitthread.Thread, p *bitpost.Post, err error) {
	t, err = u.CreateThread(db, topic, title)
	if err != nil {
		return nil, nil, err
	}
	p, err = u.CreatePost(db, body, t.ID, "")
	if err != nil {
		return nil, nil, err
	}
	return
}

// DeleteThread deletes a thread owned by the receiver user
func (u *User) DeleteThread(db bitrepo.BitRepo, tid string) error {
	t, err := db.GetThread(tid)
	if err != nil {
		return err
	}
	if u.ID != t.CreatedBy {
		return fmt.Errorf("User: %s cannot delete thread: %s as they are not the thread owner", u.ID, tid)
	}
	return db.DeleteThread(tid)
}

// DeletePost deletes a post owned by the receiver user
func (u *User) DeletePost(db bitrepo.BitRepo, pid string) error {
	p, err := db.GetPost(pid)
	if err != nil {
		return err
	}
	if u.ID != p.CreatedBy {
		return fmt.Errorf("User: %s cannot delete post: %s as they are not the post owner", u.ID, pid)
	}
	return db.DeletePost(pid)
}

// Delete deletes the receiver user from the database
func (u *User) Delete(db bitrepo.BitRepo) error {
	return db.DeleteUser(u.ID)
}

// UpdateEmail updates the users email address in the database
func (u *User) UpdateEmail(db bitrepo.BitRepo, newEmail string) error {
	u.Email = newEmail
	return db.UpdateUser(u.ToDocument())
}

//UpdateName updates the users name in the database
func (u *User) UpdateName(db bitrepo.BitRepo, newName string) error {
	u.Name = newName
	return db.UpdateUser(u.ToDocument())
}

// UpdatePassword updates the users password in the database
func (u *User) UpdatePassword(db bitrepo.BitRepo, newPassword string) error {
	u.Password = newPassword
	return db.UpdateUser(u.ToDocument())
}

// UpdateBio updates the users bio in the database
func (u *User) UpdateBio(db bitrepo.BitRepo, newBio string) error {
	chars := uniseg.GraphemeClusterCount(newBio)
	if chars >= 250 {
		return errors.New("Bio must be less than 250 characters")
	}
	u.Bio = newBio
	return db.UpdateUser(u.ToDocument())
}
