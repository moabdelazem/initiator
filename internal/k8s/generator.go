package k8s

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// ManifestGenerator generates Kubernetes manifests
type ManifestGenerator struct {
	AppName       string
	ProjectName   string
	ContainerName string
	Namespace     string
	Port          int
	WithService   bool
	WithIngress   bool
}

// NewManifestGenerator creates a new ManifestGenerator
func NewManifestGenerator(appName, projectName, containerName, namespace string, port int, withService, withIngress bool) *ManifestGenerator {
	return &ManifestGenerator{
		AppName:       appName,
		ProjectName:   projectName,
		ContainerName: containerName,
		Namespace:     namespace,
		Port:          port,
		WithService:   withService,
		WithIngress:   withIngress,
	}
}

// Generate generates Kubernetes manifests in the specified directory
func (g *ManifestGenerator) Generate(outputDir string) error {
	// Create manifest directory if it doesn't exist
	k8sDir := filepath.Join(outputDir, "k8s")
	if err := os.MkdirAll(k8sDir, 0755); err != nil {
		return fmt.Errorf("failed to create k8s directory: %v", err)
	}

	// Always generate deployment manifest
	if err := g.generateDeployment(k8sDir); err != nil {
		return err
	}

	// Generate service manifest if requested
	if g.WithService {
		if err := g.generateService(k8sDir); err != nil {
			return err
		}
	}

	// Generate ingress manifest if requested
	if g.WithIngress {
		if err := g.generateIngress(k8sDir); err != nil {
			return err
		}
	}

	return nil
}

func (g *ManifestGenerator) generateDeployment(k8sDir string) error {
	deploymentPath := filepath.Join(k8sDir, "deployment.yaml")
	return g.generateManifest(deploymentPath, deploymentTemplate)
}

func (g *ManifestGenerator) generateService(k8sDir string) error {
	servicePath := filepath.Join(k8sDir, "service.yaml")
	return g.generateManifest(servicePath, serviceTemplate)
}

func (g *ManifestGenerator) generateIngress(k8sDir string) error {
	ingressPath := filepath.Join(k8sDir, "ingress.yaml")
	return g.generateManifest(ingressPath, ingressTemplate)
}

func (g *ManifestGenerator) generateManifest(filePath string, templateContent string) error {
	// Parse template
	tmpl, err := template.New(filepath.Base(filePath)).Parse(templateContent)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	// Create file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", filePath, err)
	}
	defer file.Close()

	// Execute template
	if err := tmpl.Execute(file, g); err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	return nil
}
