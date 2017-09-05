package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Instance Model
type Instance struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Token     string        `json:"token" bson:"token"`
	IP        string        `json:"ip" bson:"ip"`
	Hostname  string        `json:"hostname" bson:"hostname"`
	DbType    string        `json:"db_type" bson:"db_type"`
	DbVersion string        `json:"db_version" bson:"db_version"`
	DbName    string        `json:"db_name" bson:"db_name"`
	CreatedAt time.Time     `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at,omitempty" bson:"updated_at"`
}
