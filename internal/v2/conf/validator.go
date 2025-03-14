package conf

import (
	"log"
	"os"

	e "s3p/internal/errors"

	z "github.com/rs/zerolog"
)

const (
	DefObjectNameTypeRelative = "relative"
	DefObjectNameTypeAbsolute = "absolute"
)

type Validator struct {
	y *YamlConfig
}

type valFunc func() error

func NewValidator(y *YamlConfig) *Validator {
	y.CleanAll()
	return &Validator{
		y: y,
	}
}

func (v *Validator) Validate() []error {
	if v.y == nil {
		return []error{e.ErrConfConfigNil}
	}
	var errs []error
	valFuncs := []valFunc{
		v.valAppSimultaneousUploads,
		v.valAppWalkDepth,
		v.valAppObjectNameType,
		v.valBucket,
		v.valLoggingFile,
		v.valLoggingSeverity,
		v.valPaths,
		v.valVersion,
	}

	for _, fn := range valFuncs {
		if fn == nil {
			log.Println("fn is nil")
			continue
		}
		if err := fn(); err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

func (v *Validator) valBucket() error {
	if v.y.Bucket.Name == EmptyStr {
		return e.ErrConfBucketNameMissing
	}
	return nil
}
func (v *Validator) valAppSimultaneousUploads() error {
	if v.y.App.SimultaneousUploads <= 0 {
		v.y.App.SimultaneousUploads = 1
		return e.ErrConfSimultaneousUploadsInvalid
	}
	return nil
}

func (v *Validator) valAppWalkDepth() error {
	if v.y.App.WalkDepth < 0 {
		return e.ErrConfWalkDepthInvalid
	}
	return nil
}

func (v *Validator) valAppObjectNameType() error {
	nameType, err := matchNameType(v.y.App.ObjectNameType)
	v.y.App.ObjectNameType = nameType
	if err != nil {
		if err == e.ErrConfObjectNameTypeNone {
			return err
		}
		return e.ErrConfObjectNameTypeInvalid
	}
	return nil
}

func (v *Validator) valLoggingSeverity() error {
	if v.y.Logging.Severity < int(z.TraceLevel) || v.y.Logging.Severity > int(z.PanicLevel) {
		v.y.Logging.Severity = int(z.InfoLevel)
		v.y.Logging.ZSev = z.InfoLevel
		return e.ErrConfLoggingSeverityInvalid
	}
	v.y.Logging.ZSev = z.Level(v.y.Logging.Severity)
	return nil
}

func (v *Validator) valLoggingFile() error {
	if v.y.Logging.File {
		if v.y.Logging.LogPath == EmptyStr {
			return e.ErrConfLoggingFileInvalid
		}
		v.y.Logging.File = false
	}
	return nil
}

func (v *Validator) valPaths() error {
	var goodPaths []string
	for _, path := range v.y.Paths {
		if _, statErr := os.Stat(path); statErr != nil {
			continue
		}

		if _, openErr := os.Open(path); openErr != nil {
			continue
		}
		goodPaths = append(goodPaths, path)
	}
	if len(goodPaths) != len(v.y.Paths) {
		return e.ErrConfPathsInaccessible
	}
	v.y.Paths = goodPaths
	return nil
}

func (v *Validator) valVersion() error {
	// only 1 version supported now
	v.y.Version = 7
	return nil
}

func matchNameType(val string) (string, error) {
	var nameTypeMap = map[string]string{
		DefObjectNameTypeRelative: DefObjectNameTypeRelative,
		DefObjectNameTypeAbsolute: DefObjectNameTypeAbsolute,
		EmptyStr:                  DefObjectNameTypeAbsolute,
	}
	if result, exists := nameTypeMap[val]; exists {
		if val == EmptyStr {
			return DefObjectNameTypeAbsolute, e.ErrConfObjectNameTypeNone
		}
		return result, nil
	}
	return EmptyStr, e.ErrConfObjectNameTypeInvalid
}
