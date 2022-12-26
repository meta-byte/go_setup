package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: must specify project name.")
		os.Exit(1)
	}
	projectName := os.Args[1]

	err := os.Chdir("/Users/garretthoward/Developer/Go")
	if err != nil {
		fmt.Printf("Error changing to Go workspace directory %s\n", err)
		os.Exit(1)
	}

	err = os.Mkdir(projectName, 0755)
	if err != nil {
		fmt.Printf("Error creating project directory: %s\n", err)
		os.Exit(1)
	}

	err = os.Chdir(projectName)
	if err != nil {
		fmt.Printf("Error changing to project directory: %s\n", err)
		os.Exit(1)
	}

	pathString := "github.com/meta-byte/" + projectName
	cmd := exec.Command("go", "mod", "init", pathString)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error initializing go package: \n%s \n%s", cmd, err)
		os.Exit(1)
	}

	cmd = exec.Command("git", "init")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error initializing Git repository: %s\n", err)
		os.Exit(1)
	}

	mainFile, err := os.Create("main.go")
	if err != nil {
		fmt.Printf("Error creating main package: %s\n", err)
		os.Exit(1)
	}

	defer mainFile.Close()
	mainFile.WriteString("package main\n\nfunc main(){\n}")

	projectPath, _ := os.Getwd()
	fmt.Printf("Successfully created new Go project %s\n", projectPath)
}
