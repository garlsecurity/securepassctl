package realm

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/garlsecurity/securepassctl/spctl/service"
)

func init() {
	Command.Subcommands = append(Command.Subcommands,
		cli.Command{
			Name:        "xattrs",
			Usage:       "access realm xattrs",
			ArgsUsage:   "REALM [list,set,delete] [ATTRIBUTE [VALUE]]",
			Description: "Operate on realms' extended attributes in SecurePass.",
			Action:      ActionXattrs,
		})
}

// ActionXattrs provides the provision subcommand
func ActionXattrs(c *cli.Context) {
	if len(c.Args()) < 2 {
		log.Fatal("error: too few arguments")
	}

	if len(c.Args()) > 4 {
		log.Fatal("error: too many arguments")
	}

	realm := c.Args()[0]
	operation := c.Args()[1]

	switch operation {

	case "list":
		if len(c.Args()) != 2 {
			log.Fatalf("error: too many arguments")
		}

		resp, err := service.Service.RealmXattrsList(realm)

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
		_, err := service.Service.RealmXattrsDelete(realm, attribute)

		if err != nil {
			log.Fatalf("error: %v", err)
		}

		log.Printf("Realm %q attribute %q deleted", realm, attribute)

	case "set":
		if len(c.Args()) != 4 {
			log.Fatalf("error: too few arguments")
		}

		attribute, value := c.Args()[2], c.Args()[3]

		_, err := service.Service.RealmXattrsSet(realm, attribute, value)

		if err != nil {
			log.Fatalf("error: %v", err)
		}
		log.Printf("Attribute %q set for realm %q", attribute, realm)

	default:
		log.Fatalf("error: invalid operation: %q", operation)
	}
}
