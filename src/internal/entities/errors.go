package entities

import "errors"

var (
	ErrIdIsRequired            = errors.New("id is required")
	ErrIdIsInvalid             = errors.New("id is not valid")
	ErrNameIsRequired          = errors.New("name is required")
	ErrSurnameIsRequired       = errors.New("surname is required")
	ErrEmailIsRequired         = errors.New("email is required")
	ErrJobPositionIsRequired   = errors.New("job position is required")
	ErrPasswordIsRequired      = errors.New("password is required")
	ErrPasswordIsInvalid       = errors.New("error while trying to save password")
	ErrPasswordLengthIsInvalid = errors.New("password has to have at least 8 characteres")

	ErrStockIsRequired = errors.New("stock is required")
)
