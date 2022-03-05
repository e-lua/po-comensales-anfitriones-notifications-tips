package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*------------------------BASIC DATA FOR SEARCH------------------------*/

type Mo_Tips struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	URLimg         string             `bson:"urlimg" json:"urlimg,omitempty"`
	DateRegistered time.Time          `bson:"dateregistered" json:"dateregistered,omitempty"`
	TypeUser       int                `bson:"typeuser" json:"typeuser,omitempty"`
	Typetip        int                `bson:"typetip" json:"typetip,omitempty"`
}

type Mo_Notifications struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Message        string             `bson:"message" json:"message,omitempty"`
	DateRegistered time.Time          `bson:"dateregistered" json:"dateregistered,omitempty"`
	IDUser         int                `bson:"iduser" json:"iduser,omitempty"`
	WasView        bool               `bson:"wasview" json:"wasview,omitempty"`
	TypeUser       int                `bson:"typeuser" json:"typeuser,omitempty"`
}

type Mo_NotificationShow struct {
	Message  string `bson:"message" json:"message,omitempty"`
	IDUser   int    `bson:"iduser" json:"iduser,omitempty"`
	TypeUser int    `bson:"typeuser" json:"typeuser,omitempty"`
}

type Mo_TipsShow struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	URLimg  string             `bson:"urlimg" json:"urlimg,omitempty"`
	Typetip int                `bson:"typetip" json:"typetip,omitempty"`
}
