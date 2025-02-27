package cmd

import (
	"fmt"

	"github.com/moabdelazem/initiator/internal/version"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of initiator",
	Long:  `All software has versions. This is initiator's.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initiator CLI Tool")
		fmt.Println(version.GetVersion())
		fmt.Println("Created by @moabdelazem")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
