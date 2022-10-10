package main

import (
	"time"
)

type User struct {
	Name     string
	Password string
	tweets   []tweet
	Likes    []*tweet
}

type tweet struct {
	author       *User
	ID           int
	Content      string
	PostDateTime string
	Likes        []*User
	comments     []*tweet
	parent       *tweet
	status       bool
}

//When a user created an account
func CreateAccount(name, password string) User {
	var u User
	u.Name = name
	u.Password = password
	return u
}

//When user u posts a tweet
func (u *User) PostTweet(content string) {
	var newt tweet
	newt.Content = content
	newt.PostDateTime = time.Now().String()
	newt.ID = len((*u).tweets)
	newt.status = true
	newt.author = u
	newt.parent = nil
	(*u).tweets = append((*u).tweets, newt)
}

//When user u comments a tweet
func (u *User) CommentTweet(content string, t *tweet) {
	var newc tweet
	newc.author = u
	newc.Content = content
	newc.PostDateTime = time.Now().String()
	newc.ID = len((*u).tweets)
	newc.status = true
	newc.author = u
	newc.parent = t
	(*u).tweets = append((*u).tweets, newc)
	a := &((*u).tweets[newc.ID])
	t.comments = append(t.comments, a)
}

//When user u deletes a tweet/comment
func (u *User) DeleteTweet(t *tweet) {
	if u == (*t).author {
		(*u).tweets[(*t).ID].status = false
	} else {
		panic("Cannot delete")
	}
}

//When user u likes a tweet/commet
func (u *User) LikeTweet(t *tweet) {
	u.Likes = append(u.Likes, t)
	(*t).Likes = append((*t).Likes, u)
}
