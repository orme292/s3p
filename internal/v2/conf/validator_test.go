package conf

import (
	"errors"
	"os"
	"testing"

	e "s3p/internal/errors"

	z "github.com/rs/zerolog"

	"github.com/stretchr/testify/require"
)

func TestValidatorStruct(t *testing.T) {
	v := getValObj(t)
	defer deleteTempFiles(v)
	errs := v.Validate()
	require.True(t, expectErrCount(0, errs), "Expected 0 error")
}

func TestValFuncAppSimultaneousUploads(t *testing.T) {
	v := getValObj(t)
	defer deleteTempFiles(v)
	v.y.App.SimultaneousUploads = 0
	errs := v.Validate()
	require.True(t, expectErrCount(1, errs), "Expected 1 error")
	require.True(t, expectErr(t, e.ErrConfSimultaneousUploadsInvalid, errs), "Expected ErrConfSimultaneousUploadsInvalid")
}

func TestValFuncBucket(t *testing.T) {
	v := getValObj(t)
	defer deleteTempFiles(v)
	v.y.Bucket.Name = ""
	errs := v.Validate()
	require.True(t, expectErrCount(1, errs), "Expected 1 error")
	require.True(t, expectErr(t, e.ErrConfBucketNameMissing, errs), "Expected ErrConfBucketNameMissing")
}

func TestValFuncAppObjectNameType(t *testing.T) {
	v := getValObj(t)
	defer deleteTempFiles(v)
	t.Logf("Testing default %q\n", v.y.App.ObjectNameType)
	errs := v.Validate()
	require.True(t, expectErrCount(0, errs), "Expected 0 error")

	v.y.App.ObjectNameType = "invalid"
	t.Logf("Testing %q\n", v.y.App.ObjectNameType)
	errs = v.Validate()
	require.True(t, expectErrCount(1, errs), "Expected 1 error")
	require.True(t, expectErr(t, e.ErrConfObjectNameTypeInvalid, errs), "Expected ErrConfObjectNameTypeInvalid")

	v.y.App.ObjectNameType = "rElAtIvE"
	t.Logf("Testing %q\n", v.y.App.ObjectNameType)
	errs = v.Validate()
	require.True(t, expectErrCount(1, errs), "Expected 1 error")
	require.True(t, expectErr(t, e.ErrConfObjectNameTypeInvalid, errs), "Expected ErrConfObjectNameTypeInvalid")

	v.y.App.ObjectNameType = "abSolute"
	t.Logf("Testing %q\n", v.y.App.ObjectNameType)
	errs = v.Validate()
	require.True(t, expectErrCount(1, errs), "Expected 1 error")
	require.True(t, expectErr(t, e.ErrConfObjectNameTypeInvalid, errs), "Expected ErrConfObjectNameTypeInvalid")

	v.y.App.ObjectNameType = EmptyStr
	t.Logf("Testing %q\n", v.y.App.ObjectNameType)
	errs = v.Validate()
	require.Equal(t, v.y.App.ObjectNameType, DefObjectNameTypeAbsolute, "Expected DefObjectNameTypeAbsolute")
	require.True(t, expectErrCount(1, errs), "Expected 1 error")
	require.True(t, expectErr(t, e.ErrConfObjectNameTypeNone, errs), "Expected ErrConfObjectNameTypeNone")
}

func TestValFuncLoggingSeverity(t *testing.T) {
	v := getValObj(t)
	defer deleteTempFiles(v)
	v.y.Logging.Severity = 6
	errs := v.Validate()
	require.True(t, expectErrCount(1, errs), "Expected 1 error")
	require.True(t, expectErr(t, e.ErrConfLoggingSeverityInvalid, errs), "Expected ErrConfObjectNameTypeInvalid")

	v.y.Logging.Severity = -1
	errs = v.Validate()
	require.True(t, expectErrCount(0, errs), "Expected 0 error")

	v.y.Logging.Severity = int(z.NoLevel)
	errs = v.Validate()
	require.True(t, expectErrCount(1, errs), "Expected 1 error")
	require.True(t, expectErr(t, e.ErrConfLoggingSeverityInvalid, errs), "Expected ErrConfObjectNameTypeInvalid")
}

func TestValFuncWalkDepth(t *testing.T) {
	v := getValObj(t)
	defer deleteTempFiles(v)
	v.y.App.WalkDepth = -1
	errs := v.Validate()
	require.True(t, expectErrCount(1, errs), "Expected 1 error")
	require.True(t, expectErr(t, e.ErrConfWalkDepthInvalid, errs), "Expected ErrConfWalkDepthInvalid")
}

func getValObj(t *testing.T) *Validator {
	v := NewValidator(T_GetYaml())
	v.y.Paths = []string{}
	createTempFiles(t, v)
	return v
}

func createTempFiles(t *testing.T, v *Validator) {
	for i := 0; i < 10; i++ {
		file, err := os.CreateTemp("", "s3pvaltest_*.txt")
		v.y.Paths = append(v.y.Paths, file.Name())
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
			deleteTempFiles(v)
			t.FailNow()
		}
	}
}

func deleteTempFiles(v *Validator) {
	for i := 0; i < len(v.y.Paths); i++ {
		os.Remove(v.y.Paths[i])
	}
}

func printErrs(t *testing.T, errs []error) {
	for _, err := range errs {
		if err == nil {
			t.Log("nil error")
		}
		if appErr, ok := err.(*e.AppError); ok {
			t.Log(appErr.Debug())
		} else {
			t.Log(err.Error())
		}

	}
}

func expectErr(t *testing.T, err error, errs []error) bool {
	for _, e := range errs {
		if errors.Is(e, err) {
			return true
		}
	}
	printErrs(t, errs)
	return false
}

func expectErrCount(count int, errs []error) bool {
	return len(errs) == count
}
