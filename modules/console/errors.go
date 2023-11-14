package console

type AbortError struct {
	NestedError error
}

func (e AbortError) Error() string {
	return e.NestedError.Error()
}

func NewAbortError(err error) AbortError {
	return AbortError{NestedError: err}
}
