package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const (
	CreateUserQuery         = `INSERT INTO users(name, email, password, role, created_at, updated_at) VALUES($1,$2,$3,$4,$5,$6)`
	getUserListQuery        = `SELECT * FROM users`
	getUserDetailsQuery     = `SELECT * FROM users WHERE email=$1 AND password=$2`
	getUserDetailsByIdQuery = `SELECT * FROM users WHERE id=$1`
)

type User struct {
	ID        uuid.UUID `db:"id"`
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
			CreateUserQuery,
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

func (d *userStore) GetUsers(ctx context.Context) (users []User, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return d.db.SelectContext(ctx, &users, getUserListQuery)
	})
	if err == sql.ErrNoRows {
		return users, ErrUserNotExist
	}

	return
}

func (d *userStore) GetUserDetails(ctx context.Context, email string, password string) (user []User, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return d.db.SelectContext(ctx, &user, getUserDetailsQuery, email, password)
	})
	if len(user) == 0 {
		return user, ErrUserNotExist
	}
	return
}

func (d *userStore) GetUserDetailsById(ctx context.Context, userId string) (users []User, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return d.db.SelectContext(ctx, &users, getUserDetailsByIdQuery, userId)
	})
	if len(users) == 0 {
		return users, ErrUserNotExist
	}
	return
}

func (d *userStore) UpdateUser(ctx context.Context) {

}
