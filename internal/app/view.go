package app

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/alanloffler/notes-tui/internal/styles"
)

func (m Model) View() tea.View {
	s := styles.AppNameStyle.Render("Notes App") + "\n\n"

	if m.state == titleView {
		s += styles.TitleStyle.Render("Note title:")
		s += "\n\n"
		s += m.textinput.View()
		s += "\n\n\n"
		s += styles.MutedForegroundStyle.Render("enter ") + styles.MutedStyle.Render("save • ")
		s += styles.MutedForegroundStyle.Render("esc ") + styles.MutedStyle.Render("discard")
	}

	if m.state == bodyView {
		s += styles.TitleStyle.Render("Note:")
		s += "\n\n"
		s += m.textarea.View()
		s += "\n\n\n"
		s += styles.MutedForegroundStyle.Render("ctrl+s ") + styles.MutedStyle.Render("save • ")
		s += styles.MutedForegroundStyle.Render("esc ") + styles.MutedStyle.Render("discard")
	}

	if m.state == listView {
		for i, n := range m.notes {
			prefix := " "

			if i == m.listIndex {
				prefix = styles.CursorStyle.Render("▸")
			}

			shortBody := strings.ReplaceAll(n.Body, "\n", " ")
			if len(shortBody) > 30 {
				shortBody = shortBody[:30] + "..."
			}

			if i == m.listIndex && m.deleteNote != nil {
				s += prefix + " " + styles.DangerStyle.Render(n.Title) + " " + styles.MutedStyle.Render(shortBody)
			} else {
				s += prefix + " " + styles.TextStyle.Render(n.Title) + " " + styles.MutedStyle.Render(shortBody)
			}

			s += "\n"
		}

		if m.deleteNote != nil {
			s += "\n\n"
			s += styles.DangerStyle.Render("Delete this note?") + "  " + styles.TextStyle.Render(m.deleteNote.Title)
			s += "\n"
			s += styles.MutedForegroundStyle.Render("[Y/n]")
			s += "\n"
		}

		s += "\n\n"
		s += styles.MutedForegroundStyle.Render("n ") + styles.MutedStyle.Render("new • ")
		s += styles.MutedForegroundStyle.Render("enter ") + styles.MutedStyle.Render("edit • ")
		s += styles.MutedForegroundStyle.Render("space ") + styles.MutedStyle.Render("delete • ")
		s += styles.MutedForegroundStyle.Render("q ") + styles.MutedStyle.Render("quit")
	}

	v := tea.NewView(s)
	v.AltScreen = true
	return v
}
