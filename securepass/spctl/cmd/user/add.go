package user

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/garlsecurity/go-securepass/securepass"
	"github.com/garlsecurity/go-securepass/securepass/spctl/service"
)

func init() {
	Command.Subcommands = append(Command.Subcommands,
		cli.Command{
			Name:        "add",
			Usage:       "add user",
			ArgsUsage:   "USERNAME",
			Description: "Add user to SecurePass.",
			Action:      ActionAdd,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, n",
					Usage: "First name",
				},
				cli.StringFlag{
					Name:  "surname, s",
					Usage: "Last name",
				},
				cli.StringFlag{
					Name:  "email, e",
					Usage: "E-mail",
				},
				cli.StringFlag{
					Name:  "mobile, m",
					Usage: "Group name (restriction)",
				},
				cli.StringFlag{
					Name:  "nin",
					Usage: "National Identification Number (optional)",
				},
				cli.StringFlag{
					Name:  "rfid",
					Usage: "RFID tag (optional)",
				},
				cli.StringFlag{
					Name:  "manager",
					Usage: "Manager user id (optional)",
				},
			},
		})
}

// ActionAdd provides the add subcommand
func ActionAdd(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatal("error: must specify a username")
	}
	username := c.Args()[0]

	if c.Bool("debug") {
		log.Println("Adding user", username)
	}
	_, err := service.Service.UserAdd(&securepass.UserDescriptor{
		Username: username,
		Name:     c.String("name"),
		Surname:  c.String("surname"),
		Email:    c.String("email"),
		Mobile:   c.String("mobile"),
		Nin:      c.String("nin"),
		Rfid:     c.String("rfid"),
		Manager:  c.String("manager"),
	})
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
