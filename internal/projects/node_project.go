package projects

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/moabdelazem/initiator/internal/utils"
)

// NodeProject represents a Node.js project.
const typeScript = true

// NodeProject represents a Node.js project.
type NodeProject struct {
	Name string
	Dir  string
}

// Create initializes a new Node.js project in the specified directory.
// It first changes to the project directory, then runs 'npm init -y' to create
// a new Node.js Typescript project with default settings.
// Returns an error if directory change fails or if npm initialization fails.
func (p *NodeProject) Create() error {
	// Define color functions
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	fmt.Printf("\n📦 Creating new TypeScript project: %s\n\n", cyan(p.Name))

	// Navigate to the project directory
	if err := ChangeDirectory(p.Dir); err != nil {
		return fmt.Errorf("%s Failed to access directory: %v", red("✘"), err)
	}

	// Define project setup steps
	steps := []ProjectSteps{
		{
			Name: "Initialize Node.js project",
			Action: func() error {
				cmd := exec.Command("npm", "init", "-y")
				return cmd.Run()
			},
			Message: "Node.js project initialized",
		},
		{
			Name: "Setup TypeScript",
			Action: func() error {
				return p.setupTypeScriptProject()
			},
			Message: "TypeScript configuration completed",
		},
		{
			Name: "Install additional packages",
			Action: func() error {
				return p.installPackages()
			},
			Message: "Additional packages installed",
		},
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

	// Print success message with project information
	fmt.Printf("\n%s Project created successfully!\n\n", green("✨"))
	p.printProjectInfo()
	return nil
}

// printProjectInfo prints the project information to the console.
// It displays the project name, location, and type (TypeScript).
// It also provides next steps for the user to run the project.
// The project information is displayed in cyan color.
func (p *NodeProject) printProjectInfo() {
	// Define color functions
	cyan := color.New(color.FgCyan).SprintFunc()

	// Print project information
	info := []struct {
		label string
		value string
	}{
		{"Project Name", p.Name},
		{"Location", p.Dir},
		{"Type", "TypeScript"},
	}

	// Print project information
	fmt.Println("Project Information:")
	fmt.Println(strings.Repeat("-", 40))
	for _, item := range info {
		fmt.Printf("%-12s: %s\n", item.label, cyan(item.value))
	}
	fmt.Println(strings.Repeat("-", 40))

	// Print next steps
	fmt.Println("\nNext steps:")
	fmt.Printf("  cd %s\n", cyan(p.Name))
	fmt.Printf("  %s\n", cyan("npm run dev"))
	fmt.Printf("  %s\n", cyan("npm run build"))
	fmt.Println()
}

// setupTypeScriptProject configures a new TypeScript project by:
// - Creating a tsconfig.json file with standard TypeScript configuration
// - Setting up the project directory structure with a src folder
// - Creating an initial index.ts file with basic content
// - Installing required TypeScript dependencies (typescript, @types/node, ts-node)
//
// Returns an error if any step in the setup process fails, such as:
// - Failed to create config files or directories
// - Failed to install npm dependencies
func (p *NodeProject) setupTypeScriptProject() error {
	// Create tsconfig.json file
	tsConfig := `{
  "compilerOptions": {
	"target": "es6",
	"module": "commonjs",
	"outDir": "./dist",
	"rootDir": "./src",
	"strict": true,
	"esModuleInterop": true,
	"skipLibCheck": true,
	"forceConsistentCasingInFileNames": true
  }
}`
	if err := os.WriteFile("tsconfig.json", []byte(tsConfig), 0644); err != nil {
		return fmt.Errorf("failed to create tsconfig.json: %v", err)
	}

	// Create project structure
	if err := os.MkdirAll("src", 0755); err != nil {
		return fmt.Errorf("failed to create src directory: %v", err)
	}

	// Create index.ts file
	indexContent := `console.log('Hello from TypeScript!');`
	if err := os.WriteFile("src/index.ts", []byte(indexContent), 0644); err != nil {
		return fmt.Errorf("failed to create index.ts: %v", err)
	}

	// Install TypeScript dependencies with spinner
	s := utils.CreateSpinner("Installing TypeScript dependencies...")
	s.Start()
	cmd := exec.Command("npm", "install", "--save-dev", "typescript", "@types/node", "ts-node")
	err := cmd.Run()
	s.Stop()
	if err != nil {
		return fmt.Errorf("failed to install TypeScript dependencies: %v", err)
	}

	// Update package.json with scripts
	scripts := `{
  "scripts": {
    "start": "node dist/index.js",
    "dev": "ts-node src/index.ts",
    "build": "tsc",
    "watch": "tsc -w"
  }
}`

	// Read existing package.json
	content, err := os.ReadFile("package.json")
	if err != nil {
		return fmt.Errorf("failed to read package.json: %v", err)
	}

	// Parse and update package.json with new scripts
	pkgJson := string(content)
	pkgJson = strings.Replace(pkgJson, `"scripts": {
    "test": "echo \\"Error: no test specified\\" && exit 1"
  }`, scripts, 1)

	// Write updated package.json
	if err := os.WriteFile("package.json", []byte(pkgJson), 0644); err != nil {
		return fmt.Errorf("failed to update package.json: %v", err)
	}

	return nil
}

// TODO: Implement this function
func (p *NodeProject) installPackages() error {
	// Add loading indicators to package installation logic when implemented
	return nil
}
