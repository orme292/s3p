package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	GroupID: "commands",
	Use:     "sync",
	Short:   "Synchronize local and remote files and folder structures.",
	Long:    "Synchronize local and remote files and folder structures with support for bi-directional or uni-directional syncing.",
	Run:     useSync,
}

func useSync(cmd *cobra.Command, args []string) {
	fmt.Println("sync called")
}
