package main

import (
	"path/filepath"
	"strings"
)

func handleCompatMode(args []string) []string {
	dir, basename := filepath.Split(args[0])
	if strings.HasPrefix(basename, "sp-") {
		newArgs := strings.Split(basename, "-")
		newArgs[0] = dir + newArgs[0]
		// Handle the sp-apps exception, see #17
		if newArgs[1] == "apps" {
			newArgs[1] = "app"
			newArgs = append(newArgs[:2], append([]string{"list"}, newArgs[2:]...)...)
		}
		newArgs = append(newArgs, args[1:]...)
		return newArgs
	}
	return args
}
