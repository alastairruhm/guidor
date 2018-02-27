package config

import (
	"github.com/spf13/viper"
)

// InitConfig load config  file from:
// - ./guidor_client.yml
// - ~/.guidor/guidor_client.yml
// - /etc/guidor/guidor_client.yml
func InitConfig() error {
	viper.SetConfigType("yaml")
	viper.SetConfigName("guidor_client")

	viper.AddConfigPath("/etc/guidor")
	viper.AddConfigPath("$HOME/.guidor")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&Config)

	return err
}

// Config global config
var Config AppConfig

// AppConfig ...
type AppConfig struct {
	AdminUser      string `mapstructure:"admin_user"`
	AdminPassword  string `mapstructure:"admin_password"`
	BackupUser     string `mapstructure:"backup_user"`
	BackupPassword string `mapstructure:"backup_password"`
	BackupDir      string `mapstructure:"backup_dir"`
	TempDir        string `mapstructure:"temp_dir"`
}
