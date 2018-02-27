package schema

import (
	"encoding/json"
	"time"
)

const (
	// ModeAuto ...
	ModeAuto = "auto"
	// ModeManual ...
	ModeManual = "manual"
)

// BackupEvent ...
type Event struct {
	ID             string            `json:"id"`
	StartAt        time.Time         `json:"start_at"`
	EndAt          time.Time         `json:"end_at"`
	Mode           string            `json:"mode"`
	RemoteFileHash string            `json:"remote_file_hash"`
	DbExporter     Exporter          `json:"database_exporter"`
	Storers        map[string]Storer `json:"storers"`
	CreatedAt      time.Time         `json:"created_at"`
}

type Exporter struct {
	DbType  string
	Program string
}

type Storer struct {
}

// DatabaseFrom parse JSON string and returns a Team intance.
func EventFrom(str string) (*Database, error) {
	db := new(Database)
	if err := json.Unmarshal([]byte(str), db); err != nil {
		return nil, err
	}
	return db, nil
}

// String returns JSON string with database instance info
func (e Event) String() string {
	return jsonMarshal(e)
}

// Result returns TeamResult intance
// func (e Event) Result() *DatabaseResult {
// 	return &DatabaseResult{
// 		ID:        e.ID.String(),
// 		IP:        e.IP,
// 		Hostname:  e.Hostname,
// 		DbType:    e.DbType,
// 		DbVersion: e.DbVersion,
// 		DbName:    e.DbName,
// 		CreatedAt: e.CreatedAt,
// 		UpdatedAt: e.UpdatedAt,
// 	}
// }

// DatabaseResult represents desensitized database
// type DatabaseResult struct {
// 	ID        string    `json:"id"`
// 	IP        string    `json:"ip"`
// 	Hostname  string    `json:"hostname"`
// 	DbType    string    `json:"db_type"`
// 	DbVersion string    `json:"db_version"`
// 	DbName    string    `json:"db_name"`
// 	CreatedAt time.Time `json:"created_at,omitempty"`
// 	UpdatedAt time.Time `json:"updated_at,omitempty"`
// }

// String returns JSON string with desensitized team info
// func (t *DatabaseResult) String() string {
// 	return jsonMarshal(t)
// }
