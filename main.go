package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"uugnet/internal/commands"
	"uugnet/internal/db"
	"uugnet/internal/logger"
	"uugnet/internal/user"

	"github.com/charmbracelet/lipgloss"
)

func handleArgs(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: uugnet <command>")
		fmt.Println("Commands:")
		fmt.Println("serve\t\t\tRun the uugnet server")
		fmt.Println("userlist\t\tList users")
		fmt.Println("useradd <username>\tAdd user")
		fmt.Println()
		os.Exit(0)
	}
	switch args[0] {
	case "userlist":
		user.CLI.UserList()
	case "useradd":
		user.CLI.AddUser(args)
	case "serve":
		return
	}
	os.Exit(0)
}

var enableForgotPassword = false

func main() {
	err := db.InitDatabase()
	commands.InitCommands()
	logger.Fatal(err)
	serveCmd := flag.NewFlagSet("serve", flag.ExitOnError)
	serveForgot := serveCmd.Bool("f", false, "uugnet serve -f")
	if serveForgot != nil && *serveForgot {
		enableForgotPassword = true
	}
	serveCmd.Parse(os.Args[2:])
	flag.Parse()
	args := flag.Args()
	handleArgs(args)
	port := ":23"
	listener, err := net.Listen("tcp", port)
	logger.Fatal(err)
	defer listener.Close()

	fmt.Printf("uugnet server started on %s\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func generatePrompt(name string) string {
	nameStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("5"))
	symbolStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	uugnetStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("99"))

	prompt := "\n" + nameStyle.Render(name) + symbolStyle.Render("@") + uugnetStyle.Render("uugnet") + symbolStyle.Render("> ")
	return prompt
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	fmt.Fprintf(conn, "\nuugnet login: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading username:", err)
		return
	}

	// This is a stupid thing that telnet on macOS sends when you don't specify a port, so i'm manually removing it for the time being
	prefix := []byte{255, 251, 37, 255, 253, 3, 255, 251, 24, 255, 251, 31, 255, 251, 32, 255, 251, 33, 255, 251, 34, 255, 251, 39, 255, 253, 5}

	username = strings.TrimSpace(username)
	username = strings.Replace(username, string(prefix), "", 1)

	fmt.Fprintf(conn, "Password: ")
	password, err := reader.ReadString('\n')
	fmt.Fprintln(conn)
	if err != nil {
		fmt.Println("Error reading password:", err)
		return
	}
	password = strings.TrimSpace(password)

	userRow, err := db.GetUser(username)

	if err != nil {
		fmt.Fprintf(conn, "User not found: '%s'\n", username)
		fmt.Println(err)
		return
	} else if userRow.Password != password {
		fmt.Fprintln(conn, "Incorrect username or password")
		if enableForgotPassword {
			fmt.Fprintf(conn, "\nForgot password? [Y/n]: ")
			reader.ReadString('\n')
			fmt.Fprintf(conn, "Your password is '%s'\n\n", userRow.Password)
		}
		return
	}

	fmt.Fprintf(conn, "%s\n\n", generateBanner())
	fmt.Fprintf(conn, "Welcome to uugnet, %s! Type 'help' for commands.\n", username)

	for true {
		fmt.Fprint(conn, generatePrompt(username))
		in, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		args := strings.Split(strings.TrimSpace(in), " ")
		commands.HandleCommands(userRow, args, conn, *reader)
	}

}
