package models

type Address struct {
	Street  string `json:"street" bson:"street"`
	City    string `json:"city" bson:"city"`
	ZipCode string `json:"zipcode" bson:"zipcode"`
}

type User struct {
	Name    string  `json:"name" bson:"username"`
	Age     uint8   `json:"age" bson:"userage"`
	Address Address `json:"address" bson:"useraddress"`
}
