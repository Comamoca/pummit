package cmd

import (
	"github.com/HidemaruOwO/pummit/src/config"
	//  "github.com/fatih/color"
	"fmt"
)

func HelpCmd() {
	fmt.Printf("%s\n", help())
}

func help() string {
	return fmt.Sprintf(`pummit - %s
Usage:
chuno [PATH] [Option]
PATH Preview file path
- Option
--help Show help
--version Show version`, config.Version)
}