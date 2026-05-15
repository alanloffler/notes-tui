package app

import (
	"log"

	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/alanloffler/notes-tui/internal/note"
)

const (
	listView uint = iota
	titleView
	bodyView
)

type Model struct {
	state     uint
	store     note.Store
	notes     []note.Note
	currNote  note.Note
	listIndex int
	textarea  textarea.Model
	textinput textinput.Model
}

func New(store note.Store) Model {
	notes, err := store.GetNotes()
	if err != nil {
		log.Fatalf("unable to get notes: %v", err)
	}

	return Model{
		state:     listView,
		store:     store,
		notes:     notes,
		textarea:  textarea.New(),
		textinput: textinput.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
