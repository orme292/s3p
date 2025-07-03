package auth

import (
    "fmt"

    "github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
    Use:   "add filename",
    Short: "Add an auth profile to the library.",
    Long:  "Adds an auth profile to the library.",
    Run:   useAdd,
}

func useAdd(cmd *cobra.Command, args []string) {
    fmt.Println("auth.add called")
}
