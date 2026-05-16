package app

import (
	"log"

	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/alanloffler/notes-tui/internal/note"
	"github.com/alanloffler/notes-tui/internal/styles"
)

const (
	listView uint = iota
	titleView
	bodyView
)

type Model struct {
	state      uint
	store      note.Store
	notes      []note.Note
	currNote   note.Note
	listIndex  int
	textarea   textarea.Model
	textinput  textinput.Model
	deleteNote *note.Note
}

func New(store note.Store) Model {
	notes, err := store.GetNotes()
	if err != nil {
		log.Fatalf("unable to get notes: %v", err)
	}

	model := Model{
		state:     listView,
		store:     store,
		notes:     notes,
		textarea:  textarea.New(),
		textinput: textinput.New(),
	}

	model.textarea.SetWidth(50)
	model.textarea.Placeholder = "Write the note content"

	// tmp fix to set color on border
	model.textarea.SetPromptFunc(2, func(_ textarea.PromptInfo) string {
		return "│ "
	})

	st := model.textarea.Styles()
	st.Focused.Prompt = styles.PromptStyle.Faint(true)
	st.Focused.LineNumber = styles.LineNumberStyle.Faint(true)
	st.Focused.CursorLineNumber = styles.CursorLineNumberStyle
	st.Blurred.Prompt = styles.PromptStyle.Faint(true)
	st.Blurred.LineNumber = styles.LineNumberStyle.Faint(true)
	st.Blurred.CursorLineNumber = styles.CursorLineNumberStyle

	model.textarea.SetStyles(st)

	return model
}

func (m Model) Init() tea.Cmd {
	return nil
}
