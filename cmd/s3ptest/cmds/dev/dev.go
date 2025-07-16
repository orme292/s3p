package dev

import (
    "github.com/spf13/cobra"
)

var DevCmd = &cobra.Command{
    GroupID: "dev",
    Use:     "dev",
    Short:   "Development.",
    Long:    "Development Options.",
}

func GetDevCmd() *cobra.Command {
    DevCmd.AddCommand(getCertsCmd())

    return DevCmd
}
