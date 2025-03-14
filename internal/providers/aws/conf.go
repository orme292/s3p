package aws

import (
	"fmt"
	"strings"

	"s3p/internal/v2/conf"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type ConfAws struct {
	authProfile  string
	authKey      string
	authSecret   string
	acl          types.ObjectCannedACL
	storageClass types.StorageClass
	region       string

	awsChecksumAlgo types.ChecksumAlgorithm
	awsChecksumMode types.ChecksumMode
}

const (
	defaultProfile = "default"

	keyAuthProfile  = "profile"
	keyAuthKey      = "key"
	keyAuthSecret   = "secret"
	keyAcl          = "acl"
	keyStorageClass = "storageclass"
	keyBucketRegion = "region"
)

func NewConfAWS(p conf.YamlProvider, b conf.YamlBucket) (conf.ProviderConfig, error) {
	c := &ConfAws{}

	fmt.Printf("provider: %+v\n", p)
	err := c.parseProvider(&p)
	if err != nil {
		return c, err
	}

	fmt.Printf("bucket: %+v\n", b)
	err = c.parseBucket(&b)
	if err != nil {
		return c, err
	}

	c.setInternalOpts()

	return c, err
}

func (c *ConfAws) parseProvider(p *conf.YamlProvider) error {
	c.authKey = conf.MapValueSafe(keyAuthKey, p.Auth)
	c.authSecret = conf.MapValueSafe(keyAuthSecret, p.Auth)
	c.authProfile = conf.MapValueSafe(keyAuthProfile, p.Auth)

	err := c.parseAcl(conf.MapValueSafe(keyAcl, p.Options))
	if err != nil {
		return err
	}

	err = c.parseStorageClass(conf.MapValueSafe(keyStorageClass, p.Options))
	if err != nil {
		return err
	}
	return nil
}

func (c *ConfAws) parseBucket(b *conf.YamlBucket) error {
	c.region = conf.MapValueSafe(keyBucketRegion, b.Options)
	return nil
}

func (c *ConfAws) Validate() error {
	if c.authProfile == EmptyString {
		if c.authKey == EmptyString || c.authSecret == EmptyString {
			return ErrAwsMissingAuthKey
		}
	}
	if c.authProfile != EmptyString {
		if c.authKey != EmptyString || c.authSecret != EmptyString {
			return ErrAwsMissingAuthKey
		}
	}
	if c.region == EmptyString {
		return ErrAwsMissingBucketRegion
	}

	return nil
}

func (c *ConfAws) parseAcl(val string) error {
	for _, acl := range types.ObjectCannedACLPrivate.Values() {
		if string(acl) == val {
			c.acl = acl
			return nil
		}
	}
	c.acl = types.ObjectCannedACLPrivate
	return ErrAwsInvalidObjectAcl
}

func (c *ConfAws) parseRegion(val string) error {
	defer func() {
		c.region = val
	}()
	// could validate the region by comparing it against endpoints.json
	// but currently, we'll let it fail on operation with an invalid region
	// TODO: Implement region validator in aws conf
	// see: https://stackoverflow.com/questions/65843812/how-to-get-a-list-of-aws-regions-programmatically-in-go-sdk-v2
	return nil
}

func (c *ConfAws) parseStorageClass(val string) error {
	for _, class := range types.StorageClassStandard.Values() {
		if strings.ToUpper(string(class)) == strings.ToUpper(val) {
			c.storageClass = class
			return nil
		}
	}
	c.storageClass = types.StorageClassStandard
	return ErrAwsInvalidStorageClass
}

func (c *ConfAws) setInternalOpts() {
	c.awsChecksumAlgo = types.ChecksumAlgorithmSha256
	c.awsChecksumMode = types.ChecksumModeEnabled
}
