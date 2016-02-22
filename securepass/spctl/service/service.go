package service

import (
	"io"
	"os"

	"github.com/garlsecurity/go-securepass/securepass"
	"gopkg.in/ini.v1"
)

var (
	// Service holds the global settings
	Service *securepass.SecurePass
	// NSSSettings holds the NSS settings
	NSSSettings *securepass.NSSConfig
	// SSHSettings holds the SSH settings
	SSHSettings *securepass.SSHConfig
)

// LoadConfiguration reads configuration from files
func LoadConfiguration(conffiles []string) error {
	cfg, _ := ini.Load([]byte("[default]\nAPP_ID=\nAPP_SECRET=\n[nss]\n\n[ssh]\n"))
	for _, filename := range conffiles {
		if fp, err := os.Open(filename); err == nil {
			fp.Close()
			cfg.Append(filename)
		}
	}
	defaultSection, _ := cfg.GetSection("default")
	nssSection, _ := cfg.GetSection("nss")
	sshSection, _ := cfg.GetSection("ssh")
	Service = &securepass.SecurePass{Endpoint: securepass.DefaultRemote}
	NSSSettings = &securepass.NSSConfig{
		DefaultGid:   100,
		DefaultHome:  "/home",
		DefaultShell: "/bin/bash",
	}
	SSHSettings = &securepass.SSHConfig{}
	if err := defaultSection.MapTo(Service); err != nil {
		return err
	}
	if err := nssSection.MapTo(NSSSettings); err != nil {
		return err
	}
	if err := sshSection.MapTo(SSHSettings); err != nil {
		return err
	}
	return nil
}

// WriteConfiguration saves configuration to file
func WriteConfiguration(writer io.Writer, appid, endpoint, appsecret, realm, root string) (int64, error) {
	globalConfig := securepass.GlobalConfig{
		*Service,
		*NSSSettings,
		*SSHSettings,
	}

	cfg := ini.Empty()
	if err := ini.ReflectFrom(cfg, &globalConfig); err != nil {
		return 0, err
	}
	return cfg.WriteTo(writer)
}
