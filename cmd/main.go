package main

import (
	"log"

	tea "charm.land/bubbletea/v2"
	"github.com/alanloffler/notes-tui/internal/app"
	"github.com/alanloffler/notes-tui/internal/note"
)

func main() {
	store := &note.SQLiteStore{}
	if err := store.Init(); err != nil {
		log.Fatalf("unable to init store: %v", err)
	}

	m := app.New(store)

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Fatalf("unable to run app: %v", err)
	}
}
