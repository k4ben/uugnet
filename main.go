package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var users = map[string]string{
	"u":   "p",
	"ben": "password",
}

func main() {
	port := ":8080"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
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

func handleConnection(conn net.Conn) {
	defer conn.Close()

	shouldClose := false

	reader := bufio.NewReader(conn)

	fmt.Fprintf(conn, "Enter username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading username:", err)
		return
	}
	username = strings.TrimSpace(username)

	fmt.Fprintf(conn, "Enter password: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading password:", err)
		return
	}
	password = strings.TrimSpace(password)

	storedPassword, ok := users[username]
	if !ok || storedPassword != strings.ToLower(password) {
		fmt.Fprintf(conn, "Incorrect username or password\n")
		return
	}
	fmt.Fprint(conn, generateBanner())

	for !shouldClose {
		fmt.Fprint(conn, generatePrompt(username))
		in, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		args := strings.Split(strings.TrimSpace(in), " ")
		switch strings.ToLower(args[0]) {
		case "help":
			fmt.Fprint(conn, help(args))
		case "exit":
			fmt.Fprintln(conn, "Leaving uugnet... Bye!")
			shouldClose = true
		default:
			fmt.Fprintf(conn, "Unknown command: %s\n", args[0])
		}
	}

}
