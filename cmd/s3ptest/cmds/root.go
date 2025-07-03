package cmds

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "s3p/cmd/s3ptest/cmds/auth"
    "s3p/cmd/s3ptest/cmds/plan"
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
    rootCmd.AddCommand(downloadCmd)
    rootCmd.AddCommand(syncCmd)
    rootCmd.AddCommand(uploadCmd)
}

func addFlags() {
    return
}

func init() {
    // This hides the default completion option provided by cobra
    rootCmd.CompletionOptions.HiddenDefaultCmd = true

    addCommands()

    // Starts a goroutine that listens for interrupts (ctrl+c)
    utils.SigIntListener()

    fmt.Printf("\033[1;36m%s\033[0m\n\n", "s3p -a multi-service object storage tool")

    // Load INI file
    _ = inits.Retrieve()
}
