package apperror

type ErrAgeValidation struct {
	Err error
}

func (e *ErrAgeValidation) Error() string {
	return "the minimum age for user is 18"
}
