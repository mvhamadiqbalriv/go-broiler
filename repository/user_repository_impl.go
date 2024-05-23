package repository

import (
	"context"
	"database/sql"
	"errors"
	"mvhamadiqbalriv/belajar-golang-restful-api/helper"
	"mvhamadiqbalriv/belajar-golang-restful-api/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING id"
    var id int

	//encrypt password bcrypt
	password := helper.HashAndSalt([]byte(user.Password))

    err := tx.QueryRowContext(ctx, SQL, user.Name, user.Email, password).Scan(&id)
    helper.PanicIfError(err)

    user.ID = id
    return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE users SET name = $1, email = $2 WHERE id = $3"
	_, err := tx.ExecContext(ctx, SQL, user.Name, user.Email, user.ID)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "DELETE FROM users WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, user.ID)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "SELECT id, name, email, password, profile_picture FROM users WHERE id = $1"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.ProfilePicture)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	SQL := "SELECT id, name, email, password FROM users WHERE email = $1"
	rows, err := tx.QueryContext(ctx, SQL, email)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "SELECT id, name, email FROM users"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		helper.PanicIfError(err)
		users = append(users, user)
	}

	return users
}

func (repository *UserRepositoryImpl) ChangePassword(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE users SET password = $1 WHERE id = $2"

	//encrypt password bcrypt
	user.Password = helper.HashAndSalt([]byte(user.Password))

	_, err := tx.ExecContext(ctx, SQL, user.Password, user.ID)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) ChangeProfilePicture(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE users SET profile_picture = $1 WHERE id = $2"
	_, err := tx.ExecContext(ctx, SQL, user.ProfilePicture, user.ID)
	helper.PanicIfError(err)

	return user
}
