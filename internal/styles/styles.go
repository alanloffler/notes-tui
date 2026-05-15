package styles

import "charm.land/lipgloss/v2"

var (
	appColor = lipgloss.Color("99")
)

var (
	AppNameStyle    = lipgloss.NewStyle().Background(appColor).Padding(0, 1)
	FaintStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Faint(true)
	EnumerableStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(1)
)
