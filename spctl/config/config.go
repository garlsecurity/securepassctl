package config

import (
	"os"

	"github.com/garlsecurity/go-securepass/securepass"
	"gopkg.in/ini.v1"
)

var (
	Configuration *securepass.SecurePass
)

// FIXME: Handle errors better

func LoadConfiguration(conffiles []string) error {
	cfg, _ := ini.Load([]byte("[default]\nAPP_ID=\nAPP_SECRET=\n"))
	for _, filename := range conffiles {
		if fp, err := os.Open(filename); err == nil {
			fp.Close()
			cfg.Append(filename)
		}
	}
	section, _ := cfg.GetSection("default")
	Configuration = &securepass.SecurePass{Endpoint: securepass.DefaultRemote}
	if err := section.MapTo(Configuration); err != nil {
		return err
	}
	return nil
}
