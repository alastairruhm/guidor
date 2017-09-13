package register

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"os"

	"github.com/alastairruhm/guidor/client"
	"github.com/alastairruhm/guidor/db"
	"github.com/alastairruhm/guidor/src/schema"
	logging "github.com/op/go-logging"
	"github.com/spf13/cobra"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// ip          string
	// hostname string
	dbUser string
	dbPass string
	dbName string
	dbHost string
	dbPort string
	dbType string
	// dbtype      string
	// dbversion   string
	// serviceName string
)

var log = logging.MustGetLogger("guidor")

func init() {
	// Cmd.Flags().StringVarP(&ip, "ip", "", "", "guidor client ip")
	Cmd.Flags().StringVarP(&dbUser, "dbuser", "", "", "database user")
	Cmd.Flags().StringVarP(&dbPass, "dbpass", "", "", "database password")
	Cmd.Flags().StringVarP(&dbName, "dbname", "", "", "database name")
	Cmd.Flags().StringVarP(&dbHost, "dbhost", "", "127.0.0.1", "database name")
	Cmd.Flags().StringVarP(&dbPort, "dbport", "", "", "database name")
	Cmd.Flags().StringVarP(&dbType, "dbtype", "", "", "database type")
}

// Cmd Register
var Cmd = &cobra.Command{
	Use:     "register",
	Short:   "Register commands: guidor client register --help",
	Long:    "Register commands: guidor register [command]",
	Aliases: []string{"r"},
	Run:     registerClient,
}

func registerClient(cmd *cobra.Command, args []string) {
	// check database connection with dsn params
	config := db.DatabaseConfig{
		Username: dbUser,
		Password: dbPass,
		Database: dbName,
		Host:     dbHost,
		Port:     dbPort,
	}
	database, err := db.NewMySQL(config)
	defer database.Close()

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Info("check database authentication success")

	// collect database version
	c := client.NewClient(nil)
	ctx := context.TODO()

	d := schema.Database{
		IP:        GetLocalIP(),
		Hostname:  MustHostName(),
		DbType:    "mysql",
		DbVersion: "5.5",
		DbName:    "test",
	}
	database, _, err := c.Databases.Register(ctx, d)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Infof("Token: %s", database.Token)
}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// MustHostName get the hostname, return empty string if err occur
func MustHostName() string {
	n, err := os.Hostname()
	if err != nil {
		return ""
	}
	return n
}

// CheckDBConnection ...
func CheckDBConnection(dbuser string, dbpasswd string, host string, port string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?timeout=5s", dbuser, dbpasswd, host, port)
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		// log.Fatal(err)
		return err
	}
	err = db.Ping()
	return err
}
