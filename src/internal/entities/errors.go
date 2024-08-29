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

	ErrPriceIsInvalid         = errors.New("Pizza price cannot be 0 or less")
	ErrSizeIsRequired         = errors.New("Pizza size is required")
	ErrSizeIsInvalid          = errors.New("Pizza size has to be SMALL, MEDIUM or BIG")
	ErrSauceIsRequired        = errors.New("Pizza sauce is required")
	ErrIngredientsAreRequired = errors.New("Pizza ingredients are required")
	ErrIngredientIdIsInvalid  = errors.New("Pizza ingredient ID is invalid")
	ErrSauceIdIsInvalid       = errors.New("Pizza sauce ID is invalid")
)
