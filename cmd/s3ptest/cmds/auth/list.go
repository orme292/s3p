package auth

import (
    "fmt"

    "github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
    Use:   "list",
    Short: "List stored auth profiles.",
    Long:  "Display a list of the auth profiles stored in the library.",
    Run:   useList,
}

func useList(cmd *cobra.Command, args []string) {
    fmt.Println("auth.list called")
}
