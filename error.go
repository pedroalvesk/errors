package GoErrors

type Error struct {
	msg  string
	Code int
}

func (err *Error) Error() string {
	return err.msg
}

func New(msg string, code int) error {
	return &Error{msg: msg, Code: code}
}

func getCode(err error) int {
	cErr, ok := err.(*Error)
	if !ok || err == nil {
		return -1
	}

	return cErr.Code
}
