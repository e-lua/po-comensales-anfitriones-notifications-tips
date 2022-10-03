package models

import (
	"time"
)

/*------------------------BASIC DATA FOR SEARCH------------------------*/

type Pg_Tips struct {
	ID             string    `bson:"_id" json:"id"`
	URLimg         string    `bson:"urlimg" json:"urlimg"`
	DateRegistered time.Time `bson:"dateregistered" json:"dateregistered"`
	TypeUser       int       `bson:"typeuser" json:"typeuser"`
	Typetip        int       `bson:"typetip" json:"typetip"`
	ViewBusiness   []int     `bson:"viewbusiness" json:"viewbusiness,"`
}

type Pg_Notifications struct {
	ID             string    `bson:"_id" json:"id"`
	Title          string    `bson:"title" json:"title"`
	Message        string    `bson:"message" json:"message"`
	DateRegistered time.Time `bson:"dateregistered" json:"dateregistered"`
	IDUser         int       `bson:"iduser" json:"iduser"`
	MultipleUser   []int     `bson:"multipleuser" json:"multipleuser"`
	Priority       int       `bson:"priority" json:"priority"`
	WasView        bool      `bson:"wasview" json:"wasview"`
	TypeUser       int       `bson:"typeuser" json:"typeuser"`
	CodeNotify     int       `bson:"codenotify" json:"codenotify"`
}

type Pg_NotificationShow struct {
	Message        string    `bson:"message" json:"message"`
	IDUser         int       `bson:"iduser" json:"iduser"`
	WasView        bool      `bson:"wasview" json:"wasview"`
	DateRegistered time.Time `bson:"dateregistered" json:"dateregistered"`
}

type Pg_TipsShow struct {
	ID      string `bson:"_id" json:"id"`
	URLimg  string `bson:"urlimg" json:"urlimg"`
	Typetip int    `bson:"typetip" json:"typetip"`
}

type Pg_Anfitrion_PlanBusiness struct {
	IdBusiness     int       `bson:"idbusiness" json:"idbusiness"`
	DateRegistered time.Time `bson:"dateregistered" json:"dateregistered"`
}
