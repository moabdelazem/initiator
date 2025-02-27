package projects

import (
	"os/exec"
)

// NodeProjectType represents the type of Node.js project
type NodeProjectType string

const (
	// Node.js project types
	TypeScriptBasic NodeProjectType = "typescript-basic"
	NextJS          NodeProjectType = "nextjs"
	Remix           NodeProjectType = "remix"
	Express         NodeProjectType = "express"
	NestJS          NodeProjectType = "nestjs"
)

// NodeProjectOption represents a Node.js project option in the selection menu
type NodeProjectOption struct {
	Type        NodeProjectType
	Name        string
	Description string
}

// GetNodeProjectOptions returns all available Node.js project options
func GetNodeProjectOptions() []NodeProjectOption {
	return []NodeProjectOption{
		{
			Type:        TypeScriptBasic,
			Name:        "TypeScript Basic",
			Description: "A simple TypeScript project with minimal configuration",
		},
		{
			Type:        NextJS,
			Name:        "Next.js",
			Description: "React framework with server-side rendering and static site generation",
		},
		{
			Type:        Remix,
			Name:        "Remix",
			Description: "Full stack web framework focusing on web standards and modern UX",
		},
		{
			Type:        Express,
			Name:        "Express",
			Description: "Fast, unopinionated, minimalist web framework for Node.js",
		},
		{
			Type:        NestJS,
			Name:        "NestJS",
			Description: "Progressive Node.js framework for building server-side applications",
		},
	}
}

// SetupNextJS configures a Next.js project
func SetupNextJS(name string) error {
	cmd := exec.Command("npx", "create-next-app@latest", ".", "--typescript", "--eslint", "--tailwind", "--app", "--src-dir", "--import-alias", "@/*", "--use-npm")
	cmd.Env = append(cmd.Environ(), "npm_config_yes=true")
	return cmd.Run()
}

// SetupRemix configures a Remix project
func SetupRemix(name string) error {
	cmd := exec.Command("npx", "create-remix@latest", ".", "--typescript", "--install")
	cmd.Env = append(cmd.Environ(), "npm_config_yes=true")
	return cmd.Run()
}

// SetupExpress configures an Express.js project with TypeScript
func SetupExpress(name string) error {
	// Install dependencies
	installCmd := exec.Command("npm", "install", "express", "@types/express", "--save")
	if err := installCmd.Run(); err != nil {
		return err
	}

	// We'll rely on the standard TypeScript setup function to create the project structure
	// Additional Express-specific files will be created by the caller
	return nil
}

// SetupNestJS configures a NestJS project
func SetupNestJS(name string) error {
	cmd := exec.Command("npx", "@nestjs/cli", "new", ".", "--package-manager", "npm", "--language", "ts")
	cmd.Env = append(cmd.Environ(), "npm_config_yes=true")
	return cmd.Run()
}
