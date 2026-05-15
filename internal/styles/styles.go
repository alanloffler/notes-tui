package styles

import "charm.land/lipgloss/v2"

var (
	primaryColor           = lipgloss.Color("#f59e0b")
	primaryForegroundColor = lipgloss.Color("#000000")
	foregroundColor        = lipgloss.Color("#ffffff")
	accentColor            = primaryColor
	mutedColor             = lipgloss.Color("#737373")
	dangerColor            = lipgloss.Color("#ef4444")
)

var (
	AppNameStyle         = lipgloss.NewStyle().Background(primaryColor).Foreground(primaryForegroundColor).Padding(0, 1).Bold(true)
	CursorStyle          = lipgloss.NewStyle().Foreground(accentColor)
	TitleStyle           = lipgloss.NewStyle().Foreground(foregroundColor).Bold(true)
	TextStyle            = lipgloss.NewStyle().Foreground(foregroundColor)
	MutedStyle           = lipgloss.NewStyle().Foreground(mutedColor).Faint(true)
	MutedForegroundStyle = lipgloss.NewStyle().Foreground(mutedColor)
	DangerStyle          = lipgloss.NewStyle().Foreground(dangerColor)
)
