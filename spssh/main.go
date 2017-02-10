package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/garlsecurity/securepassctl"
	"github.com/garlsecurity/securepassctl/spctl/service"
)

var (
	// OptionDebug contains the --debug flag
	OptionDebug bool
	// Version of spctl
	Version string
	// GNUHelpStyle indicates whether enable GNU-like help screen
	GNUHelpStyle string
)

func init() {
	log.SetPrefix("")
	log.SetFlags(0)
}

func loadConfiguration(userConfigFile string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	if userConfigFile != "" {
		if info, err := os.Stat(userConfigFile); err != nil {
			return err
			//log.Fatalf("error: %v", err)
		} else if !info.Mode().IsRegular() {
			//			log.Fatalf("error: %q is not a regular file", userConfigFile)
			return fmt.Errorf("%q is not a regular file", userConfigFile)
		}
		service.LoadConfiguration([]string{userConfigFile})
		return nil
	}
	SystemConfigFiles := []string{"/etc/securepass.conf",
		"/usr/local/etc/securepass.conf",
		filepath.Join(cwd, "securepass.conf")}
	service.LoadConfiguration(SystemConfigFiles)
	return nil
}

func main() {
	//if b, e := strconv.ParseBool(GNUHelpStyle); e != nil || b {

	a := cli.NewApp()

	a.Name = "spssh"
	a.Usage = "Get SSH keys"
	a.Author = "Giuseppe Paterno'"
	a.Email = "gpaterno@gpaterno.com"
	a.Copyright = "Copyright (c) 2016-2017 Giuseppe Paterno' <gpaterno@gpaterno.com>"
	a.Version = Version

	// Set command action
	a.Action = ActionSSHKey

	a.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug, D",
			Usage:       "enable debug output",
			Destination: &OptionDebug,
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "configuration file",
		},
	}

	a.OnUsageError = func(context *cli.Context, err error, isSubcommand bool) error {
		log.Fatalf("error: %v", err)
		return err
	}

	// a.Commands = cmd.Commands

	a.Before = func(c *cli.Context) error {
		if c.GlobalBool("debug") {
			securepassctl.DebugLogger.SetOutput(os.Stderr)
		}
		err := loadConfiguration(c.GlobalString("config"))
		return err
	}

	a.Run(os.Args)
}

// Get user's ssh keys
func ActionSSHKey(c *cli.Context) {
	// Check that we have exactly one argument
	if len(c.Args()) != 1 {
		log.Fatal("error: wrong number of arguments")
	}

	username := c.Args()[0]

	log.Println("Username: ", username)

}
