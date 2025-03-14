package cmds_unstable

import (
    "fmt"
    "log"

    "github.com/spf13/cobra"
    "s3p/internal/providers/aws"
    "s3p/internal/v2/conf"
)

var UseCmd = &cobra.Command{
    Use:   "use",
    Short: "use an s3p configuration file",
    Long:  "use an s3p configuration file to upload files to object storage services",
    Run:   useCmdFunc,
}

func addUseCmd() {
    rootCmd.AddCommand(UseCmd)
}

func useCmdFunc(cmd *cobra.Command, args []string) {
    filename, err := cmd.Flags().GetString(UseProfileFilenameFlag)
    if err != nil {
        log.Fatalf(fmt.Sprintf("Failed to retrieve '%s' flag: %v", UseProfileFilenameFlag, err))
    }

    y := conf.NewYamlConfig()
    err = y.LoadFromFile(filename)
    if err != nil {
        log.Fatalf(fmt.Sprintf("Failed to load file '%s': %v", filename, err))
    }

    err = y.InjectProviderConfig(aws.NewConfAWS)
    if err != nil {
        log.Fatalf(fmt.Sprintf("Failed to inject provider config: %v", err))
    }

    fmt.Printf("%+v\n", y)
}
