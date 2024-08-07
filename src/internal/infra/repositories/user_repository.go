package repositories

import (
	"context"
	"database/sql"

	"github.com/CaiqueRibeiro/pizza-rute/src/internal/entities"
)

type UserRepositoryInterface interface {
	Create(user *entities.User) error
	List() ([]*entities.User, error)
	FindByID(id string) (*entities.User, error)
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

func (ur *UserRepository) List() ([]*entities.User, error) {
	rows, err := ur.db.Query("SELECT id, name, surname, email, photo_url, job_position from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*entities.User
	for rows.Next() {
		var u entities.User
		err = rows.Scan(&u.ID, &u.Name, &u.Surname, &u.Email, &u.PhotoUrl, &u.JobPosition)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}

func (ur *UserRepository) FindByID(id string) (*entities.User, error) {
	stmt, err := ur.db.Prepare("SELECT * FROM users WHERE id=?")
	if err != nil {
		return nil, err
	}
	var u entities.User
	err = stmt.QueryRowContext(context.Background(), id).Scan(&u.ID, &u.Name, &u.Surname, &u.Email, &u.PhotoUrl, &u.JobPosition, &u.Password)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
