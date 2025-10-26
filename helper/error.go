package helper

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

type NotFoundError struct {
	Err string
}

func (e NotFoundError) Error() string {
	return e.Err
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Err: error}
}

type ResourceConflictError struct {
	Err string
}

func (e ResourceConflictError) Error() string {
	return e.Err
}

func NewResourceConflictError(error string) ResourceConflictError {
	return ResourceConflictError{Err: error}
}

type BadRequestError struct {
	Err string
}

func (e BadRequestError) Error() string {
	return e.Err
}

func NewBadRequestError(error string) BadRequestError {
	return BadRequestError{Err: error}
}

type UnauthorizedError struct {
	Err string
}

func (e UnauthorizedError) Error() string {
	return e.Err
}

func NewUnauthorizedError(error string) UnauthorizedError {
	return UnauthorizedError{Err: error}
}
