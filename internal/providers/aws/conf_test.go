package aws

import (
	"testing"

	"s3p/internal/v2/conf"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/stretchr/testify/require"
)

// This should succeed. A profile name with no key/secret pair is set
// No error, not nil ConfAws object returned
func TestConfAwsValidProfile(t *testing.T) {
	p := *conf.T_GetYamlProvider()
	b := *conf.T_GetOkBucket()

	p.Auth[keyAuthProfile] = defaultProfile
	p.Auth[keyAuthKey] = EmptyString
	p.Auth[keyAuthSecret] = EmptyString
	c, err := NewConfAWS(p, b)
	require.NoError(t, err, "Expected no error. Provided a profile and no key / secret pair")
	require.NotNil(t, c, "ConfAws object should never be nil")
}

// This should succeed. A valid key/secret pair is set with no profile
// No error, not nil ConfAws object returned
func TestConfAwsValidKeyPair(t *testing.T) {
	p := *conf.T_GetYamlProvider()
	b := *conf.T_GetOkBucket()

	p.Auth[keyAuthProfile] = EmptyString
	p.Auth[keyAuthKey] = "key"
	p.Auth[keyAuthSecret] = "secret"
	c, err := NewConfAWS(p, b)
	require.NoError(t, err, "Expected no error. Provided a key/secret pair and NO profile")
	require.NotNil(t, c, "ConfAws object should never be nil")
}

func TestConfAwsInvalidNoAuth(t *testing.T) {
	p := *conf.T_GetYamlProvider()
	b := *conf.T_GetOkBucket()

	p.Auth[keyAuthProfile] = EmptyString
	p.Auth[keyAuthKey] = EmptyString
	p.Auth[keyAuthSecret] = EmptyString
	c, err := NewConfAWS(p, b)
	require.Error(t, err, "Expected an error. Didn't provide a profile or a key/secret pair")
	require.NotNil(t, c, "ConfAws object should never be nil")
}

func TestConfAwsInvalidOnlyKey(t *testing.T) {
	p := *conf.T_GetYamlProvider()
	b := *conf.T_GetOkBucket()

	p.Auth[keyAuthProfile] = EmptyString
	p.Auth[keyAuthKey] = "key"
	p.Auth[keyAuthSecret] = EmptyString
	c, err := NewConfAWS(p, b)
	require.Error(t, err, "Expected an error. Only provided key, but no secret")
	require.NotNil(t, c, "ConfAws object should never be nil")
}

func TestConfAwsInvalidOnlySecret(t *testing.T) {
	p := *conf.T_GetYamlProvider()
	b := *conf.T_GetOkBucket()

	p.Auth[keyAuthProfile] = EmptyString
	p.Auth[keyAuthKey] = EmptyString
	p.Auth[keyAuthSecret] = "secret"
	c, err := NewConfAWS(p, b)
	require.Error(t, err, "Expected an error. Only provided secret, but no key")
	require.NotNil(t, c, "ConfAws object should never be nil")
}

func TestConfAwsValidObjectACLs(t *testing.T) {
	p := *conf.T_GetYamlProvider()
	b := *conf.T_GetOkBucket()

	for _, acl := range types.ObjectCannedACLPrivate.Values() {
		p.Options[keyAcl] = string(acl)
		t.Logf("Testing storage class %q", string(acl))
		c, err := NewConfAWS(p, b)
		require.NoError(t, err, "No expected error. Tested valid canned bucket acl strings")
		require.NotNil(t, c, "ConfAws object should never be nil")
	}
}

func TestConfAwsInvalidObjectACLs(t *testing.T) {
	p := *conf.T_GetYamlProvider()
	b := *conf.T_GetOkBucket()

	for _, acl := range []string{"FaKe_AcL", "notPrivate", "Private Read Only", "Public Read Only"} {
		p.Options[keyAcl] = string(acl)
		t.Logf("Testing storage class %q", string(acl))
		c, err := NewConfAWS(p, b)
		require.Error(t, err, "Expected an error. Tested invalid canned bucket acl strings")
		require.NotNil(t, c, "ConfAws object should never be nil")
	}
}

func TestConfValidStorageClasses(t *testing.T) {
	p := *conf.T_GetYamlProvider()
	b := *conf.T_GetOkBucket()

	for _, class := range types.StorageClassStandard.Values() {
		p.Options[keyStorageClass] = string(class)
		t.Logf("Testing storage class %q", string(class))
		c, err := NewConfAWS(p, b)
		require.NoError(t, err, "No expected error. Tested valid storage class strings")
		require.NotNil(t, c, "ConfAws object should never be nil")
	}
}

func TestConfInvalidStorageClasses(t *testing.T) {
	p := *conf.T_GetYamlProvider()
	b := *conf.T_GetOkBucket()

	for _, class := range []string{"FaKe_StOrAgE_Class", "other class", "brokenclass", "standards"} {
		p.Options[keyStorageClass] = class
		t.Logf("Testing invalid storage class %q", class)
		c, err := NewConfAWS(p, b)
		require.Error(t, err, "Expected an error. Tested invalid storage class strings")
		require.NotNil(t, c, "ConfAws object should never be nil")
	}
}
