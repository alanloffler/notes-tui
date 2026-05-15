package note

type Store interface {
	Init() error
	GetNotes() ([]Note, error)
	SaveNote(note Note) error
	DeleteNote(id int64) error
}
