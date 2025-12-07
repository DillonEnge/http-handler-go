package handle

type HandlerError struct {
	Status int
	Err    error
}

func (a *HandlerError) Error() string {
	if a.Err == nil {
		return "error not provided"
	}

	return a.Err.Error()
}
