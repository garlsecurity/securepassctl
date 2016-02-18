//    Copyright © 2016 Alessio Treglia <alessio@debian.org>
//
//    This program is free software; you can redistribute it and/or modify
//    it under the terms of the GNU General Public License as published by
//    the Free Software Foundation; either version 2 of the License, or
//    (at your option) any later version.
//
//    This program is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//    GNU General Public License for more details.
//
//    You should have received a copy of the GNU General Public License along
//    with this program; if not, write to the Free Software Foundation, Inc.,
//    51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.

package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/garlsecurity/go-securepass/spctl/cmd"
	"github.com/garlsecurity/go-securepass/spctl/config"
)

var (
	SystemConfigFiles []string
)

func init() {
	log.SetPrefix("spctl: ")
	log.SetFlags(0)
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	SystemConfigFiles = []string{"/etc/securepass.conf",
		"/usr/local/etc/securepass.conf",
		filepath.Join(cwd, "securepass.conf")}
	config.LoadConfiguration(SystemConfigFiles)
}

func main() {
	app := cli.NewApp()
	app.Name = "spctl"
	app.Usage = "manage distributed identities"
	app.Commands = []cli.Command{
		{
			Name:        "ping",
			Usage:       "ping a SecurePass's remote endpoint",
			Description: "test SecurePass's service status and user's configuration",
			Action:      cmd.ActionPing,
		},
	}

	app.Run(os.Args)
}
