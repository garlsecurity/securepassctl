package app

import (
	"log"
	"os"
	"text/template"

	"github.com/codegangsta/cli"
	"github.com/garlsecurity/securepassctl/spctl/service"
)

func init() {
	Command.Subcommands = append(Command.Subcommands,
		cli.Command{
			Name:        "info",
			Usage:       "retrieve application's details from SecurePass",
			ArgsUsage:   "APP",
			Description: "Retrieve application's details from SecurePass.",
			Action:      ActionInfo,
		})
}

// ActionInfo provides the info subcommand
func ActionInfo(c *cli.Context) {
	const templ = `Label.................: {{.Label}}
Realm.................: {{.Realm}}
Restrict to group.....: {{.Group}}
Permissions...........: {{.Write | permissions}}
IPv4 Network ACL......: {{.AllowNetworkIPv4}}
IPv6 Network ACL......: {{.AllowNetworkIPv6}}
Privacy mode..........: {{.Privacy}}
`
	if len(c.Args()) != 1 {
		log.Fatal("error: must specify a label")
	}
	app := c.Args()[0]
	resp, err := service.Service.AppInfo(app)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	report := template.Must(template.New("report").Funcs(template.FuncMap{"permissions": func(b bool) string {
		if b {
			return "read/write"
		}
		return "read-only"
	}}).Parse(templ))

	if err := report.Execute(os.Stdout, resp); err != nil {
		log.Fatalf("error: %v", err)
	}
}
