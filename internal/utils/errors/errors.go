package e

import "fmt"

type (
	ErrForbidden struct {
		err error
	}
	ErrNotFound struct {
		err error
	}
	ErrUnprocessableEntity struct {
		err error
	}
	ErrInvalidData struct {
		err error
	}
	ErrConflict struct {
		err error
	}
)

func (e *ErrForbidden) Error() string {
	return fmt.Sprintf("forbidden: %v", e.err)
}

func (e *ErrForbidden) Unwrap() error {
	return e.err
}

func NewErrForbidden(err error) *ErrForbidden {
	return &ErrForbidden{err}
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("not found: %v", e.err)
}

func (e *ErrNotFound) Unwrap() error {
	return e.err
}

func NewErrNotFound(err error) *ErrNotFound {
	return &ErrNotFound{err}
}

func (e *ErrUnprocessableEntity) Error() string {
	return fmt.Sprintf("unprocessable entity: %v", e.err)
}

func (e *ErrUnprocessableEntity) Unwrap() error {
	return e.err
}

func NewErrUnprocessableEntity(err error) *ErrUnprocessableEntity {
	return &ErrUnprocessableEntity{err}
}

func (e *ErrInvalidData) Error() string {
	return fmt.Sprintf("invalid value: %v", e.err)
}

func (e *ErrInvalidData) Unwrap() error {
	return e.err
}

func NewErrConflict(err error) *ErrConflict {
	return &ErrConflict{err}
}

func (e *ErrConflict) Error() string {
	return fmt.Sprintf("conflict: %v", e.err)
}

func (e *ErrConflict) Unwrap() error {
	return e.err
}

func NewErrInvalidData(err error) *ErrInvalidData {
	return &ErrInvalidData{err}
}
