package model

// Error is for errors in the business domain. See the constants below.
type Error string

const (
	ErrorNoteNotFound = Error("NOTE_NOT_FOUND")
)

func (e Error) Error() string {
	return string(e)
}
