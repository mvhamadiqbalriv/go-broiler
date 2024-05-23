package repository

import (
	"context"
	"database/sql"
	"mvhamadiqbalriv/belajar-golang-restful-api/model/domain"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, user domain.User)
	FindByID(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
	
	//ChangeProfilePicture
	ChangeProfilePicture(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	ChangePassword(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
}