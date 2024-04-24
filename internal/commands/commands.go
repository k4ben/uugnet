package commands

import (
	"bufio"
	"fmt"
	"net"
	"uugnet/internal/db"
)

type Command struct {
	Name        string
	Usage       string
	Description string
	Handler     func(user db.UserDbRow, args []string, conn net.Conn, reader bufio.Reader)
}

var Commands = []Command{}

func InitCommands() {
	Commands = []Command{HelpCommand, ExitCommand, ChatCommand}
}

func HandleCommands(user db.UserDbRow, args []string, conn net.Conn, reader bufio.Reader) {
	for _, c := range Commands {
		if args[0] == c.Name {
			c.Handler(user, args, conn, reader)
			return
		}
	}
	fmt.Fprintf(conn, "Unknown command: %s\n", args[0])
}
