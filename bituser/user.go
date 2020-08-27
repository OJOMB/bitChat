package bituser

import (
	"time"

	"github.com/rivo/uniseg"	

	"github.com/OJOMB/bitChat/bitpost"
	"github.com/OJOMB/bitChat/bitrepo"
	"github.com/OJOMB/bitChat/bitthread"
)

type User struct {
	ID        string
	Bio       string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time 
}

// GenerateDummyUser generates a dummy user for testing purposes
func GenerateDummyUser() *User{
	return &User{
		ID: "b916bfc4-4a6f-4559-9cfa-213f7a9c3b73",
		Bio: "Hey there I'm a user",
		Name: "Oscar",
		Email: "oscarm-b@tutamail.com",
		Password: "password",
		CreatedAt: time.Parse(time.iso8601,"2020-08-27T15:32:11+0000"),
}

// Creates a new thread in in the given db
func (u *User) CreateThread(db *bitrepo.BitRepo, topic string, title string) (t *bitthread.Thread, err error) {
	tid := bitUtils.CreateUUID()
	t = &bitthread.Thread{
		Id:        tid,
		Topic:     topic,
		Title: 	   title,
		CreatedBy: u.Id,
		CreatedAt: time.Now()
	}
	err := db.createThread(t)
	return
}

// Creates a new post in an existing thread in the given db
// param pid refers to the post id of the post to which the new post is a response
func (u *User) CreatePost(db *bitrepo.BitRepo, body string, tid string, pid string) (p *bitpost.Post, err error) {
	pid := bitUtils.CreateUUID()
	p = &bitpost.Post{
		Id:           pid,
		Body:         body,
		ThreadId:     tid,
		CreatedBy:    u.Id,
		CreatedAt:    time.Now(),
		InResponseTo: pid,
	}
	err := db.createPost(p)
	return
}

// Creates a thread and the initial post in the given db
func (u *User) CreateThreadPost(db *bitrepo.BitRepo, body string, topic string, title string) (t *bitthread.Thread, p *bitpost.Post, err error) {
	t, err := u.CreateThread(db, topic, title)
	if err != nil {
		return nil, nil, err
	}
	p, err := u.CreatePost(db, body, t.Id, nil)
	if err != nil {
		return nil, nil, err
	}
	return
}

// 
func (u *User) DeleteThread(db *bitrepo.BitRepo, tid string) error {
	t := db.GetThread(tid)
	if u.Id != t.CreatedBy {
		return errors.New(
			fmt.Sprintf("User: %s cannot delete thread: %s as they are not the thread owner", u.Id, tid)
		)
	}
	return db.DeleteThread(tid)
}

func (u *User) DeletePost(db *bitrepo.BitRepo, pid string) error {
	p := db.GetPost(pid)
	if u.Id != p.CreatedBy {
		return errors.New(
			fmt.Sprintf("User: %s cannot delete post: %s as they are not the post owner", u.Id, tid)
		)
	}
	return db.DeletePost(pid)
}

func (u *User) Delete(db *bitrepo.BitRepo) error {
	return db.DeleteUser(u.Id)
}

func (u *User) UpdateEmail(db *bitrepo.BitRepo, newEmail) error {
	u.Email = newEmail
	return db.UpdateUser(u)
}

func (u *User) UpdateName(db *bitrepo.BitRepo, newName) error {
	u.Name = newName
	return db.UpdateUser(u)
}

func (u *User) UpdatePassword(db *bitrepo.BitRepo, newPassword) error {
	u.Password = newPassword
	return db.UpdateUser(u)
}

func (u *User) UpdateBio(db *bitrepo.BitRepo, newBio) error {
	chars := uniseg.GraphemeClusterCount(newBio)
	if chars >= 250 {
		return errors.New("Bio must be less than 250 characters")
	}
	u.Bio = newBio
	return db.UpdateUser(u)
}