package logger

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

func Fatal(err error) {
	if err != nil {
		fmt.Println(lipgloss.NewStyle().
			Foreground(lipgloss.Color("1")).
			Render(fmt.Sprintf("Error: %s", err.Error())))
		os.Exit(1)
	}
}
