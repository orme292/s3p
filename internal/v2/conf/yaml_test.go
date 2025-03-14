package conf

import (
    "fmt"
    "os"
    "path/filepath"
    "testing"

    "github.com/stretchr/testify/require"
)

func TestYamlConfig_WriteLoad(t *testing.T) {
    y := T_GetYaml()

    tempPath := getTempPath()
    err := y.WriteToFile(tempPath)
    defer func(t *testing.T, path string) {
        err := deleteTempFile(path)
        if err != nil {
            t.Errorf("error deleting temp file %s", path)
        }
    }(t, tempPath)
    if err != nil {
        require.Nil(t, err, fmt.Sprintf("Shouldn't return error: %v", err))
    }

    // Read and print the file contents
    content, err := os.ReadFile(tempPath)
    require.Nil(t, err, "Failed to read file contents")
    t.Logf("File contents:\n%s\n", string(content))
    t.Logf("Length: %d\n", len(string(content)))

    y = NewYamlConfig()
    err = y.LoadFromFile(tempPath)
    require.Nil(t, err, fmt.Sprintf("Shouldn't return error: %v", err))
}

// meh
func Test_expandHome(t *testing.T) {
    path := expandHome("~/")
    require.Equal(t, filepath.Clean(os.Getenv("HOME")),
        filepath.Clean(path))
}

func Test_writeYamlFile(t *testing.T) {
    f := &os.File{}
    err := writeYamlFile(f, []byte("aabbccddeeff"))
    require.Error(t, err)

}
