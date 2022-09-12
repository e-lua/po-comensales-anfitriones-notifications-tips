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
	ViewBusiness   []int              `bson:"viewbusiness" json:"viewbusiness,"`
}

type Mo_Notifications struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title          string             `bson:"title" json:"title,omitempty"`
	Message        string             `bson:"message" json:"message,omitempty"`
	DateRegistered time.Time          `bson:"dateregistered" json:"dateregistered,omitempty"`
	IDUser         int                `bson:"iduser" json:"iduser,omitempty"`
	MultipleUser   []int              `bson:"multipleuser" json:"multipleuser,omitempty"`
	Priority       int                `bson:"priority" json:"priority"`
	WasView        bool               `bson:"wasview" json:"wasview,omitempty"`
	TypeUser       int                `bson:"typeuser" json:"typeuser,omitempty"`
}

type Mo_NotificationShow struct {
	Message        string    `bson:"message" json:"message,omitempty"`
	IDUser         int       `bson:"iduser" json:"iduser,omitempty"`
	WasView        bool      `bson:"wasview" json:"wasview"`
	DateRegistered time.Time `bson:"dateregistered" json:"dateregistered,omitempty"`
}

type Mo_TipsShow struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	URLimg  string             `bson:"urlimg" json:"urlimg,omitempty"`
	Typetip int                `bson:"typetip" json:"typetip,omitempty"`
}

type Mo_Anfitrion_PlanBusiness struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	IdBusiness     int                `bson:"idbusiness" json:"idbusiness"`
	DateRegistered time.Time          `bson:"dateregistered" json:"dateregistered"`
}
