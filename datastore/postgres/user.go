package postgres

import (
	"database/sql"

	"github.com/anthonynsimon/parrot/errors"
	"github.com/anthonynsimon/parrot/model"
)

func (db *PostgresDB) GetUser(id int) (*model.User, error) {
	u := model.User{}
	row := db.QueryRow("SELECT * FROM users WHERE id = $1", id)

	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return &u, nil
}

func (db *PostgresDB) GetUserByEmail(email string) (*model.User, error) {
	u := model.User{}
	row := db.QueryRow("SELECT * FROM users WHERE email = $1", email)

	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return &u, nil
}

func (db *PostgresDB) CreateUser(u *model.User) error {
	row := db.QueryRow("INSERT INTO users (email, password, role) VALUES($1, $2, $3) RETURNING id", u.Email, u.Password, u.Role)
	return row.Scan(&u.ID)
}

func (db *PostgresDB) UpdateUser(u *model.User) error {
	return errors.ErrInternal
}

func (db *PostgresDB) DeleteUser(id int) (int, error) {
	err := db.QueryRow("DELETE FROM users WHERE id = $1 RETURNING id", id).Scan(&id)
	if err == sql.ErrNoRows {
		return -1, errors.ErrNotFound
	}
	return id, err
}