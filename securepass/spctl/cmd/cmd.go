package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/garlsecurity/go-securepass/securepass/spctl/cmd/app"
	"github.com/garlsecurity/go-securepass/securepass/spctl/cmd/config"
	"github.com/garlsecurity/go-securepass/securepass/spctl/cmd/logs"
	"github.com/garlsecurity/go-securepass/securepass/spctl/cmd/ping"
)

// Commands contains the app's commands
var Commands = []cli.Command{
	ping.Command,
	app.Command,
	config.Command,
	logs.Command,
}
