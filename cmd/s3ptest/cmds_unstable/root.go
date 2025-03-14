package cmds_unstable

import (
    "os"

    "github.com/spf13/cobra"
)

const (
    UseProfileFilenameFlag  = "filename"
    UseProfileFilenameFlagS = "f"
)

var rootCmd = &cobra.Command{
    Use:   "s3ptest",
    Short: "s3ptest is a test tool for s3p",
    Long:  "s3ptest is a test tool for s3p with unstable features",
}

func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}

func init() {
    rootCmd.CompletionOptions.HiddenDefaultCmd = true

    addUseCmd()

    UseCmd.Flags().StringP(UseProfileFilenameFlag, UseProfileFilenameFlagS, "./sample-profile.yml", "filename for the new profile")
    _ = UseCmd.MarkFlagRequired(UseProfileFilenameFlag)
}
