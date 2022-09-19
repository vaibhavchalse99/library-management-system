package db

import (
	"context"
	"database/sql"
	"time"
)

const (
	CreateUser = `INSERT INTO users(name, email, password, role, created_at, updated_at) VALUES($1,$2,$3,$4,$5,$6)`
)

type User struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (d *userStore) CreateUser(ctx context.Context, user *User) (err error) {
	now := time.Now()
	return Transact(ctx, d.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = d.db.Exec(
			CreateUser,
			user.Name,
			user.Email,
			user.Password,
			user.Role,
			now,
			now,
		)
		return err
	})
}

func (d *userStore) GetUsers(ctx context.Context) {

}

func (d *userStore) UpdateUser(ctx context.Context) {

}
