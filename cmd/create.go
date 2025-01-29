package cmd

import (
	"fmt"

	"github.com/moabdelazem/initiator/internal/utils"
	"github.com/spf13/cobra"
)

var targetDir string = "." // if not provided, default to current directory

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [project-name]",
	Short: "create new project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		// Get The Target Directory And Get The absolute path
		path, err := utils.GetAbsPath(targetDir, projectName)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		// Create The Project Directory
		if err := utils.CreateProjectDir(path, 0755); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		// Print Success Message
		fmt.Printf("Project '%s' created successfully at: %s\n", projectName, path)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	createCmd.Flags().StringVarP(&targetDir, "dir", "d", ".", "parent directory for the project")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
