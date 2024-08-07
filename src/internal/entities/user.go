package entities

import (
	"errors"

	"github.com/CaiqueRibeiro/pizza-rute/src/internal/dtos"
	"github.com/CaiqueRibeiro/pizza-rute/src/pkg/entities"
	"golang.org/x/crypto/bcrypt"
)

var (
	minPassLength              = 8
	bCryptCost                 = 12
	ErrIdIsRequired            = errors.New("id is required")
	ErrIdIsInvalid             = errors.New("id is not valid")
	ErrNameIsRequired          = errors.New("name is required")
	ErrSurnameIsRequired       = errors.New("surname is required")
	ErrEmailIsRequired         = errors.New("email is required")
	ErrJobPositionIsRequired   = errors.New("job position is required")
	ErrPasswordIsRequired      = errors.New("password is required")
	ErrPasswordIsInvalid       = errors.New("error while trying to save password")
	ErrPasswordLengthIsInvalid = errors.New("password has to have at least 8 characteres")
)

type User struct {
	ID          entities.ID `json:"id"`
	Name        string      `json:"name"`
	Surname     string      `json:"surname"`
	Email       string      `json:"email"`
	PhotoUrl    string      `json:"photo_url"`
	JobPosition string      `json:"job_position"`
	Password    string      `json:"-"`
}

func (u *User) Validate() []error {
	errs := []error{}

	if u.ID.String() == "" {
		errs = append(errs, ErrIdIsRequired)
	}
	if _, err := entities.ParseID(u.ID.String()); err != nil {
		errs = append(errs, ErrIdIsInvalid)
	}
	if u.Name == "" {
		errs = append(errs, ErrNameIsRequired)
	}
	if u.Surname == "" {
		errs = append(errs, ErrSurnameIsRequired)
	}
	if u.Email == "" {
		errs = append(errs, ErrEmailIsRequired)
	}
	if u.JobPosition == "" {
		errs = append(errs, ErrJobPositionIsRequired)
	}
	if u.Password == "" {
		errs = append(errs, ErrPasswordIsRequired)
	}
	return errs
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func NewUser(newUser dtos.CreateUserInput) (*User, []error) {
	if len(newUser.TempPassword) < minPassLength {
		return nil, []error{ErrPasswordLengthIsInvalid}
	}

	encpw, err := bcrypt.GenerateFromPassword([]byte(newUser.TempPassword), bCryptCost)

	if err != nil {
		return nil, []error{ErrPasswordIsInvalid}
	}

	user := &User{
		ID:          entities.NewID(),
		Name:        newUser.Name,
		Surname:     newUser.Surname,
		Email:       newUser.Email,
		PhotoUrl:    newUser.PhotoUrl,
		JobPosition: newUser.JobPosition,
		Password:    string(encpw),
	}

	errs := user.Validate()

	if len(errs) > 0 {
		return nil, errs
	}

	return user, []error{}
}
