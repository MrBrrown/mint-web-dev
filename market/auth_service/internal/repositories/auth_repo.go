package repo

import (
	"database/sql"
	"errors"
)

type AuthRepo interface {
	TryLogin(login string) (User, error)
}

type authRepoImpl struct {
	db *sql.DB
}

func New(database *sql.DB) AuthRepo {
	return &authRepoImpl{db: database}
}

func (r *authRepoImpl) TryLogin(login string) (User, error) {
	var user User

	query := `
		SELECT u.id, u.username, u.password_hash, r.name as role
		FROM users u
		JOIN roles r ON u.role_id = r.id
		WHERE u.username = $1
	`

	err := r.db.QueryRow(query, login).Scan(
		&user.ID,
		&user.Login,
		&user.PasswordHash,
		&user.Role,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, errors.New("user not found")
		}
		return User{}, err
	}

	return user, nil
}
