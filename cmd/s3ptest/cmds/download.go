package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	GroupID: "commands",
	Use:     "download",
	Short:   "Download files from a remote object storage service",
	Long:    "Download file and folder structures from a remote object storage service.",
	Run:     useDownload,
}

func useDownload(cmd *cobra.Command, args []string) {
	fmt.Println("download called")
}
