package cmds

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "s3p/cmd/s3ptest/cmds/auth"
    "s3p/cmd/s3ptest/cmds/dev"
    "s3p/cmd/s3ptest/cmds/download"
    "s3p/cmd/s3ptest/cmds/plan"
    "s3p/cmd/s3ptest/cmds/sync"
    "s3p/cmd/s3ptest/cmds/upload"
    "s3p/cmd/s3ptest/utils"
    "s3p/internal/inits"
)

var rootCmd = &cobra.Command{
    Use:   "s3ptest",
    Short: "s3ptest is a tool used to interact with a remote object storage service",
    Long: "s3ptest is a tool used to interact with remote object storage services. It supports uni and bi-directional" +
        " uploads and downloads, syncing, and backup plans.",
}

func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
}

func addCommands() {
    rootCmd.AddGroup(&cobra.Group{
        ID:    "config",
        Title: "Configuration:",
    })
    rootCmd.AddCommand(auth.GetAuthCmd())
    rootCmd.AddCommand(plan.GetPlanCmd())

    rootCmd.AddGroup(&cobra.Group{
        ID:    "commands",
        Title: "Storage Commands:",
    })
    rootCmd.AddCommand(download.GetDownloadCmd())
    rootCmd.AddCommand(sync.GetSyncCmd())
    rootCmd.AddCommand(upload.GetUploadCmd())
}

func addDev() {
    rootCmd.AddGroup(&cobra.Group{
        ID:    "dev",
        Title: "Development Commands:",
    })
    rootCmd.AddCommand(dev.GetDevCmd())
}

func addFlags() {
    rootCmd.PersistentFlags().Bool("debug", false, "enable debug mode")
}

func init() {
    // This hides the default completion option provided by cobra
    rootCmd.CompletionOptions.HiddenDefaultCmd = true

    addCommands()
    addFlags()

    // Starts a goroutine that listens for interrupts (ctrl+c)
    utils.SigIntListener()

    // Display the program title
    fmt.Printf("\033[1;36m%s\033[0m\n\n", "s3p - a multi-service object storage tool")

    // Load INI file
    cfg := inits.Retrieve()

    if cfg.DevMode {
        addDev()
    }
}
