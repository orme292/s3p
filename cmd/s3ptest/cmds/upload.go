package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	GroupID: "commands",
	Use:     "upload",
	Short:   "Upload local files",
	Long:    "Upload local files and folder structures to a remote object storage service",
	Run:     useUpload,
}

func useUpload(cmd *cobra.Command, args []string) {
	fmt.Println("upload called")
}
