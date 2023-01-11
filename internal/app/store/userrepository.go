package store

import "github.com/Abylkaiyr/forum/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow("INSERT INTO users(email,username, password) VALUES ($1, $2, $3) RETURNING id",
		u.Email,
		u.Username,
		u.Password,
	).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}
