package commands

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"uugnet/internal/db"
)

type ChatMessage struct {
	Username string
	Message  string
}

var messages = []ChatMessage{}

var ChatCommand = Command{
	Name:        "chat",
	Usage:       "chat <read | write <message>>",
	Description: "A not-so-good chat client.",
	Handler: func(user db.UserDbRow, args []string, conn net.Conn, reader bufio.Reader) {
		result := ""
		if len(args) < 2 {
			result += "Usage: chat <read | write <message>>\n"
		} else {
			if args[1] == "read" {
				for _, m := range messages {
					result += fmt.Sprintf("%s: %s\n", m.Username, m.Message)
				}
			} else if args[1] == "write" {
				text := strings.Join(args[2:], " ")
				messages = append(messages, ChatMessage{
					Username: user.Username,
					Message:  text,
				})
			}
		}
		fmt.Fprint(conn, result)
	},
}
