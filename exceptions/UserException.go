package exceptions

type errorMsg struct {
	s string
}

func (e errorMsg) Error() string {
	return e.s
}

//UserServiceException --
func UserServiceException(text string) error {
	return errorMsg{text}
}
