package main

import (
	"path/filepath"
	"strings"
)

func handleCompatMode(args []string) []string {
	var isListCommandAlias bool

	dir, basename := filepath.Split(args[0])
	if strings.HasPrefix(basename, "sp-") {
		newArgs := strings.Split(basename, "-")
		newArgs[0] = dir + newArgs[0]
		// Handle the sp-apps exception, see #17
		switch newArgs[1] {
		case "apps":
			newArgs[1] = "app"
			isListCommandAlias = true
		case "users":
			newArgs[1] = "user"
			isListCommandAlias = true
		}
		if isListCommandAlias {
			newArgs = append(newArgs[:2], append([]string{"list"}, newArgs[2:]...)...)
		}
		newArgs = append(newArgs, args[1:]...)
		return newArgs
	}
	return args
}
