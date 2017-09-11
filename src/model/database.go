package model

import (
	"time"

	"github.com/alastairruhm/guidor/src/schema"
	"github.com/alastairruhm/guidor/src/service/mongodb"
	"github.com/alastairruhm/guidor/src/util"
	"gopkg.in/mgo.v2/bson"
)

const (
	databaseCollectionName string = "databases"
)

// Database ...
type Database struct {
	conn *mongodb.Collection
}

func newDatabaseCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession(databaseCollectionName)
}

// Create create database record
func (d *Database) Create(database *schema.Database) (*schema.DatabaseResult, error) {
	d.conn = newDatabaseCollection()
	defer d.conn.Close()

	// set default mongodb ID  and created date
	database.ID = bson.NewObjectId()

	// instance.Token = bson.NewObjectId().Hex()
	database.Token = util.GenerateToken()
	database.CreatedAt = time.Now()
	// Insert database record to mongodb
	err := d.conn.Session.Insert(&database)
	if err != nil {
		return nil, dbError(err)
	}
	ds := database.Result()
	return ds, nil
}
