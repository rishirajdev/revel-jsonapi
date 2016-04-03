package model

// Chocolate is the chocolate that a user consumes in order to get fat and happy
type Comment struct {
	Id     int    `jsonapi:"primary,comments"`
	PostId int    `jsonapi:"attr,post_id"`
	Body   string `jsonapi:"attr,body"`
}