package user

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
			Usage:       "retrieve user's details from SecurePass",
			ArgsUsage:   "USERNAME",
			Description: "Retrieve user's details from SecurePass.",
			Action:      ActionInfo,
		})
}

// ActionInfo provides the info subcommand
func ActionInfo(c *cli.Context) {
	const templ = `Name....................: {{.Name}}
Surname.................: {{.Surname}}
E-mail..................: {{.Email}}
Mobile nr...............: {{.Mobile}}
National ID.............: {{.Nin}}
RFID tag................: {{.Rfid}}
Token...................: {{.Token}}
User status.............: {{.Enabled | boolToString}}
Password status.........: {{.Password | boolToString}}
`
	if len(c.Args()) != 1 {
		log.Fatal("error: must specify a username")
	}
	username := c.Args()[0]
	resp, err := service.Service.UserInfo(username)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	report := template.Must(template.New("report").Funcs(template.FuncMap{"boolToString": func(b bool) string {
		if b {
			return "Enabled"
		}
		return "Disabled"
	}}).Parse(templ))

	if err := report.Execute(os.Stdout, resp); err != nil {
		log.Fatalf("error: %v", err)
	}
}
