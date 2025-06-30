package inits

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/ini.v1"
)

/*
What should this library do?

1. Look for the configuration INI file
- default location will be ~/.s3p
2. Create the file if it is not found
3. Read the contents
- If the needed values are not there, then they should be added

What will be in the INI file?
1. auth file storage location
- /var/s3p/auth
2. plan file storage location
- /var/s3p/plan
3. defaults file location
- etc/s3p
*/

// the ini filename (<homedir>/.s3p - i.e. /Users/admin/.s3p)
var iniFileName string

type Init struct {
	AuthPath   string // the location where auth files will be stored (/etc/s3p/auth/)
	PlanPath   string // the location where plan files will be stored (/etc/s3p/plan/)
	ConfigPath string // the location of the default s3p configuration file (/etc/s3p/config)
}

func Retrieve() Init {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	iniFileName = filepath.Join(home, ".s3p")

	// LooseLoad does not throw an error if the file does not exist
	// If the file does not exist, the default values are set in withDefaults()
	// and then a new ini file is saved.
	cfg, err := ini.LooseLoad(iniFileName)
	if err != nil {
		log.Fatal(err)
	}

	return withDefaults(cfg)
}

func withDefaults(cfg *ini.File) Init {
	init := Init{
		AuthPath:   cfg.Section("Paths").Key("auth").String(),
		PlanPath:   cfg.Section("Paths").Key("plan").String(),
		ConfigPath: cfg.Section("Paths").Key("config").String(),
	}

	if strings.TrimSpace(init.AuthPath) == "" {
		init.AuthPath = "/etc/s3p/auth/"
		cfg.Section("Paths").Key("auth").SetValue(init.AuthPath)
	}
	if strings.TrimSpace(init.PlanPath) == "" {
		init.PlanPath = "/etc/s3p/plan/"
		cfg.Section("Paths").Key("plan").SetValue(init.PlanPath)
	}
	if strings.TrimSpace(init.ConfigPath) == "" {
		init.ConfigPath = "/etc/s3p/config"
		cfg.Section("Paths").Key("config").SetValue(init.ConfigPath)
	}

	saveOut(cfg)

	return init
}

func saveOut(cfg *ini.File) {
	err := cfg.SaveTo(iniFileName)
	if err != nil {
		log.Fatal(err)
	}
}
