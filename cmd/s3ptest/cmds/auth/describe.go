package auth

import (
    "fmt"

    "github.com/spf13/cobra"
)

var DescribeCmd = &cobra.Command{
    Use:   "describe",
    Short: "Describe a stored auth profile.",
    Long:  "Display the contents of a stored auth profile.",
    Run:   useDescribe,
}

func useDescribe(cmd *cobra.Command, args []string) {
    fmt.Println("auth.describe called")
}
