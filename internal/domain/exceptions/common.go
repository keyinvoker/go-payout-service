package exceptions

import "fmt"

type ErrNotFound struct {
	Resource string
	ID       int
}

func (e ErrNotFound) Error() string {
	msg := "not found"
	if e.Resource != "" {
		msg = e.Resource + " " + msg
	}
	if e.ID != 0 {
		msg += fmt.Sprintf("with ID: %d", e.ID)
	}
	return msg
}

type ErrDatabase struct {
	Op       string
	Err      error
	Resource string
}

func (e ErrDatabase) Error() string {
	msg := fmt.Sprintf("database error during %s", e.Op)
	if e.Resource != "" {
		msg += " on " + e.Resource
	}
	if e.Err != nil {
		msg += ": " + e.Err.Error()
	}
	return msg
}
