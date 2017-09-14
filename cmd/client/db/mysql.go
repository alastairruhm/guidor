package db

import (
	"database/sql"
	"fmt"
)

// MySQL ...
type MySQL struct {
	config   Config
	DBClient *sql.DB
}

// ensure MySQL impl Base
var _ Base = &MySQL{}

// NewMySQL ...
func NewMySQL(c Config) (*MySQL, error) {
	var mysql MySQL
	mysql.config = c
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?timeout=5s", c.Username, c.Password, c.Host, c.Port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	mysql.DBClient = db
	return &mysql, nil
}

// Version ...
func (ctx *MySQL) Version() (string, error) {
	var version string
	err := ctx.DBClient.QueryRow("SELECT VERSION() as verion;").Scan(&version)
	if err != nil {
		return "", err
	}
	return version, nil
}

// Type ...
func (ctx *MySQL) Type() string {
	return "mysql"
}

// Auth ...
func (ctx *MySQL) Auth() error {
	err := ctx.DBClient.Ping()
	return err
}

// Backup ...
func (ctx *MySQL) Backup() error {
	fmt.Println("do backup")
	return nil
}

// Close ...
func (ctx *MySQL) Close() error {
	return ctx.DBClient.Close()
}
