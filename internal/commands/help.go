package commands

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"uugnet/internal/db"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var HelpCommand = Command{
	Name:        "help",
	Usage:       "help [command]",
	Description: "Lists commands and their usage",
	Handler: func(user db.UserDbRow, args []string, conn net.Conn, reader bufio.Reader) {
		result := ""
		if len(args) < 2 {
			result += "Type 'help [command]' for more info about a command.\n"
			result += "Available commands:\n"
			table := helpTable(Commands)
			result += table.Render()
		} else {
			c, err := findCommand(args[1])
			if err != nil {
				result += fmt.Sprintf("help: command not found: %s\n", args[1])
			}
			result += fmt.Sprintf("%s: %s\t\t%s\n", c.Name, c.Usage, c.Description)
		}
		fmt.Fprint(conn, result)
	},
}

func helpTable(commands []Command) table.Table {
	t := table.New().Border(lipgloss.HiddenBorder()).Width(80)
	for _, c := range commands {
		t.Row(c.Usage, c.Description)
	}
	return *t
}

func findCommand(name string) (*Command, error) {
	for _, c := range Commands {
		if c.Name == name {
			return &c, nil
		}
	}
	return nil, errors.New("command not found")
}
