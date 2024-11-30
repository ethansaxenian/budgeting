package tui

import "github.com/charmbracelet/lipgloss"

var (
	emptyStyle        = lipgloss.NewStyle()
	modelStyle        = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).BorderStyle(lipgloss.HiddenBorder())
	focusedModelStyle = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("#555555"))
	tableHeaderStyle  = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).Bold(true).PaddingBottom(1)
	bold              = lipgloss.NewStyle().Bold(true)
)

const (
	red   = lipgloss.Color("#FF0000")
	green = lipgloss.Color("#00FF00")
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
