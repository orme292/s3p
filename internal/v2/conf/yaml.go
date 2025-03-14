package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	e "s3p/internal/errors"

	z "github.com/rs/zerolog"
	"gopkg.in/yaml.v3"
)

type S3PYamlBlock interface {
	Format()
}

type YamlProvider struct {
	Use     string            `yaml:"Use"`
	Auth    map[string]string `yaml:"Auth"`
	Options map[string]string `yaml:"Options"`
}

func (y *YamlProvider) Format() {
	y.Use = strings.ToLower(y.Use)
	y.Options = FormatMap(y.Options, PlainMap)
	y.Auth = FormatMap(y.Auth, PlainMap)
}

type YamlBucket struct {
	Create  bool              `yaml:"Create"`
	Name    string            `yaml:"Name"`
	Options map[string]string `yaml:"Options"`
}

func (y *YamlBucket) Format() {
	y.Name = strings.ToLower(y.Name)
	y.Options = FormatMap(y.Options, PlainMap)
}

type YamlApp struct {
	SimultaneousUploads int    `yaml:"SimultaneousUploads"`
	FollowSymlinks      bool   `yaml:"FollowSymlinks"`
	WalkDepth           int    `yaml:"WalkDepth"`
	RemoteOverwrite     bool   `yaml:"RemoteOverwrite"`
	ObjectNameType      string `yaml:"ObjectNameStyle"`
	ObjectNamePrefix    string `yaml:"ObjectNamePrefix"`
	ObjectPathPrefix    string `yaml:"ObjectPathPrefix"`
}

func (y *YamlApp) Format() {
	y.ObjectNameType = strings.ToLower(y.ObjectNameType)
	y.ObjectNamePrefix = strings.ToLower(y.ObjectNamePrefix)
	y.ObjectPathPrefix = strings.ToLower(y.ObjectPathPrefix)
}

type YamlTags struct {
	LocalPaths     bool              `yaml:"LocalPaths"`
	ChecksumSHA256 bool              `yaml:"ChecksumSHA256"`
	Custom         map[string]string `yaml:"Custom"`
}

func (y *YamlTags) Format() {
	y.Custom = FormatMap(y.Custom, PlainMap)
}

type YamlLogging struct {
	Severity int `yaml:"Severity"`
	ZSev     z.Level
	Console  bool   `yaml:"Console"`
	File     bool   `yaml:"File"`
	LogPath  string `yaml:"LogPath"`
}

func (y *YamlLogging) Format() {
	// nothing to do
}

type YamlConfig struct {
	Version  int           `yaml:"Version"`
	Provider *YamlProvider `yaml:"Provider"`
	Bucket   *YamlBucket   `yaml:"Bucket"`
	App      *YamlApp      `yaml:"App"`
	Tags     *YamlTags     `yaml:"Tags"`
	Logging  *YamlLogging  `yaml:"Logging"`
	Paths    []string      `yaml:"Paths"`
	Skip     []string      `yaml:"Skip"` // TODO: Add Support for Skip paths
	P        *ProviderConfig
}

func (y *YamlConfig) CleanAll() {
	blocks := []S3PYamlBlock{
		y.Provider,
		y.Bucket,
		y.App,
		y.Tags,
		y.Logging,
	}

	for _, block := range blocks {
		block.Format()
	}
}

func NewYamlConfig() *YamlConfig {
	return &YamlConfig{}
}

func (y *YamlConfig) InjectProviderConfig(f GetProviderConfigFunc) error {
	p, err := f(*y.Provider, *y.Bucket)
	if err != nil {
		return e.NewErrorWithBase(e.ErrConfProviderInjectErr, err)
	}

	err = p.Validate()
	if err != nil {
		return e.NewErrorWithBase(e.ErrConfProviderInjectValidationErr, err)
	}

	y.P = &p
	return nil
}

func (y *YamlConfig) LoadFromFile(filename string) error {
	filename, err := fixFilename(filename)
	if err != nil {
		return e.NewErrorWithBase(e.ErrConfCannotAbsPath, err)
	}

	raw, err := loadYamlFile(filename)
	if err != nil {
		return e.NewErrorWithBase(e.ErrConfCannotLoadFile, err)
	}

	err = unmarshalYaml(raw, y)
	if err != nil {
		return e.NewErrorWithBase(e.ErrConfCannotUnmarshalYaml, err)
	}

	return nil
}

func (y *YamlConfig) WriteToFile(filename string) error {
	filename, err := fixFilename(filename)
	if err != nil {
		return e.NewErrorWithBase(e.ErrConfCannotAbsPath, err)
	}

	// Check if we can create the file
	canCreate, err := canCreate(filename)
	if err != nil {
		return e.NewErrorWithBase(e.ErrConfCannotCreateFile, err)
	}
	if !canCreate {
		return e.ErrConfFileExists
	}

	// Marshal the config to YAML
	data, err := marshalYaml(y)
	if err != nil {
		return e.ErrConfCannotMarshalYaml
	}

	f, err := createFile(filename)
	if err != nil {
		return e.NewErrorWithBase(e.ErrConfCannotCreateFile, err)
	}
	defer f.Close()

	err = writeYamlFile(f, data)
	if err != nil {
		return e.NewErrorWithBase(e.ErrConfCannotWriteFile, err)
	}

	return nil
}

func canCreate(path string) (bool, error) {
	// Resolve G304: Potential file inclusion via variable
	if isValidFilename(path) {
		return false, fmt.Errorf("invalid filename: %s", path)
	}

	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return true, nil
		}
		return false, err
	}

	return false, fmt.Errorf("file exists: %s", path)
}

func fixFilename(filename string) (string, error) {
	filename = expandHome(filename)

	filename, err := filepath.Abs(filename)
	if err != nil {
		return "", err
	}
	return filename, nil
}

func isValidFilename(filename string) bool {
	return strings.Contains(filename, "..")
}

func createFile(filename string) (*os.File, error) {
	return os.Create(filename)
}

func loadYamlFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func marshalYaml(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func unmarshalYaml(raw []byte, v interface{}) error {
	return yaml.Unmarshal(raw, v)
}

func writeYamlFile(f *os.File, data []byte) error {

	// Write YAML header
	if _, err := f.WriteString("---\n"); err != nil {
		return err
	}

	// Write the YAML data
	if _, err := f.Write(data); err != nil {
		return err
	}

	return nil
}

func expandHome(path string) string {
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
		}
		return strings.Replace(path, "~", home, 1)
	}
	return path
}
