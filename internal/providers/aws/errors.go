package aws

import (
	e "s3p/internal/errors"

	z "github.com/rs/zerolog"
)

var (
	ErrAwsMissingAuthKey      = e.NewError("aws", "need profile or key/secret pair", z.ErrorLevel)
	ErrAwsMissingBucketRegion = e.NewError("aws", "need bucket region", z.ErrorLevel)
	ErrAwsInvalidObjectAcl    = e.NewError("aws", "invalid object acl", z.ErrorLevel)
	ErrAwsInvalidStorageClass = e.NewError("aws", "invalid storage class", z.ErrorLevel)
)
