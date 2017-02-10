package radius

import (
	"log"
	"os"
	"text/template"

	"github.com/urfave/cli"
	"github.com/garlsecurity/securepassctl/spctl/service"
)

func init() {
	Command.Subcommands = append(Command.Subcommands,
		cli.Command{
			Name:        "info",
			Usage:       "retrieve RADIUS details from SecurePass",
			ArgsUsage:   "RADIUSIP",
			Description: "Retrieve RADIUS details from SecurePass.",
			Action:      ActionInfo,
		})
}

// ActionInfo provides the radius info command
func ActionInfo(c *cli.Context) {
	const templ = `FQDN/Name.................: {{.Name}}
Secret....................: {{.Secret}}
Realm.....................: {{.Realm}}
Restrict to group.........: {{.Group}}
Rfid tag..................: {{.Rfid | rfid}}
`

	if len(c.Args()) != 1 {
		log.Fatal("error: must specify an IP address")
	}
	radius := c.Args()[0]
	resp, err := service.Service.RadiusInfo(radius)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	report := template.Must(template.New("report").Funcs(template.FuncMap{"rfid": func(b bool) string {
		if b {
			return "enabled"
		}
		return "disabled"
	}}).Parse(templ))

	if err := report.Execute(os.Stdout, resp); err != nil {
		log.Fatalf("error: %v", err)
	}
}
