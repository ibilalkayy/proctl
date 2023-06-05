// Package main is the entry point of the program.
package main

// Import necessary packages and libraries.
import (
	// This imports the cmd package from the github.com/ibilalkayy/proctl repository.
	"github.com/ibilalkayy/proctl/cmd"

	// The following import statements import specific commands from the proctl tool.
	// The "_" before the package means that the package is being imported for its side-effects only
	// i.e., it's only being imported to execute an initialization function.
	_ "github.com/ibilalkayy/proctl/cmd/board"
	_ "github.com/ibilalkayy/proctl/cmd/member"
	_ "github.com/ibilalkayy/proctl/cmd/profile"
	_ "github.com/ibilalkayy/proctl/cmd/user"
	_ "github.com/ibilalkayy/proctl/cmd/work"
	_ "github.com/ibilalkayy/proctl/cmd/workspace"
)

// Function main is the entry point of the program.
func main() {
	// Call the Execute function from the cmd package.
	// This will start the execution of the program.
	cmd.Execute()
}
