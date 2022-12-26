## Overview
This is a simple Go program for creating new Go projects. The program will create a new directory with the specified project name and initialize a new Git repository inside it. It will also create a main package with a main function in a main.go file inside the project directory. It then also initializes the go module. You can then add your Go code to the main.go file and use Git to manage your project as needed.

## Usage

1. Download or clone go_setup repository onto your local machine.
2. Edit the path for the working directory where you would like your go projects to be created.
3. Edit the path string for initializing the go module to be your preferred format. 
3. Navigate into the go_setup directory and run the make create command followed by the name of the project you want to create. `make create project=<projectName>`

- Optionally you can run `go run <filepath to go_setup>/main.go <projectName` from anywhere

This will create a new Go project with the specified project name in the parent directory specified in step two above.

