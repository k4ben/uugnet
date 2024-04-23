package main

import (
	"uugnet/art"

	"github.com/charmbracelet/lipgloss"
)

func generateBanner() string {

	logoStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("99"))
	titleStyle := lipgloss.NewStyle().Align(lipgloss.Left)
	subtitleStlye := lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	bannerStyle := lipgloss.NewStyle().BorderStyle(lipgloss.DoubleBorder()).BorderForeground(lipgloss.Color("8")).Padding(0, 4)

	logo := logoStyle.Render(art.Tux)

	titleWithSubtitle := lipgloss.JoinVertical(lipgloss.Center, titleStyle.Render(art.Title), subtitleStlye.Render("A modern BBS solution for UUGers."))

	msg := lipgloss.JoinHorizontal(lipgloss.Center, logo, "    ", titleWithSubtitle)
	return bannerStyle.Render(msg)
}

func generatePrompt(name string) string {

	nameStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("5"))
	symbolStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	uugnetStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("99"))

	prompt := "\n" + nameStyle.Render(name) + symbolStyle.Render("@") + uugnetStyle.Render("uugnet") + symbolStyle.Render("> ")
	return prompt
}
