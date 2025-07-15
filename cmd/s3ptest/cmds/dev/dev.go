package dev

import (
    "fmt"

    "github.com/spf13/cobra"
)

var DevCmd = &cobra.Command{
    GroupID: "dev",
    Use:     "dev",
    Short:   "Development.",
    Long:    "Development Options.",
    Run:     useDev,
}

func GetDevCmd() *cobra.Command {
    DevCmd.AddCommand(getCertsCmd())

    return DevCmd
}

func useDev(cmd *cobra.Command, args []string) {
    fmt.Println("dev called")
}
