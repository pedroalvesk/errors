package errors

type Error struct {
	msg  string
	code int
}

func (err *Error) Error() string {
	return err.msg
}

func New(msg string, code int) error {
	return &Error{msg: msg, code: code}
}
