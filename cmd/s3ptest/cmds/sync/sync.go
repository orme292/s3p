package sync

import (
    "fmt"

    "github.com/spf13/cobra"
)

var SyncCmd = &cobra.Command{
    GroupID: "commands",
    Use:     "sync",
    Short:   "Sync files.",
    Long:    "Sync files between this device and a remote object storage service.",
    Run:     useSync,
}

func GetSyncCmd() *cobra.Command {
    SyncCmd.Flags().StringP("path", "p", "", "local path to sync with")
    SyncCmd.Flags().BoolP("recursive", "r", false, "process subdirectories")
    SyncCmd.Flags().StringP("ignore", "i", "", "regular expression, ignore files that match")
    SyncCmd.Flags().String("remote-prefix", "/", "prefix for remote directory to sync with")
    SyncCmd.Flags().BoolP("overwrite", "x", false, "overwrite existing local and remote files")
    SyncCmd.Flags().BoolP("skip-empty", "k", true, "skip empty files and directories")
    SyncCmd.Flags().StringP("auth", "a", "", "auth profile to authenticate with")
    SyncCmd.Flags().BoolP("new-session-log", "l", false, "generate a new log file for this session")

    return SyncCmd
}

func useSync(cmd *cobra.Command, args []string) {
    fmt.Println("sync called")
}
