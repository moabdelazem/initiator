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
type NodeProject struct {
	Name        string
	Dir         string
	ProjectType NodeProjectType
}

// Create initializes a new Node.js project in the specified directory.
// It first changes to the project directory, then sets up the project
// based on the selected project type (TypeScript, Next.js, Remix, etc.)
// Returns an error if directory change fails or if project setup fails.
func (p *NodeProject) Create() error {
	// Define color functions
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	// If project type is not selected, prompt the user
	if p.ProjectType == "" {
		p.ProjectType = promptNodeProjectType()
	}

	fmt.Printf("\n📦 Creating new Node.js project: %s (%s)\n\n",
		cyan(p.Name),
		cyan(p.ProjectType))

	// Navigate to the project directory
	if err := ChangeDirectory(p.Dir); err != nil {
		return fmt.Errorf("%s Failed to access directory: %v", red("✘"), err)
	}

	// Define project setup steps based on project type
	var steps []ProjectSteps

	switch p.ProjectType {
	case TypeScriptBasic:
		steps = p.getTypeScriptSteps()
	case NextJS:
		steps = p.getNextJSSteps()
	case Remix:
		steps = p.getRemixSteps()
	case Express:
		steps = p.getExpressSteps()
	case NestJS:
		steps = p.getNestJSSteps()
	default:
		return fmt.Errorf("%s Unsupported project type: %s", red("✘"), p.ProjectType)
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

// getTypeScriptSteps returns the steps needed to set up a basic TypeScript project
func (p *NodeProject) getTypeScriptSteps() []ProjectSteps {
	return []ProjectSteps{
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
}

// getNextJSSteps returns the steps needed to set up a Next.js project
func (p *NodeProject) getNextJSSteps() []ProjectSteps {
	return []ProjectSteps{
		{
			Name: "Creating Next.js project",
			Action: func() error {
				return SetupNextJS(p.Name)
			},
			Message: "Next.js project created",
		},
	}
}

// getRemixSteps returns the steps needed to set up a Remix project
func (p *NodeProject) getRemixSteps() []ProjectSteps {
	return []ProjectSteps{
		{
			Name: "Creating Remix project",
			Action: func() error {
				return SetupRemix(p.Name)
			},
			Message: "Remix project created",
		},
	}
}

// getExpressSteps returns the steps needed to set up an Express.js project
func (p *NodeProject) getExpressSteps() []ProjectSteps {
	return []ProjectSteps{
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
			Name: "Setup Express",
			Action: func() error {
				return SetupExpress(p.Name)
			},
			Message: "Express.js installed",
		},
		{
			Name: "Create Express starter files",
			Action: func() error {
				// Create a basic Express app
				appContent := `import express from 'express';

const app = express();
const port = process.env.PORT || 3000;

app.get('/', (req, res) => {
  res.send('Hello from Express with TypeScript!');
});

app.listen(port, () => {
  console.log(\"Server running on port \${port}\");
});`
				return os.WriteFile("src/index.ts", []byte(appContent), 0644)
			},
			Message: "Express starter files created",
		},
	}
}

// getNestJSSteps returns the steps needed to set up a NestJS project
func (p *NodeProject) getNestJSSteps() []ProjectSteps {
	return []ProjectSteps{
		{
			Name: "Creating NestJS project",
			Action: func() error {
				return SetupNestJS(p.Name)
			},
			Message: "NestJS project created",
		},
	}
}

// printProjectInfo prints the project information to the console.
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
		{"Type", string(p.ProjectType)},
	}

	// Print project information
	fmt.Println("Project Information:")
	fmt.Println(strings.Repeat("-", 40))
	for _, item := range info {
		fmt.Printf("%-12s: %s\n", item.label, cyan(item.value))
	}
	fmt.Println(strings.Repeat("-", 40))

	// Print next steps based on project type
	fmt.Println("\nNext steps:")
	fmt.Printf("  cd %s\n", cyan(p.Name))

	switch p.ProjectType {
	case TypeScriptBasic, Express:
		fmt.Printf("  %s\n", cyan("npm run dev"))
		fmt.Printf("  %s\n", cyan("npm run build"))
	case NextJS:
		fmt.Printf("  %s\n", cyan("npm run dev"))
		fmt.Printf("  %s\n", cyan("npm run build"))
		fmt.Printf("  %s\n", cyan("npm start"))
	case Remix:
		fmt.Printf("  %s\n", cyan("npm run dev"))
		fmt.Printf("  %s\n", cyan("npm run build"))
		fmt.Printf("  %s\n", cyan("npm start"))
	case NestJS:
		fmt.Printf("  %s\n", cyan("npm run start:dev"))
		fmt.Printf("  %s\n", cyan("npm run build"))
		fmt.Printf("  %s\n", cyan("npm run start:prod"))
	}
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

// promptNodeProjectType prompts the user to select a Node.js project type
func promptNodeProjectType() NodeProjectType {
	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	white := color.New(color.FgWhite, color.Bold).SprintFunc()

	options := GetNodeProjectOptions()

	fmt.Printf("\n%s Select Node.js project type:\n\n", white("📋"))

	// Print options with descriptions
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
