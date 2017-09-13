package db

import (
	"database/sql"
	"fmt"
	"log"
)

// MySQL ...
type MySQL struct {
	config   DatabaseConfig
	DBClient *sql.DB
}

// NewMySQL ...
func NewMySQL(c DatabaseConfig) (*MySQL, error) {
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
		log.Fatal(err)
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

// Close ...
func (ctx *MySQL) Close() error {
	return ctx.DBClient.Close()
}
