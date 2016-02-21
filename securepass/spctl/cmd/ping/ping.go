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

package ping

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/garlsecurity/go-securepass/securepass/spctl/config"
)

// Command holds the ping command
var Command = cli.Command{
	Name:      "ping",
	Usage:     "ping a SecurePass's remote endpoint",
	ArgsUsage: " ",
	Description: "ping tests a SecurePass's endpoint service status. " +
		"It comes in handy to to test user's configuration.",
	Action: ActionPing,
}

// ActionPing is the ping command handler
func ActionPing(c *cli.Context) {
	resp, err := config.Service.Ping()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("Ping from IPv%d address %s", resp.IPVersion, resp.IP)
}
