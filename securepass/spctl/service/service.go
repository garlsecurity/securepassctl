package service

import (
	"io"
	"os"

	"github.com/garlsecurity/securepassctl/securepass"
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
	cfg := ini.Empty()
	Service = &securepass.SecurePass{
		Endpoint: securepass.DefaultRemote,
	}
	NSSSettings = &securepass.NSSConfig{
		DefaultGid:   100,
		DefaultHome:  "/home",
		DefaultShell: "/bin/bash",
	}
	SSHSettings = &securepass.SSHConfig{}

	for _, filename := range conffiles {
		if fp, err := os.Open(filename); err == nil {
			fp.Close()
			cfg.Append(filename)
		}
	}

	defaultSection, err := cfg.GetSection("default")
	if err != nil {
		defaultSection, _ = cfg.NewSection("default")
		err = defaultSection.ReflectFrom(Service)
		if err != nil {
			panic(err)
		}
	} else {
		if err = defaultSection.MapTo(Service); err != nil {
			panic(err)
		}
	}

	nssSection, err := cfg.GetSection("nss")
	if err != nil {
		nssSection, _ = cfg.NewSection("nss")
		err = nssSection.ReflectFrom(NSSSettings)
		if err != nil {
			panic(err)
		}
	} else {
		if err = nssSection.MapTo(NSSSettings); err != nil {
			panic(err)
		}
	}

	sshSection, err := cfg.GetSection("ssh")
	if err != nil {
		sshSection, _ = cfg.NewSection("ssh")
		err = sshSection.ReflectFrom(SSHSettings)
		if err != nil {
			panic(err)
		}
	} else {
		if err = sshSection.MapTo(SSHSettings); err != nil {
			panic(err)
		}
	}

	return nil
}

// WriteConfiguration saves configuration to file
func WriteConfiguration(writer io.Writer, s *securepass.SecurePass, nss *securepass.NSSConfig, ssh *securepass.SSHConfig) (int64, error) {
	globalConfig := securepass.GlobalConfig{
		*s, *nss, *ssh,
	}

	cfg := ini.Empty()
	if err := ini.ReflectFrom(cfg, &globalConfig); err != nil {
		return 0, err
	}
	return cfg.WriteTo(writer)
}
