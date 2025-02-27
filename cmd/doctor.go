package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/moabdelazem/initiator/pkg"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check the required dependencies for the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		headerColor := color.New(color.FgCyan, color.Bold)
		successPrefix := color.GreenString("✓")
		errorPrefix := color.RedString("✗")

		headerColor.Println("Checking required dependencies...")

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Dependency", "Status", "Version", "Installation URL"})
		table.SetBorder(true)
		table.SetRowLine(true)
		table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_CENTER, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT})
		table.SetHeaderAlignment(tablewriter.ALIGN_CENTER)
		table.SetAutoWrapText(false)

		// Check Go
		checkDependency(table, "Go", pkg.IsGoInstalled, pkg.GetGoVersion, "https://golang.org/dl/", successPrefix, errorPrefix)

		// Check Node.js
		checkDependency(table, "Node.js", pkg.IsNodeInstalled, pkg.GetNodeVersion, "https://nodejs.org/", successPrefix, errorPrefix)

		// Check Git
		checkDependency(table, "Git", pkg.IsGitInstalled, pkg.GetGitVersion, "https://git-scm.com/downloads", successPrefix, errorPrefix)

		// Check Docker
		checkDependency(table, "Docker", pkg.IsDockerInstalled, pkg.GetDockerVersion, "https://www.docker.com/get-started", successPrefix, errorPrefix)

		// Check Docker Compose
		checkDependency(table, "Docker Compose", pkg.IsDockerComposeInstalled, pkg.GetDockerComposeVersion, "https://docs.docker.com/compose/install/", successPrefix, errorPrefix)

		table.Render()
	},
}

// checkDependency checks a single dependency and adds its status to the table
func checkDependency(
	table *tablewriter.Table,
	name string,
	isInstalledFn func() bool,
	getVersionFn func() (string, error),
	installUrl string,
	successPrefix string,
	errorPrefix string,
) {
	if isInstalledFn() {
		version, err := getVersionFn()
		if err != nil {
			table.Append([]string{name, successPrefix + " Installed", "Version unknown", ""})
		} else {
			table.Append([]string{name, successPrefix + " Installed", version, ""})
		}
	} else {
		table.Append([]string{name, errorPrefix + " Not installed", "N/A", installUrl})
	}
}
