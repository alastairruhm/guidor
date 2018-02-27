package config

import (
	"bytes"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var configExample = []byte(`admin_user: root
admin_password: password

backup_user: backup_user
backup_password: password

backup_dir:  /Users/leon/.dbtool/backup
temp_dir: /Users/leon/.dbtool/tmp
`)

func TestConfigInit(t *testing.T) {
	viper.Reset()

	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(bytes.NewReader(configExample)); err != nil {
		t.Error(err)
	}

	assert.Equal(t, "root", viper.GetViper().GetString("admin_user"))
}
