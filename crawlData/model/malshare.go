package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type MalshareDaily struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Md5    string        `bson:"md5"json :"md5"`
	Sha256 string        `bson:"Sha256"json:"Sha256"`

	Sha1 string `bson:"sha1"json:"sha1"`

	Base64 string    `bson:"base64"json:"base64"`
	Date   time.Time `bson:"date"json:"date"`
}
