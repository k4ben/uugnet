package main

import (
	"uugnet/internal/art"

	"github.com/charmbracelet/lipgloss"
)

func generateBanner() string {

	logoStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("99"))
	titleStyle := lipgloss.NewStyle().Align(lipgloss.Left)
	subtitleStlye := lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	bannerStyle := lipgloss.NewStyle().BorderStyle(lipgloss.DoubleBorder()).BorderForeground(lipgloss.Color("8")).Padding(0, 2)

	logo := logoStyle.Render(art.Tux)

	titleWithSubtitle := lipgloss.JoinVertical(lipgloss.Center, titleStyle.Render(art.Title), subtitleStlye.Render("A telnet server for UUGers."))

	msg := lipgloss.JoinHorizontal(lipgloss.Center, logo, "  ", titleWithSubtitle)
	return bannerStyle.Render(msg)
}
