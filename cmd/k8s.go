package cmd

import (
	"fmt"
	"os"

	"github.com/moabdelazem/initiator/internal/k8s"
	"github.com/moabdelazem/initiator/internal/utils"
	"github.com/spf13/cobra"
)

var (
	namespace     string = "default"
	createService bool   = false
	createIngress bool   = false
	port          int    = 8080
	outputDir     string = "."
	containerName string = "" // New variable for container name
	projectName   string = "" // New variable for project name
)

// k8sCmd represents the k8s command
var k8sCmd = &cobra.Command{
	Use:   "k8s [app-name]",
	Short: "Generate Kubernetes manifests",
	Long: `Generate Kubernetes manifests for your application.
Optionally include service and ingress resources.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appName := args[0]

		// Validate application name
		if err := utils.ValidateProjectName(appName); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// If container name is not provided, use app name
		if containerName == "" {
			containerName = appName
		}

		// If project name is not provided, use app name
		if projectName == "" {
			projectName = appName
		}

		// Get the output directory path
		path, err := utils.GetAbsPath(outputDir, "")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		// Create k8s manifest generator with separate container and project names
		generator := k8s.NewManifestGenerator(appName, projectName, containerName, namespace, port, createService, createIngress)

		// Generate the manifests
		if err := generator.Generate(path); err != nil {
			fmt.Printf("Error generating Kubernetes manifests: %v\n", err)
			return
		}

		fmt.Printf("Kubernetes manifests for '%s' generated successfully at: %s\n", appName, path)
	},
}

func init() {
	rootCmd.AddCommand(k8sCmd)

	// Add flags for k8s command
	k8sCmd.Flags().StringVarP(&namespace, "namespace", "n", "default", "Kubernetes namespace for the application")
	k8sCmd.Flags().BoolVarP(&createService, "service", "s", false, "Create a Kubernetes Service manifest")
	k8sCmd.Flags().BoolVarP(&createIngress, "ingress", "i", false, "Create a Kubernetes Ingress manifest")
	k8sCmd.Flags().IntVarP(&port, "port", "p", 8080, "Container port for the application")
	k8sCmd.Flags().StringVarP(&outputDir, "output", "o", ".", "Output directory for the manifest files")
	k8sCmd.Flags().StringVarP(&containerName, "container-name", "c", "", "Container name (defaults to app-name if not provided)")
	k8sCmd.Flags().StringVarP(&projectName, "project-name", "r", "", "Project name for labels and selectors (defaults to app-name if not provided)")
}
