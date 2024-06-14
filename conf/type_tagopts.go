package conf

import (
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// TagOpts contain the object tagging configuration, but only the ones handled internally by the application.
// Custom tags are put in a separate map named "Tags" inside the AppConfig struct.
type TagOpts struct {
	ChecksumSHA256 bool
	OriginPath     bool

	AwsChecksumAlgorithm types.ChecksumAlgorithm
	AwsChecksumMode      types.ChecksumMode
}

func (to *TagOpts) build(inc *ProfileIncoming) error {

	to.OriginPath = inc.TagOptions.OriginPath
	to.ChecksumSHA256 = inc.TagOptions.ChecksumSHA256
	to.AwsChecksumAlgorithm = types.ChecksumAlgorithmSha256
	to.AwsChecksumMode = types.ChecksumModeEnabled

	return to.validate()

}

func (to *TagOpts) validate() error {

	// nothing to validate yet
	return nil

}
