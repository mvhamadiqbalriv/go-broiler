package repository

import (
	"context"
	"mvhamadiqbalriv/belajar-golang-restful-api/model/domain"
	"mvhamadiqbalriv/belajar-golang-restful-api/pkg"
	"net/http"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, tx *gorm.DB, user domain.User) domain.User
	Update(ctx context.Context, tx *gorm.DB, user domain.User) domain.User
	Delete(ctx context.Context, tx *gorm.DB, user domain.User)
	FindByID(ctx context.Context, tx *gorm.DB, userId int) (domain.User, error)
	FindByEmail(ctx context.Context, tx *gorm.DB, email string) (domain.User, error)
	FindAll(ctx context.Context, tx *gorm.DB, r *http.Request) (*pkg.PaginationImpl, error)
	
	//ChangeProfilePicture
	ChangeProfilePicture(ctx context.Context, tx *gorm.DB, user domain.User) domain.User
	ChangePassword(ctx context.Context, tx *gorm.DB, user domain.User) domain.User
}