package main

import (
	"errors"
	"fmt"
)

type CommandInfo struct {
	Name        string
	Usage       string
	Description string
}

var commands = []CommandInfo{
	{"help", "help [command]", "Lists commands and their usage"},
}

func findCommand(name string) (*CommandInfo, error) {
	for _, c := range commands {
		if c.Name == name {
			return &c, nil
		}
	}
	return nil, errors.New("command not found")
}

func help(args []string) string {
	result := ""
	if len(args) < 2 {
		result += "Type 'help [command]' for more info about a command.\n"
		result += "Available commands:\n"
		for _, c := range commands {
			result += fmt.Sprintf("\t%s\t\t%s\n", c.Usage, c.Description)
		}
	} else {
		c, err := findCommand(args[0])
		if err != nil {
			result += fmt.Sprintf("help: command not found: %s\n", args[0])
		}
		result += fmt.Sprintf("%s: %s\t\t%s\n", c.Name, c.Usage, c.Description)
	}
	return result
}
