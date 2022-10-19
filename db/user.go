package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const (
	CreateUserQuery             = `INSERT INTO users(name, email, password, role, created_at, updated_at) VALUES($1,$2,$3,$4,$5,$6)`
	getUserListQuery            = `SELECT * FROM users`
	getUserDetailsQuery         = `SELECT * FROM users WHERE email=$1 AND password=$2`
	getUserDetailsByIdQuery     = `SELECT * FROM users WHERE id=$1`
	UpdateUserQueryWithName     = `UPDATE users SET name=$1 updated_at=$2 WHERE ID = $3 RETURNING *`
	UpdateUserQueryWithPassword = `UPDATE users SET password=$1 updated_at=$2 WHERE ID = $3 RETURNING *`
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
		//returns multiple rows
		return d.db.SelectContext(ctx, &users, getUserListQuery)
	})
	if err == sql.ErrNoRows {
		return users, ErrUserNotExist
	}

	return
}

func (d *userStore) GetUserDetails(ctx context.Context, email string, password string) (user User, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return d.db.GetContext(ctx, &user, getUserDetailsQuery, email, password)
	})
	if err == sql.ErrNoRows {
		return user, ErrUserNotExist
	}
	return
}

func (d *userStore) GetUserDetailsById(ctx context.Context, userId string) (user User, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return d.db.GetContext(ctx, &user, getUserDetailsByIdQuery, userId)
	})
	if err == sql.ErrNoRows {
		return user, ErrUserNotExist
	}
	return
}

func (d *userStore) UpdateUserDetailsById(ctx context.Context, userId string, name string, password string) (user User, err error) {
	now := time.Now()
	if name != "" {
		err = d.db.GetContext(ctx, &user, UpdateUserQueryWithName, name, now, userId)
	}
	if password != "" {
		err = d.db.GetContext(ctx, &user, UpdateUserQueryWithPassword, password, now, userId)
	}
	return
}
