package models

type User struct {
	ID       int    `bson:"_id" json:"id"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"-" json:"-"`
}
