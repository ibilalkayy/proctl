package main

import (
	"github.com/ibilalkayy/proctl/cmd"
	_ "github.com/ibilalkayy/proctl/cmd/profile"
	_ "github.com/ibilalkayy/proctl/cmd/user"
)

func main() {
	cmd.Execute()
}
