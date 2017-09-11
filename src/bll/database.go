package bll

import (
	"github.com/alastairruhm/guidor/src/schema"
)

// Database Business Logic Layer for database
type Database struct {
	*Bll
}

// Register register one database instance to system
func (b *Database) Register(database *schema.Database) (*schema.DatabaseResult, error) {
	res, err := b.Models.Database.Create(database)
	if err != nil {
		return nil, err
	}
	return res, nil
}
