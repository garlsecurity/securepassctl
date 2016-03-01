package user

import (
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/codegangsta/cli"
	"github.com/garlsecurity/securepassctl/spctl/service"
	"golang.org/x/crypto/ssh/terminal"
)

func init() {
	Command.Subcommands = append(Command.Subcommands,
		cli.Command{
			Name:        "passwd",
			Usage:       "modify user's password",
			ArgsUsage:   "USERNAME",
			Description: "Change or disable user password for SecurePass.",
			Action:      ActionPasswd,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "disable, d",
					Usage: "Disable user's password.",
				},
			},
		})
}

// ActionPasswd provides the passwd subcommand
func ActionPasswd(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatal("error: must specify a username")
	}

	username := c.Args()[0]
	if c.Bool("disable") {
		_, err := service.Service.UserPasswordDisable(username)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		log.Println("Password removed")
		os.Exit(0)
	}
	fmt.Fprintf(os.Stderr, "Enter new password: ")
	newPassword, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Fprintln(os.Stderr)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Fprintf(os.Stderr, "Confirm new password: ")
	newPasswordConfirm, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Fprintln(os.Stderr)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	newPasswordString := strings.TrimSpace(string(newPassword))
	if newPasswordString != strings.TrimSpace(string(newPasswordConfirm)) {
		log.Fatalf("error: Password mismatch.")
	}
	_, err = service.Service.UserPasswordChange(username, newPasswordString)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Fprintln(os.Stderr, "Password updated.")
}
