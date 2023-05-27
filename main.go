package main

import (
	"github.com/ibilalkayy/proctl/cmd"
	_ "github.com/ibilalkayy/proctl/cmd/member"
	_ "github.com/ibilalkayy/proctl/cmd/profile"
	_ "github.com/ibilalkayy/proctl/cmd/user"
	_ "github.com/ibilalkayy/proctl/cmd/work"
	_ "github.com/ibilalkayy/proctl/cmd/workspace"
)

func main() {
	cmd.Execute()
}
