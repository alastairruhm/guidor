package util

import (
	"encoding/hex"

	uuid "github.com/satori/go.uuid"
)

// GenerateToken generate token that 32-chars-length string
func GenerateToken() string {
	uuid := uuid.NewV4()
	return hex.EncodeToString(uuid.Bytes())
}
