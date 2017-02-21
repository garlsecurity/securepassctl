package group

import (
	"log"
	"os"
	"text/template"

	"github.com/garlsecurity/securepassctl/spctl/service"
	"github.com/urfave/cli"
)

func init() {
	Command.Subcommands = append(Command.Subcommands,
		cli.Command{
			Name:        "info",
			Usage:       "retrieve group's details from SecurePass",
			ArgsUsage:   "GROUP",
			Description: "Retrieve group's details from SecurePass.",
			Action:      ActionInfo,
		})
}

// ActionInfo provides the info subcommand
func ActionInfo(c *cli.Context) {
	const templ = `Group...........: {{.Group}}
Description.....: {{.Description}}
Realm...........: {{.Realm}}
`
	if len(c.Args()) != 1 {
		log.Fatal("error: must specify a group")
	}
	group := c.Args()[0]
	resp, err := service.Service.GroupInfo(group)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	report := template.Must(template.New("report").Parse(templ))

	if err := report.Execute(os.Stdout, resp); err != nil {
		log.Fatalf("error: %v", err)
	}
}
