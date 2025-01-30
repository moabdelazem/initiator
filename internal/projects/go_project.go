package projects

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/moabdelazem/initiator/internal/utils"
)

// GoProjectType represents the type of Go project.
type GoProjectType string

const (
	PlainGo GoProjectType = "plain"
	WebGo   GoProjectType = "web"
)

// GoProject represents a Go project.
type GoProject struct {
	Name        string
	Dir         string
	ProjectType GoProjectType
}

// Create initializes a new Go project in the specified directory.
// It changes to the project directory and runs 'go mod init' with the project name.
// Returns an error if directory change fails or if project initialization fails.
func (p *GoProject) Create() error {
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	fmt.Printf("\n🚀 Creating new Go project: %s\n\n", cyan(p.Name))

	// Ask for project type if not set
	if p.ProjectType == "" {
		p.ProjectType = p.promptForGoProjectType()
	}

	if err := ChangeDirectory(p.Dir); err != nil {
		return fmt.Errorf("%s Failed to access directory: %v", red("✘"), err)
	}

	steps := []ProjectSteps{
		{
			Name: "Initialize Go module",
			Action: func() error {
				cmd := exec.Command("go", "mod", "init", p.Name)
				return cmd.Run()
			},
			Message: "Go module initialized",
		},
		{
			Name: "Setup project structure",
			Action: func() error {
				return p.setupProjectStructure()
			},
			Message: "Project structure created",
		},
		{
			Name: "Create main package",
			Action: func() error {
				if p.ProjectType == WebGo {
					return p.createWebPackage()
				}
				return p.createMainPackage()
			},
			Message: "Main package created",
		},
		{
			Name: "Tidy Things Up",
			Action: func() error {
				cmd := exec.Command("go", "mod", "tidy")
				return cmd.Run()
			},
			Message: "Go modules tidied",
		},
	}

	// Add web-specific steps
	if p.ProjectType == WebGo {
		steps = append(steps, ProjectSteps{
			Name: "Install web dependencies",
			Action: func() error {
				return p.installWebDependencies()
			},
			Message: "Web dependencies installed",
		})
	}

	// Execute project setup steps
	for _, step := range steps {
		s := utils.CreateSpinner(step.Name + "...")
		s.Start()
		if err := step.Action(); err != nil {
			s.Stop()
			return fmt.Errorf("%s %s: %v", red("✘"), step.Name, err)
		}
		s.Stop()
		fmt.Printf("%s %s\n", green("✓"), step.Message)
	}

	fmt.Printf("\n%s Project created successfully!\n\n", green("✨"))
	p.printProjectInfo()
	return nil
}

// promptForGoProjectType prompts the user to select a Go project type.
// It displays a list of predefined project types with descriptions and asks the user
// to select one by entering a number.
//
// The function reads user input from standard input and validates it to ensure a valid
// project type is selected. It keeps prompting until a valid choice is made.
//
// Returns: the selected GoProjectType.
func (p *GoProject) promptForGoProjectType() GoProjectType {
	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	white := color.New(color.FgWhite, color.Bold).SprintFunc()

	options := []struct {
		Type        GoProjectType
		Name        string
		Description string
	}{
		{
			Type:        PlainGo,
			Name:        "Plain Go Project",
			Description: "Basic Go project with standard structure",
		},
		{
			Type:        WebGo,
			Name:        "Web Project",
			Description: "Go web project with Echo framework, middleware, and API structure",
		},
	}

	fmt.Printf("\n%s Select Go project type:\n\n", white("📋"))

	for i, opt := range options {
		fmt.Printf("%s %s\n", cyan(fmt.Sprintf("%d.", i+1)), opt.Name)
		fmt.Printf("   %s\n", yellow(opt.Description))
	}

	fmt.Printf("\n%s Enter your choice (1-%d): ", white("→"), len(options))

	var choice int
	for {
		fmt.Scanln(&choice)
		if choice >= 1 && choice <= len(options) {
			fmt.Printf("%s Selected: %s\n\n", white("✓"), cyan(options[choice-1].Name))
			return options[choice-1].Type
		}
		fmt.Printf("%s Please enter a number between 1 and %d: ", yellow("!"), len(options))
	}
}

// setupProjectStructure creates the initial directory structure for a Go project.
// For basic Go projects, it creates the standard layout directories:
// - cmd/: Contains main application entry points
// - internal/: Private application code
// - pkg/: Code that's safe to use by external applications
// - docs/: Documentation files
// - test/: Additional test files
//
// For web projects (when ProjectType is WebGo), it creates additional directories:
// - internal/handlers/: HTTP request handlers
// - internal/middleware/: HTTP middleware components
// - internal/models/: Data models
// - internal/routes/: Route definitions
// - internal/services/: Business logic services
//
// It also generates and writes a README.md file based on the project type.
//
// Returns an error if directory creation or README file writing fails.
func (p *GoProject) setupProjectStructure() error {
	// Define standard project directories
	dirs := []string{
		"cmd",
		"internal",
		"pkg",
		"docs",
		"test",
	}

	// Add web-specific directories
	if p.ProjectType == WebGo {
		dirs = append(dirs,
			"internal/handlers",
			"internal/middleware",
			"internal/models",
			"internal/routes",
			"internal/services",
		)
	}

	// Loop through directories and create them
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create %s directory: %v", dir, err)
		}
	}

	// Update README based on project type
	readmeContent := p.generateReadme()
	if err := os.WriteFile("README.md", []byte(readmeContent), 0644); err != nil {
		return fmt.Errorf("failed to create README.md: %v", err)
	}

	return nil
}

// createMainPackage creates a new main.go file in the cmd directory with a basic
// Go application structure. The created file contains a simple main package with
// basic imports (fmt, log) and a main function that prints startup messages.
// It returns an error if the file creation fails.
func (p *GoProject) createMainPackage() error {
	mainContent := `package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Starting application...")
	fmt.Println("Hello from Go!")
}
`

	if err := os.WriteFile("cmd/main.go", []byte(mainContent), 0644); err != nil {
		return fmt.Errorf("failed to create main.go: %v", err)
	}

	return nil
}

// createWebPackage initializes a basic web application structure by creating a main.go file
// in the cmd directory. The generated file sets up an Echo web server with basic middleware
// (Logger and Recover) and a single "/" route that returns a JSON welcome message.
// The server listens on port 8080.
//
// The function writes the content to cmd/main.go with 0644 permissions.
//
// Returns an error if file creation fails.
func (p *GoProject) createWebPackage() error {
	mainContent := `package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Welcome to the API!",
	})
}
`

	if err := os.WriteFile("cmd/main.go", []byte(mainContent), 0644); err != nil {
		return fmt.Errorf("failed to create main.go: %v", err)
	}

	return nil
}

// installWebDependencies installs required web development dependencies using go get
// and creates a default .env file with basic configuration.
//
// The function installs the following dependencies:
// - github.com/labstack/echo/v4 - Web framework
// - github.com/joho/godotenv - Environment variable loader
//
// The created .env file contains default configurations for:
// - Server settings (port, environment)
// - Database connection parameters
//
// Returns an error if dependency installation fails or if .env file creation fails.
func (p *GoProject) installWebDependencies() error {
	deps := []string{
		"github.com/labstack/echo/v4",
		"github.com/joho/godotenv",
	}

	for _, dep := range deps {
		cmd := exec.Command("go", "get", dep)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to install %s: %v", dep, err)
		}
	}

	// Create .env file
	envContent := `# Server Configuration
PORT=8080
ENV=development

# Database Configuration (if needed)
DB_HOST=localhost
DB_PORT=5432
DB_NAME=dbname
DB_USER=user
DB_PASSWORD=password
`
	if err := os.WriteFile(".env", []byte(envContent), 0644); err != nil {
		return fmt.Errorf("failed to create .env file: %v", err)
	}

	return nil
}

// generateReadme generates a README.md file content based on project type.
// For web projects (WebGo), it generates a comprehensive README with web-specific
// structure and instructions including API endpoints.
// For other project types, it generates a basic README with standard Go project structure.
//
// Returns:
//   - string: The complete README.md content with project name and instructions
func (p *GoProject) generateReadme() string {
	if p.ProjectType == WebGo {
		return fmt.Sprintf(`# %s

## Description
A Go web application with Echo framework and modern project structure.

## Project Structure
- cmd/: Main applications
- internal/
  - handlers/: HTTP request handlers
  - middleware/: Custom middleware
  - models/: Data models
  - routes/: Route definitions
  - services/: Business logic
- pkg/: Library code
- docs/: Documentation
- test/: Tests

## Getting Started
1. Install dependencies:
   ~~~
   go mod tidy
   ~~~

2. Configure environment:
   ~~~
   cp .env.example .env
   ~~~

3. Run the server:
   ~~~
   go run ./cmd/main.go
   ~~~

## API Endpoints
- GET /: Welcome message
`, p.Name)
	}

	// Return plain project README
	return fmt.Sprintf(`# %s

## Description
A Go project created with modern project structure.

## Project Structure
- cmd/: Main applications
- internal/: Private application code
- pkg/: Library code
- docs/: Documentation
- test/: Tests

## Getting Started
1. Build the project:
   ~~~
   go build ./cmd/...
   ~~~

2. Run the application:
   ~~~
   go run ./cmd/main.go
   ~~~
`, p.Name)
}

// printProjectInfo prints formatted information about the Go project.
// It displays project details including:
//   - Project name
//   - Project directory location
//   - Project type
//
// After displaying the project information, it shows next steps for the user
// including commands to:
//   - Change to the project directory
//   - Run the project
//   - Build the project
//
// The output is formatted with colored text using cyan for values and
// white bold for headers.
func (p *GoProject) printProjectInfo() {
	cyan := color.New(color.FgCyan).SprintFunc()
	white := color.New(color.FgWhite, color.Bold).SprintFunc()

	info := []struct {
		label string
		value string
	}{
		{"Project Name", p.Name},
		{"Location", p.Dir},
		{"Type", "Go"},
	}

	fmt.Println("Project Information:")
	fmt.Println(strings.Repeat("-", 40))
	for _, item := range info {
		fmt.Printf("%-12s: %s\n", item.label, cyan(item.value))
	}
	fmt.Println(strings.Repeat("-", 40))

	fmt.Printf("\n%s Next steps:\n", white("→"))
	fmt.Printf("  cd %s\n", cyan(p.Name))
	fmt.Printf("  %s\n", cyan("go run ./cmd/main.go"))
	fmt.Printf("  %s\n", cyan("go build ./cmd/..."))
	fmt.Println()
}
