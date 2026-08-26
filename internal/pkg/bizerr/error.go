package bizerr

type Error struct {
	reason string
}

func (e *Error) Error() string  { return e.reason }
func (e *Error) Reason() string { return e.reason }

func New(reason string) *Error {
	return &Error{
		reason: reason,
	}
}
