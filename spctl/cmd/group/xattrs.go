package group

import (
	"log"

	"github.com/urfave/cli"
	"github.com/garlsecurity/securepassctl/spctl/service"
)

func init() {
	Command.Subcommands = append(Command.Subcommands,
		cli.Command{
			Name:        "xattrs",
			Usage:       "access group xattrs",
			ArgsUsage:   "GROUP [list,set,delete] [ATTRIBUTE [VALUE]]",
			Description: "Operate on group's extended attributes in SecurePass.",
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

	group := c.Args()[0]
	operation := c.Args()[1]

	switch operation {

	case "list":
		if len(c.Args()) != 2 {
			log.Fatalf("error: too many arguments")
		}

		resp, err := service.Service.GroupXattrsList(group)

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
		_, err := service.Service.GroupXattrsDelete(group, attribute)

		if err != nil {
			log.Fatalf("error: %v", err)
		}

		log.Printf("Group %q attribute %q deleted", group, attribute)

	case "set":
		if len(c.Args()) != 4 {
			log.Fatalf("error: too few arguments")
		}

		attribute, value := c.Args()[2], c.Args()[3]

		_, err := service.Service.GroupXattrsSet(group, attribute, value)

		if err != nil {
			log.Fatalf("error: %v", err)
		}
		log.Printf("Attribute %q set for group %q", attribute, group)

	default:
		log.Fatalf("error: invalid operation: %q", operation)
	}
}
