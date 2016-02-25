package ping

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/garlsecurity/securepassctl/securepass/spctl/service"
)

// Command holds the ping command
var Command = cli.Command{
	Name:      "ping",
	Usage:     "ping a SecurePass's remote endpoint",
	ArgsUsage: " ",
	Description: "ping tests a SecurePass's endpoint service status. " +
		"It comes in handy to to test user's configuration.",
	Action: ActionPing,
}

// ActionPing is the ping command handler
func ActionPing(c *cli.Context) {
	resp, err := service.Service.Ping()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("Ping from IPv%d address %s", resp.IPVersion, resp.IP)
}
