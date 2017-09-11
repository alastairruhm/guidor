package register

import (
	"context"

	"github.com/alastairruhm/guidor/client"
	"github.com/alastairruhm/guidor/src/schema"
	logging "github.com/op/go-logging"
	"github.com/spf13/cobra"
)

var (
	ip          string
	hostname    string
	dbtype      string
	dbversion   string
	serviceName string
)

var log = logging.MustGetLogger("guidor")

func init() {
	Cmd.Flags().StringVarP(&ip, "ip", "", "", "guidor client ip")
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
	c := client.NewClient(nil)
	ctx := context.TODO()
	d := schema.Database{
		IP:        "10.34.50.178",
		Hostname:  "test",
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
