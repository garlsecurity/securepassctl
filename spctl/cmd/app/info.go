package app

import (
	"log"
	"os"
	"text/template"

	"github.com/codegangsta/cli"
	"github.com/garlsecurity/go-securepass/securepass"
	"github.com/garlsecurity/go-securepass/spctl/config"
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

// ActionInfo is the info subcommand
func ActionInfo(c *cli.Context) {
	const templ = `
Label.................: {{.Label}}
Realm.................: {{.Realm}}
Restrict to group.....: {{.Group}}
Permissions...........: {{.Write | permissions}}
IPv4 Network ACL......: {{.AllowNetworkIPv4}}
IPv6 Network ACL......: {{.AllowNetworkIPv6}}
Privacy mode..........: {{.Privacy}}
`
	app := func() string {
		switch len(c.Args()) {
		case 0:
			return ""
		case 1:
			return c.Args()[0]
		}
		log.Fatal("too many arguments")
		return ""
	}()
	s, err := securepass.NewSecurePass(config.Configuration.AppID,
		config.Configuration.AppSecret, config.Configuration.Endpoint)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := s.AppInfo(app)
	if resp.RC != 0 {
		log.Fatalf("error: %v", resp.ErrorMsg)
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
