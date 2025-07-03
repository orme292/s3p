package plan

import (
    "fmt"

    "github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
    Use:   "list",
    Short: "List stored plans.",
    Long:  "Display a list of the plans stored in the library.",
    Run:   useList,
}

func useList(cmd *cobra.Command, args []string) {
    fmt.Println("plan.list called")
}
