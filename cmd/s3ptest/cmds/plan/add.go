package plan

import (
    "fmt"

    "github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
    Use:   "add filename",
    Short: "Add a plan profile to the library.",
    Long:  "Adds a plan profile to the library.",
    Run:   useAdd,
}

func useAdd(cmd *cobra.Command, args []string) {
    fmt.Println("plan.add called")
}
