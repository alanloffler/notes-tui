package app

import (
	tea "charm.land/bubbletea/v2"
	"github.com/alanloffler/notes-tui/internal/note"
)

const (
	listView uint = iota
	titleView
	bodyView
)

type Model struct {
	state uint
	store note.Store
}

func New(store note.Store) Model {
	return Model{
		state: listView,
		store: store,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
