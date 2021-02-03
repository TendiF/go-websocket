package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	CreatedBy primitive.ObjectID `bson:"created_by,omitempty"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty"`
	Message string `bson:"message,omitempty"`
}

type Group struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	CreatedBy primitive.ObjectID `bson:"created_by,omitempty"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty"`
	Title string `bson:"title,omitempty"`
}

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	CreatedBy primitive.ObjectID `bson:"created_by,omitempty"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty"`
	Name string `bson:"name,omitempty"`
	Phone string `bson:"phone,omitempty"`
}

type MapUserGroup struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	CreatedBy primitive.ObjectID `bson:"created_by,omitempty"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty"`
	IdUser primitive.ObjectID `bson:"id_user,omitempty"`
	IdGroup primitive.ObjectID `bson:"id_group,omitempty"`
}
