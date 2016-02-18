package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/garlsecurity/go-securepass/spctl/cmd/app"
	"github.com/garlsecurity/go-securepass/spctl/cmd/ping"
)

// Commands contains the app's commands
var Commands = []cli.Command{
	ping.Command,
	app.Command,
}
