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
	// Add more users and passwords as needed
}

func main() {
	println(generateBanner())
	// Start a TCP server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("uugnet server started, listening on port 8080")

	for {
		// Accept connections from clients
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle each connection in a separate goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	shouldClose := false

	// Create a bufio reader to read data from the connection
	reader := bufio.NewReader(conn)

	// Read username
	fmt.Fprintf(conn, "Enter username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading username:", err)
		return
	}
	username = strings.TrimSpace(username)

	// Read password
	fmt.Fprintf(conn, "Enter password: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading password:", err)
		return
	}
	password = strings.TrimSpace(password)

	// Check if the username and password are correct
	// storedPassword, ok := users[strings.ToLower(username)]
	// if !ok || strings.ToLower(storedPassword) != strings.ToLower(password) {
	// 	fmt.Fprintf(conn, "Incorrect username or password\n")
	// 	return
	// }
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
		default:
			fmt.Fprintf(conn, "Unknown command: %s\n", args[0])
		}
	}

	// Now you can implement your event handling logic here
}
