package download

import (
    "fmt"

    "github.com/spf13/cobra"
)

var DownloadCmd = &cobra.Command{
    GroupID: "commands",
    Use:     "download",
    Short:   "Download files.",
    Long:    "Download files from a remote object storage service.",
    Run:     useDownload,
}

func GetDownloadCmd() *cobra.Command {
    DownloadCmd.Flags().StringP("path", "p", "", "local path to download to")
    DownloadCmd.Flags().BoolP("recursive", "r", false, "process subdirectories")
    DownloadCmd.Flags().StringP("ignore", "i", "", "regular expression, ignore files that match")
    DownloadCmd.Flags().String("remote-prefix", "/", "prefix for remote directory to download from")
    DownloadCmd.Flags().BoolP("overwrite", "x", false, "overwrite existing local files")
    DownloadCmd.Flags().BoolP("skip-empty", "k", true, "skip empty files and directories")
    DownloadCmd.Flags().StringP("auth", "a", "", "auth profile to authenticate with")
    DownloadCmd.Flags().BoolP("new-session-log", "l", false, "generate a new log file for this session")

    return DownloadCmd
}

func useDownload(cmd *cobra.Command, args []string) {
    fmt.Println("download called")
}
