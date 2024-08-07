package repositories

import (
	"database/sql"

	"github.com/CaiqueRibeiro/pizza-rute/src/internal/entities"
)

type UserRepositoryInterface interface {
	Create(user *entities.User) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(user *entities.User) error {
	stmt, err :=
		ur.db.Prepare("INSERT INTO users(id, name, surname, email, photo_url, job_position, password)values (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		user.ID.String(),
		user.Name,
		user.Surname,
		user.Email,
		user.PhotoUrl,
		user.JobPosition,
		user.Password,
	)
	if err != nil {
		return err
	}
	return nil
}
