package plan

import (
    "fmt"

    "github.com/spf13/cobra"
)

var DescribeCmd = &cobra.Command{
    Use:   "describe plan_name",
    Short: "Describe a stored plan.",
    Long:  "Display the contents of a stored plan.",
    Run:   useDescribe,
}

func useDescribe(cmd *cobra.Command, args []string) {
    fmt.Println("plan.describe called")
}
