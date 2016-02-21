package service

import (
	"os"

	"github.com/garlsecurity/go-securepass/securepass"
	"gopkg.in/ini.v1"
)

var (
	// Service holds the global settings
	Service *securepass.SecurePass
)

// LoadConfiguration reads configuration from files
func LoadConfiguration(conffiles []string) error {
	cfg, _ := ini.Load([]byte("[default]\nAPP_ID=\nAPP_SECRET=\n"))
	for _, filename := range conffiles {
		if fp, err := os.Open(filename); err == nil {
			fp.Close()
			cfg.Append(filename)
		}
	}
	section, _ := cfg.GetSection("default")
	Service = &securepass.SecurePass{Endpoint: securepass.DefaultRemote}
	if err := section.MapTo(Service); err != nil {
		return err
	}
	return nil
}
