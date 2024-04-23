package main

import (
	"io"
	"log"
	"os"

	"github.com/charmbracelet/lipgloss"
)

func readFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func generateBanner() string {
	a := readFile("tux.txt")
	b := readFile("title.txt")

	logoStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("99"))
	subtitleStlye := lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	bannerStyle := lipgloss.NewStyle().BorderStyle(lipgloss.DoubleBorder()).BorderForeground(lipgloss.Color("8")).Padding(0, 4)

	logo := logoStyle.Render(a)

	titleWithSubtitle := lipgloss.JoinVertical(lipgloss.Center, b, subtitleStlye.Render("A modern BBS solution for UUGers."))

	msg := lipgloss.JoinHorizontal(lipgloss.Center, logo, "    ", titleWithSubtitle)
	return bannerStyle.Render(msg)
}

func generatePrompt(name string) string {

	nameStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("5"))
	symbolStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	uugnetStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("99"))

	// Combine the styles with the respective text
	prompt := "\n" + nameStyle.Render(name) + symbolStyle.Render("@") + uugnetStyle.Render("uugnet") + symbolStyle.Render("> ")
	return prompt
}
