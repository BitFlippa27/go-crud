package models

type Todo struct {
	Id        int    `json:"id" bson:"id"`
	Title     string `json:"title" bson:"title"`
	Completed bool   `json:"completed" bson:"completed"`
}
