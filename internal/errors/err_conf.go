package errors

import (
	z "github.com/rs/zerolog"
)

var (
	ErrConfConfigNil                   = &AppError{module: "conf", Msg: "Internal error, config is empty", Sev: z.FatalLevel}
	ErrConfConflictingKeys             = &AppError{module: "conf", Msg: "Conflicting key(s) found", Sev: z.ErrorLevel}
	ErrConfKeysNotFound                = &AppError{module: "conf", Msg: "Key(s) not found", Sev: z.ErrorLevel}
	ErrConfKeyNotFound                 = &AppError{module: "conf", Msg: "Key not found", Sev: z.ErrorLevel}
	ErrConfBucketNameMissing           = &AppError{module: "conf", Msg: "Bucket name is missing", Sev: z.ErrorLevel}
	ErrConfSimultaneousUploadsInvalid  = &AppError{module: "conf", Msg: "SimultaneousUploads should be atleast 1, using `1`", Sev: z.WarnLevel}
	ErrConfLoggingSeverityInvalid      = &AppError{module: "conf", Msg: "Logging severity is invalid, using `1`)", Sev: z.WarnLevel}
	ErrConfLoggingFileInvalid          = &AppError{module: "conf", Msg: "LogPath is not given", Sev: z.ErrorLevel}
	ErrConfPathsInaccessible           = &AppError{module: "conf", Msg: "Some paths are accessible", Sev: z.WarnLevel}
	ErrConfWalkDepthInvalid            = &AppError{module: "conf", Msg: "walk depth must be 0 (infinite) or above", Sev: z.ErrorLevel}
	ErrConfObjectNameTypeNone          = &AppError{module: "conf", Msg: "Using default ObjectNameType `absolute`", Sev: z.WarnLevel}
	ErrConfObjectNameTypeInvalid       = &AppError{module: "conf", Msg: "ObjectNameType is invalid", Sev: z.ErrorLevel}
	ErrConfProviderInjectErr           = &AppError{module: "conf", Msg: "Could not build provider configuration", Sev: z.ErrorLevel}
	ErrConfProviderInjectValidationErr = &AppError{module: "conf", Msg: "Provider configuration is invalid", Sev: z.ErrorLevel}
)

var (
	ErrConfCannotAbsPath       = &AppError{module: "conf", Msg: "Cannot get absolute path", Sev: z.ErrorLevel}
	ErrConfCannotMarshalYaml   = &AppError{module: "conf", Msg: "Cannot convert to yaml", Sev: z.ErrorLevel}
	ErrConfCannotUnmarshalYaml = &AppError{module: "conf", Msg: "Cannot convert from yaml", Sev: z.ErrorLevel}
	ErrConfCannotLoadFile      = &AppError{module: "conf", Msg: "Cannot load file", Sev: z.ErrorLevel}
	ErrConfFileExists          = &AppError{module: "conf", Msg: "File already exists", Sev: z.ErrorLevel}
	ErrConfCannotCreateFile    = &AppError{module: "conf", Msg: "Cannot create file", Sev: z.ErrorLevel}
	ErrConfCannotWriteFile     = &AppError{module: "conf", Msg: "Cannot write to file", Sev: z.ErrorLevel}
)
