package commands

import (
	"bufio"
	"fmt"
	"net"
	"uugnet/internal/db"
)

var ExitCommand = Command{
	Name:        "exit",
	Usage:       "exit",
	Description: "Disconnect from uugnet.",
	Handler: func(user db.UserDbRow, args []string, conn net.Conn, reader bufio.Reader) {
		fmt.Fprint(conn, "Leaving uugnet... Bye!\n\n")
		conn.Close()
	},
}
