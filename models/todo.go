package models

type Todo struct {
	Id        int    `json:"id" bson:"id"`
	UserId    int    `json:"userId" bson:"userId,omitempty"`
	Title     string `json:"title" bson:"title"`
	Completed bool   `json:"completed" bson:"completed"`
}
