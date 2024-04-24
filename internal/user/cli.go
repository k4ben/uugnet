package user

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"uugnet/internal/db"
	"uugnet/internal/logger"

	"golang.org/x/term"
)

const UsernameRegex = "^[a-z0-9_-]{3,12}$"

const PasswordRegex = `^[a-zA-Z0-9!@#$%^&*()_+{}|:"<>?\-=[\]\\;',./~]{3,12}$`

type commandNamespace struct{}

var CLI commandNamespace

func (commandNamespace) UserList() {
	users, err := db.GetUsers()
	logger.Fatal(err)
	if len(users) == 0 {
		fmt.Print("No users found.")
	}
	for _, u := range users {
		fmt.Printf("%s\t", u.Username)
	}
	fmt.Print("\n\n")
}

func (commandNamespace) AddUser(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: uugnet useradd <username>")
	} else {
		username := strings.TrimSpace(args[1])
		validUsername, err := regexp.Match(UsernameRegex, []byte(username))
		if err != nil {
			panic(err)
		} else if !validUsername {
			fmt.Println("Error: Invalid username")
			os.Exit(0)
		}
		fmt.Printf("New password for '%s': ", username)
		oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		var line string
		fmt.Scan(&line)
		term.Restore(int(os.Stdin.Fd()), oldState)
		fmt.Println()
		logger.Fatal(err)
		password := strings.TrimSpace(line)
		validPassword, err := regexp.Match(PasswordRegex, []byte(line))
		logger.Fatal(err)
		if !validPassword {
			fmt.Println("Error: Invalid password")
		} else {
			err = db.AddUser(&db.User{
				Username: username,
				Password: password,
			})
			if err != nil {
				fmt.Println("Error: Couldn't add user")
			} else {
				fmt.Printf("Added user '%s'", username)
			}
		}
	}
	fmt.Println()
}
