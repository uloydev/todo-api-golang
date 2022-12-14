package exception

type NotFoundError struct {
	Message string
}

func (NotFoundError NotFoundError) Error() string {
	return NotFoundError.Message
}
