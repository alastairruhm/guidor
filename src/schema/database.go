package schema

import (
	"encoding/json"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Database represents team info
type Database struct {
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

// DatabaseFrom parse JSON string and returns a Team intance.
func DatabaseFrom(str string) (*Database, error) {
	db := new(Database)
	if err := json.Unmarshal([]byte(str), db); err != nil {
		return nil, err
	}
	return db, nil
}

// String returns JSON string with database instance info
func (d Database) String() string {
	return jsonMarshal(d)
}

// Result returns TeamResult intance
func (d Database) Result() *DatabaseResult {
	return &DatabaseResult{
		ID:        d.ID.String(),
		IP:        d.IP,
		Hostname:  d.Hostname,
		DbType:    d.DbType,
		DbVersion: d.DbVersion,
		DbName:    d.DbName,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

// DatabaseResult represents desensitized database
type DatabaseResult struct {
	ID        string    `json:"id"`
	IP        string    `json:"ip"`
	Hostname  string    `json:"hostname"`
	DbType    string    `json:"db_type"`
	DbVersion string    `json:"db_version"`
	DbName    string    `json:"db_name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// String returns JSON string with desensitized team info
func (t *DatabaseResult) String() string {
	return jsonMarshal(t)
}
