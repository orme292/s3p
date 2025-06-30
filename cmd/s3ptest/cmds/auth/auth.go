package auth

import (
	"fmt"

	"github.com/spf13/cobra"
)

var AuthCmd = &cobra.Command{
	GroupID: "config",
	Use:     "auth",
	Short:   "Manage the stored auth profiles.",
	Long:    "Add, remove, list, and describe the stored auth profiles.",
	Run:     useAuth,
}

func useAuth(cmd *cobra.Command, args []string) {
	fmt.Println("auth called")
}
