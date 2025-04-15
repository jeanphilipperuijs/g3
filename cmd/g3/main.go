package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run bootstrap.go <project-name>")
		return
	}

	projectName := os.Args[1]
	projectPath := filepath.Join(".", projectName)

	// Create the project directory
	err := os.MkdirAll(projectPath, 0755)
	if err != nil {
		fmt.Printf("Error creating project directory: %v\n", err)
		return
	}

	// Initialize the Go module
	cmd := exec.Command("go", "mod", "init", projectName)
	cmd.Dir = projectPath
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error initializing Go module: %v\n", err)
		return
	}

	// Create the directory structure
	dirs := []string{
		"cmd/" + projectName,
		"internal/handler",
		"internal/service",
		"internal/repository",
		"pkg/utils",
	}

	for _, dir := range dirs {
		err := os.MkdirAll(filepath.Join(projectPath, dir), 0755)
		if err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			return
		}
	}

	// Create the main.go file
	mainGoPath := filepath.Join(projectPath, "cmd", projectName, "main.go")
	err = os.WriteFile(mainGoPath, []byte(mainGoContent), 0644)
	if err != nil {
		fmt.Printf("Error creating main.go: %v\n", err)
		return
	}

	// Create the .gitignore file
	gitignorePath := filepath.Join(projectPath, ".gitignore")
	err = os.WriteFile(gitignorePath, []byte(gitignoreContent), 0644)
	if err != nil {
		fmt.Printf("Error creating .gitignore: %v\n", err)
		return
	}

	// Create the README.md file
	readmePath := filepath.Join(projectPath, "README.md")
	err = os.WriteFile(readmePath, []byte(readmeContent), 0644)
	if err != nil {
		fmt.Printf("Error creating README.md: %v\n", err)
		return
	}

	fmt.Printf("Project %s bootstrapped successfully!\n", projectName)
}

const mainGoContent = `package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}
`

const gitignoreContent = `# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Pytest stuff
.pytest_cache/

# Environments
.env
.venv
env/
venv/
ENV/
env.bak/
venv.bak/

# pycharm stuff
.idea/

# vscode stuff
.vscode/

# go stuff
*.swp
*.swo
`

const readmeContent = `# Project Name

This is a new Go project bootstrapped with a simple script.

## Directory Structure

myapp/
├── cmd/
│   └── myapp/
│       └── main.go
├── internal/
│   ├── handler/
│   ├── service/
│   └── repository/
├── pkg/
│   └── utils/
├── go.mod
├── go.sum
├── README.md
└── .gitignore

## Getting Started

1. **Initialize the Module**:
   
   go mod init github.com/yourusername/myapp

Create the Directory Structure:

mkdir -p cmd/myapp internal/{handler,service,repository} pkg/utils

Create the Main Entry Point:

    // cmd/myapp/main.go
    package main

    import (
        "fmt"
    )

    func main() {
        fmt.Println("Hello, World!")
    }



### How to Use the Script

1. Save the above code to a file named bootstrap.go.
2. Run the script with the desired project name:

	go run bootstrap.go myproject
`
