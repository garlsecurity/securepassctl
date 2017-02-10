package user

import (
	"log"

	"github.com/urfave/cli"
	"github.com/garlsecurity/securepassctl/spctl/service"
)

func init() {
	Command.Subcommands = append(Command.Subcommands,
		cli.Command{
			Name:        "xattrs",
			Usage:       "access user xattrs",
			ArgsUsage:   "USERNAME [list,set,delete] [ATTRIBUTE [VALUE]]",
			Description: "Operate on users' extended attributes in SecurePass.",
			Action:      ActionXattrs,
		})
}

// ActionXattrs provides the provision subcommand
func ActionXattrs(c *cli.Context) {
	// We do not need to check less than x, because we have checks after
	if len(c.Args()) > 4 {
		log.Fatal("error: too many arguments")
	}

	username := c.Args()[0]
	operation := c.Args()[1]

	switch operation {
	case "list":
		if len(c.Args()) != 2 {
			log.Fatalf("error: too many arguments")
		}
		resp, err := service.Service.UserXattrsList(username)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		for a, v := range *resp {
			if a != "rc" && a != "errorMsg" {
				log.Printf("%s: %s", a, v)
			}
		}
	case "delete":
		if len(c.Args()) != 3 {
			log.Fatalf("error: must specify an attribute")
		}
		attribute := c.Args()[2]
		_, err := service.Service.UserXattrsDelete(username, attribute)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		log.Printf("User %q attribute %q deleted", username, attribute)
	case "set":
		if len(c.Args()) != 4 {
			log.Fatalf("error: too few arguments")
		}
		attribute, value := c.Args()[2], c.Args()[3]
		_, err := service.Service.UserXattrsSet(username, attribute, value)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		log.Printf("Attribute %q set for user %q", attribute, username)
	default:
		log.Fatalf("error: invalid operation: %q", operation)
	}
}
