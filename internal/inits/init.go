// Package inits implements a utility that loads the environment config file (~/.s3p).
//
// The package will verify that the necessary directories (referred to as libs) exist and are
// accessible. If not, the package will create the paths. There are three possible fatal failures
// when using this package.
//
// If the user's home directory cannot be determined using os.UserHomeDir(), then a fatal failure occurs.
// If loose loading does not work, then a fatal failure is returned.
// If the INI file cannot be saved, then a fatal failure is returned.
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

const (
    defaultAuthPath   = "s3p/auth/"
    defaultPlanPath   = "s3p/plan/"
    defaultConfigPath = "s3p/config"
)

// the ini filename (<homedir>/.s3p - i.e. /Users/admin/.s3p)
var iniFileName string
var homePath string

type Init struct {
    AuthPath   string // the location where auth files will be stored (/etc/s3p/auth/)
    PlanPath   string // the location where plan files will be stored (/etc/s3p/plan/)
    ConfigPath string // the location of the default s3p configuration file (/etc/s3p/config)

    DevMode bool // whether DevMode is active
}

func Retrieve() Init {
    home, err := os.UserHomeDir()
    if err != nil {
        log.Fatal(err)
    }
    homePath = home

    iniFileName = filepath.Join(home, ".s3p")

    // LoadSources is used to pass specific load options. 'Loose' prevents the LoadSources from
    // throwing an error if the target file does not exist. 'Insensitive' ignores the case of
    // sections and keys. 'SkipUnrecognizableLines' ignores lines that aren't key/value pairs.
    cfg, err := ini.LoadSources(ini.LoadOptions{
        Loose:                   true,
        Insensitive:             true,
        SkipUnrecognizableLines: true,
    }, iniFileName)

    init := withDefaults(cfg)

    // create the library paths if they do not exist
    // config file is not checked here, it can be created when the app loads the file
    makeLibsExist(init)

    return init
}

func withDefaults(cfg *ini.File) Init {
    changed := false

    init := Init{
        AuthPath:   cfg.Section("Paths").Key("auth").String(),
        PlanPath:   cfg.Section("Paths").Key("plan").String(),
        ConfigPath: cfg.Section("Paths").Key("config").String(),
        DevMode:    cfg.Section("DevMode").Key("enabled").MustBool(false),
    }

    // If there is a "DevMode" section with an "enabled" key/value, then we read it.
    // Otherwise, DevMode is ignored and does not exist.
    if cfg.HasSection("DevMode") {
        if cfg.Section("DevMode").HasKey("enabled") {
            init.DevMode = cfg.Section("DevMode").Key("enabled").MustBool(true)
        }
    }

    // If the Paths.auth, Paths.plan, or Paths.config fields are empty, then we set
    // them to the default, and add each one to the ini object (cfg).
    // changed is set to true so that the new defaults are written out to the file.
    if strings.TrimSpace(init.AuthPath) == "" {
        init.AuthPath = filepath.Join(homePath, defaultAuthPath)
        cfg.Section("Paths").Key("auth").SetValue(init.AuthPath)
        changed = true
    }
    if strings.TrimSpace(init.PlanPath) == "" {
        init.PlanPath = filepath.Join(homePath, defaultPlanPath)
        cfg.Section("Paths").Key("plan").SetValue(init.PlanPath)
        changed = true
    }
    if strings.TrimSpace(init.ConfigPath) == "" {
        init.ConfigPath = filepath.Join(homePath, defaultConfigPath)
        cfg.Section("Paths").Key("config").SetValue(init.ConfigPath)
        changed = true
    }

    // If changes are made to the config, then save them.
    if changed {
        saveOut(cfg)
    }

    return init
}

// Save changes to the ini file.
func saveOut(cfg *ini.File) {
    err := cfg.SaveTo(iniFileName)
    if err != nil {
        log.Fatal(err)
    }
}
