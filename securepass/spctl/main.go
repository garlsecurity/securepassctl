package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/garlsecurity/go-securepass/securepass/spctl/cmd"
	"github.com/garlsecurity/go-securepass/securepass/spctl/service"
)

var (
	// OptionDebug contains the --debug flag
	OptionDebug bool
	// Version of spctl
	Version string
)

func init() {
	log.SetPrefix("")
	log.SetFlags(0)
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	SystemConfigFiles := []string{"/etc/securepass.conf",
		"/usr/local/etc/securepass.conf",
		filepath.Join(cwd, "securepass.conf")}
	service.LoadConfiguration(SystemConfigFiles)
}

func main() {
	a := cli.NewApp()
	a.Name = "spctl"
	a.Usage = "manage distributed identities"
	a.Author = "Alessio Treglia"
	a.Email = "alessio@debian.org"
	a.Copyright = "Copyright © 2016 Alessio Treglia <alessio@debian.org>"
	a.Version = Version
	a.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug, D",
			Usage:       "enable debug output",
			Destination: &OptionDebug,
		},
	}
	a.Commands = cmd.Commands

	a.Run(os.Args)
}
