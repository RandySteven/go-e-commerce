package apperror

type ErrAgeValidation struct {
	Err error
}

type ErrEmailAndPasswordNotMatch struct {
	Err error
}

type ErrBadRequest struct {
	Err error
}

func (e *ErrAgeValidation) Error() string {
	return "the minimum age for user is 18"
}

func (e *ErrEmailAndPasswordNotMatch) Error() string {
	return "the email and password didnt match"
}

func (e *ErrBadRequest) Error() string {
	return e.Err.Error()
}
