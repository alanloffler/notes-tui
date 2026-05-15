package app

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/alanloffler/notes-tui/internal/styles"
)

func (m Model) View() tea.View {
	s := styles.AppNameStyle.Render("Notes App") + "\n\n"

	if m.state == titleView {
		s += "Note title:\n\n"
		s += m.textinput.View() + "\n\n"
		s += styles.FaintStyle.Render(" enter save - esc discard")
	}

	if m.state == bodyView {
		s += "Note:\n\n"
		s += m.textarea.View() + "\n\n"
		s += styles.FaintStyle.Render("ctrl+s save - esc discard")
	}

	if m.state == listView {
		for i, n := range m.notes {
			prefix := " "
			if i == m.listIndex {
				prefix = "▸"
			}

			shortBody := strings.ReplaceAll(n.Body, "\n", " ")
			if len(shortBody) > 30 {
				shortBody = shortBody[:30]
			}
			s += styles.EnumerableStyle.Render(prefix) + n.Title + " | " + styles.FaintStyle.Render(shortBody)
			s += "\n\n"
		}

		s += styles.FaintStyle.Render("n new note - q quit")
	}

	return tea.NewView(s)
}
