package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/orme292/s3packer/logbot"
	"gopkg.in/yaml.v2"
)

/*
New returns a new Configuration object with default values for logging.
*/
func New() Configuration {
	return Configuration{
		Logger: logbot.LogBot{
			Level:       logbot.INFO,
			FlagConsole: false,
			FlagFile:    false,
			Path:        "log.log",
		},
	}
}

/*
Load loads and sanitizes the configuration from the specified file.

The file is expected to be in YAML format.
1. Load the file.
2. Unmarshal the YAML into the Configuration object.
3. Configure logging.
4. Process Authentication.
5. Process Bucket details.
6. Process Options.
7. Return the Configuration object.
*/
func (c *Configuration) Load(file string) error {
	file, err := filepath.Abs(file)
	if err != nil {
		return errors.New("unable to determine filename path: " + err.Error())
	}

	f, err := os.ReadFile(filepath.Clean(file))
	if err != nil {
		return errors.New(err.Error())
	}

	fmt.Println("Using profile", file)

	err = yaml.Unmarshal(f, &c)
	if err != nil {
		return errors.New(err.Error())
	}

	err = c.Validate()
	if err != nil {
		return err
	}

	c.Logger.FlagConsole = c.Logging[ProfileLoggingToConsole].(bool)
	c.Logger.Level = c.Logger.ParseIntLevel(c.Logging[ProfileLoggingLevel].(int))
	c.Logger.FlagFile = c.Logging[ProfileLoggingToFile].(bool)
	if c.Logging[ProfileLoggingToFile].(bool) == true && c.Logging[ProfileLoggingFilename].(string) != "" {
		path, err := filepath.Abs(c.Logging[ProfileLoggingFilename].(string))
		if err != nil {
			return errors.New("unable to get absolute path of log file: " + err.Error())
		}
		c.Logger.Path = path
	}
	c.Logger.Level = c.Logger.ParseIntLevel(c.Logging[ProfileLoggingLevel])

	if c.Files == nil && c.Dirs == nil {
		return errors.New("no files or directories specified")
	}
	if len(c.Files) == 0 && len(c.Dirs) == 0 {
		return errors.New("no files or directories specified")
	}

	return nil
}

/*
GetACL returns an ObjectCannedACL that matches the provided ACL string.
*/
func (c *Configuration) GetACL(acl string) types.ObjectCannedACL {
	switch acl {
	case ACLPrivate:
		return types.ObjectCannedACLPrivate
	case ACLPublicRead:
		return types.ObjectCannedACLPublicRead
	case ACLPublicReadWrite:
		return types.ObjectCannedACLPublicReadWrite
	case ACLAuthenticatedRead:
		return types.ObjectCannedACLAuthenticatedRead
	case ACLAwsExecRead:
		return types.ObjectCannedACLAwsExecRead
	case ACLBucketOwnerRead:
		return types.ObjectCannedACLBucketOwnerRead
	case ACLBucketOwnerFullControl:
		return types.ObjectCannedACLBucketOwnerFullControl
	default:
		return types.ObjectCannedACLPrivate
	}
}

/*
GetStorageClass returns a StorageClass that matches the provided class string.
*/
func (c *Configuration) GetStorageClass(class string) types.StorageClass {
	switch class {
	case StorageClassStandard:
		return types.StorageClassStandard
	case StorageClassReducedRedundancy:
		return types.StorageClassReducedRedundancy
	case StorageClassGlacier:
		return types.StorageClassGlacier
	case StorageClassStandardIA:
		return types.StorageClassStandardIa
	case StorageClassOneZoneIA:
		return types.StorageClassOnezoneIa
	case StorageClassIntelligentTiering:
		return types.StorageClassIntelligentTiering
	case StorageClassGlacierIR:
		return types.StorageClassGlacierIr
	case StorageClassDeepArchive:
		return types.StorageClassDeepArchive
	case StorageClassSnow:
		return types.StorageClassSnow
	default:
		return types.StorageClassStandard
	}
}
