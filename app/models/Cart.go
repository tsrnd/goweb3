package models

import (
	"time"
)
type Cart struct {
	Id int	     			`db:"id" bson:"id"`
	UserId int				`db:"user_id" bson:"user_id"`
	CreatedAt time.Time     `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `db:"updated_at" bson:"updated_at"`
	DeletedAt time.Time     `db:"deleted_at" bson:"deleted_at"`
}
