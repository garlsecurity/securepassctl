package logs

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/garlsecurity/go-securepass/securepass/spctl/service"
)

// DateFormat is the supported datetime format
const DateFormat = "2006-01-02"

// Command holds the logs command
var Command = cli.Command{
	Name:        "logs",
	Usage:       "display SecurePass logs",
	ArgsUsage:   " ",
	Description: "Show logs for SecurePass.",
	Action:      ActionLogs,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "csv, C",
			Usage: "Enable CSV full output",
		},
		cli.StringFlag{
			Name:  "start, s",
			Usage: "Start date",
		},
		cli.StringFlag{
			Name:  "end, e",
			Usage: "End date",
		},
		cli.StringFlag{
			Name:  "realm, r",
			Usage: "Set alternate realm",
		},
	},
}

// ActionLogs handles the logs command
func ActionLogs(c *cli.Context) {
	if len(c.Args()) != 0 {
		log.Fatal("error: too many parameters")
	}

	resp, err := service.Service.Logs(c.String("realm"),
		c.String("start"), c.String("end"))
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if !c.Bool("csv") {
		for tstamp, entry := range resp.Logs {
			fmt.Printf("%-19s %s\n", tstamp, entry.Message)
		}
	} else {
		w := csv.NewWriter(os.Stdout)
		defer w.Flush()
		w.Write([]string{"timestamp", "uuid", "message", "realm", "app", "level"})
		for tstamp, entry := range resp.Logs {
			record := []string{tstamp, entry.UUID, entry.Message,
				entry.Realm, entry.App, strconv.Itoa(entry.Level)}
			if err := w.Write(record); err != nil {
				log.Fatalln("error writing record to csv: ", err)
			}
		}
	}
}
