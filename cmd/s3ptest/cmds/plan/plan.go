package plan

import (
    "fmt"

    "github.com/spf13/cobra"
)

var PlanCmd = &cobra.Command{
    GroupID: "config",
    Use:     "plan",
    Short:   "Manage backup plans.",
    Long:    "Add, remove, list, and describe stored backup plans.",
    Run:     usePlan,
}

func GetPlanCmd() *cobra.Command {
    PlanCmd.AddCommand(ListCmd)
    PlanCmd.AddCommand(DescribeCmd)
    PlanCmd.AddCommand(AddCmd)

    return PlanCmd
}

func usePlan(cmd *cobra.Command, args []string) {
    fmt.Println("plan called")
}
