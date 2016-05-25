package group

import "github.com/codegangsta/cli"
import "github.com/garlsecurity/securepassctl/spctl/cmd/group/member"

// Command holds the user subcommands
var Command = cli.Command{
	Name:        "group",
	Usage:       "manage groups",
	ArgsUsage:   "",
	Description: "Manage SecurePass groups.",
	Subcommands: []cli.Command{
			groupmember.Command,
		},
}
