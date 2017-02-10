package radius

import (
	"log"

	"github.com/urfave/cli"
	"github.com/garlsecurity/securepassctl"
	"github.com/garlsecurity/securepassctl/spctl/service"
)

func init() {
	Command.Subcommands = append(Command.Subcommands,
		cli.Command{
			Name:        "add",
			Usage:       "add a RADIUS",
			ArgsUsage:   "RADIUS",
			Description: "Add a RADIUS to SecurePass.",
			Action:      ActionAdd,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "realm, r",
					Usage: "Set alternate realm",
				},
				cli.StringFlag{
					Name:  "fqdn, f",
					Usage: "FQDN/Name",
				},
				cli.StringFlag{
					Name:  "secret, s",
					Usage: "Shared secret",
				},
				cli.StringFlag{
					Name:  "group, g",
					Usage: "Group name (restriction)",
				},
				cli.BoolFlag{
					Name:  "rfid, R",
					Usage: "Enable RFID tag as username",
				},
			},
		})
}

// ActionAdd provides the add subcommand
func ActionAdd(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatal("error: must specify an IP address")
	}
	radius := c.Args()[0]
	securepassctl.DebugLogger.Printf("Adding RADIUS %q", radius)
	_, err := service.Service.RadiusAdd(&securepassctl.RadiusDescriptor{
		Radius: radius,
		Name:   c.String("fqdn"),
		Secret: c.String("secret"),
		Group:  c.String("group"),
		Realm:  c.String("realm"),
		Rfid:   c.Bool("rfid"),
	})
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
