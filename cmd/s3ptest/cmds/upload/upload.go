package upload

import (
    "fmt"

    "github.com/spf13/cobra"
)

var UploadCmd = &cobra.Command{
    GroupID: "commands",
    Use:     "upload",
    Short:   "Upload files.",
    Long:    "Upload files to an object storage service.",
    Run:     useUpload,
}

func GetUploadCmd() *cobra.Command {
    UploadCmd.Flags().StringP("path", "p", "", "local path to upload")
    UploadCmd.Flags().BoolP("recursive", "r", false, "process subdirectories")
    UploadCmd.Flags().StringP("ignore", "i", "", "regular expression, ignore files that match")
    UploadCmd.Flags().String("remote-prefix", "/", "prefix for remote directory")
    UploadCmd.Flags().BoolP("overwrite", "x", false, "overwrite existing files on remote")
    UploadCmd.Flags().BoolP("skip-empty", "k", true, "skip empty files and directories")
    UploadCmd.Flags().StringP("auth", "a", "", "auth profile to authenticate with")
    UploadCmd.Flags().BoolP("new-session-log", "l", false, "generate a new log file for this session")

    return UploadCmd
}

func useUpload(cmd *cobra.Command, args []string) {
    fmt.Println("upload called")
}
