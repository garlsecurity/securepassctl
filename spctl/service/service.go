package service

import (
	"io"
	"os"

	"github.com/garlsecurity/securepassctl"
	"gopkg.in/ini.v1"
)

var (
	// Service holds the global settings
	Service *securepassctl.SecurePass
	// NSSSettings holds the NSS settings
	NSSSettings *securepassctl.NSSConfig
	// SSHSettings holds the SSH settings
	SSHSettings *securepassctl.SSHConfig
)

// LoadConfiguration reads configuration from files
func LoadConfiguration(conffiles []string) error {
	cfg := ini.Empty()
	Service = &securepassctl.SecurePass{
		Endpoint: securepassctl.DefaultRemote,
	}
	NSSSettings = &securepassctl.NSSConfig{
		DefaultGid:   100,
		DefaultHome:  "/home",
		DefaultShell: "/bin/bash",
	}
	SSHSettings = &securepassctl.SSHConfig{}

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
func WriteConfiguration(writer io.Writer, s *securepassctl.SecurePass, nss *securepassctl.NSSConfig, ssh *securepassctl.SSHConfig) (int64, error) {
	globalConfig := securepassctl.GlobalConfig{
		SecurePass: *s,
		NSSConfig:  *nss,
		SSHConfig:  *ssh,
	}

	cfg := ini.Empty()
	if err := ini.ReflectFrom(cfg, &globalConfig); err != nil {
		return 0, err
	}
	return cfg.WriteTo(writer)
}
