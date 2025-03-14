package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	z "github.com/rs/zerolog"
)

// TGetOkYaml is used for testing. It returns a valid YamlConfig
// type object to be used for testing.
func T_GetYaml() *YamlConfig {
	y := &YamlConfig{
		Version:  7,
		App:      T_GetYamlApp(),
		Bucket:   T_GetOkBucket(),
		Provider: T_GetYamlProvider(),
		Tags:     T_GetYamlTags(),
		Logging:  T_GetYamlLogging(),
		Paths:    []string{"/tmp/test"},
	}
	return y
}

func T_GetYamlApp() *YamlApp {
	return &YamlApp{
		SimultaneousUploads: 1,
		FollowSymlinks:      false,
		WalkDepth:           0,
		ObjectNameType:      DefObjectNameTypeRelative,
	}
}

func T_GetYamlTags() *YamlTags {
	return &YamlTags{
		Custom: map[string]string{
			"s3ptest": "yes",
		},
	}
}

// T_GetYamlProvider is used for testing. It returns a valid YamlProvider
// type object to be used for testing.
func T_GetYamlProvider() *YamlProvider {
	c := &YamlProvider{
		Use: "aws",
		Auth: map[string]string{
			"profile": "default",
		},
		Options: map[string]string{
			"acl":   "private",
			"class": "standard",
		},
	}
	return c
}

// T_GetOkBucket is used for testing. It returns a valid YamlBucket
// type object to be used for testing.
func T_GetOkBucket() *YamlBucket {
	c := &YamlBucket{
		Create: true,
		Name:   "go_test_bucket",
		Options: map[string]string{
			"acl":    "private",
			"region": "us-east-1",
		},
	}
	return c
}

func T_GetYamlLogging() *YamlLogging {
	return &YamlLogging{
		Console:  true,
		Severity: int(z.DebugLevel),
	}
}

func getTempPath() string {
	return filepath.Join(os.TempDir(), getEpochFilename("", "yaml"))
}

func deleteTempFile(path string) error {
	return os.Remove(path)
}

// getEpochFilename returns a filename based on the current Unix epoch time
// with an optional prefix and suffix (extension)
func getEpochFilename(prefix, suffix string) string {
	return fmt.Sprintf("%s%d%s", prefix, time.Now().Unix(), suffix)
}
