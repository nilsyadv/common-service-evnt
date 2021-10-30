package error

import (
	"errors"

	"gorm.io/gorm"
)

// NewDatabaseError creates a new database error
func NewDatabaseError(err error) IDatabaseError {
	return &DatabaseError{CreateUnexpectedError(ErrorCodeDatabaseFailure, err)}
}

// DatabaseError represents an database query failure error interface
type IDatabaseError interface {
	IUnexpected
	IsRecordNotFoundError() bool
}

type DatabaseError struct {
	UnexpectedError
}

func (e *DatabaseError) IsRecordNotFoundError() bool {
	return errors.Is(e.cause, gorm.ErrRecordNotFound)
}
