package cmd

import (
	"fmt"
	"os"

	"github.com/moabdelazem/initiator/internal/projects"
	"github.com/moabdelazem/initiator/internal/utils"
	"github.com/spf13/cobra"
)

var targetDir string = "." // if not provided, default to current directory
var initGit bool = true    // create git repository by default

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [project-name]",
	Short: "create new project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		// Validate the project name before proceeding
		if err := utils.ValidateProjectName(projectName); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Get The Target Directory And Get The absolute path
		path, err := utils.GetAbsPath(targetDir, projectName)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		// Create The Project Directory with git init option
		if err := utils.CreateProjectDir(path, 0755, initGit); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		// Print Success Message
		fmt.Printf("Project '%s' created successfully at: %s\n", projectName, path)

		// Get The Project Type From The User
		projectType := projects.PromptUserForProjectType()

		// initialize the project
		project := projects.NewProject(projectName, path, projectType)

		// Create The Project
		if err := project.Create(); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// -d flag to specify the parent directory
	createCmd.Flags().StringVarP(&targetDir, "dir", "d", ".", "parent directory for the project")
	// -ng flag to disable git init
	createCmd.Flags().BoolVarP(&initGit, "no-git", "", true, "do not initialize a git repository")
}
