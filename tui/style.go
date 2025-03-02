package tui

import "github.com/charmbracelet/lipgloss"

const (
	red   = lipgloss.Color("#FF0000")
	green = lipgloss.Color("#00FF00")
	gray  = lipgloss.Color("#555555")
)

var (
	modelStyle        = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).BorderStyle(lipgloss.HiddenBorder())
	focusedModelStyle = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).BorderStyle(lipgloss.NormalBorder()).BorderForeground(gray)
	tableHeaderStyle  = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).Bold(true).PaddingBottom(1)
	bold              = lipgloss.NewStyle().Bold(true)
)

func colorAmount(amount float64) lipgloss.Style {
	style := lipgloss.NewStyle().UnsetPadding().UnsetMargins()

	if amount > 0 {
		return style.Foreground(green)
	} else if amount < 0 {
		return style.Foreground(red)
	} else {
		return style
	}
}
