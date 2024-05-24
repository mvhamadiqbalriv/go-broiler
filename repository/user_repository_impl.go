package repository

import (
	"context"
	"errors"
	"mvhamadiqbalriv/belajar-golang-restful-api/helper"
	"mvhamadiqbalriv/belajar-golang-restful-api/model/domain"
	"mvhamadiqbalriv/belajar-golang-restful-api/pkg"
	"net/http"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
    return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, user domain.User) domain.User {
    // Encrypt password using helper function
    user.Password = helper.HashAndSalt([]byte(user.Password))
    tx.Create(&user)
    return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, user domain.User) domain.User {
    tx.Save(&user)
    return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, user domain.User) {
    tx.Delete(&user)
}

func (repository *UserRepositoryImpl) FindByID(ctx context.Context, tx *gorm.DB, userId int) (domain.User, error) {
    var user domain.User
    result := tx.First(&user, userId)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return user, errors.New("user not found")
        }
        return user, result.Error
    }
    return user, nil
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *gorm.DB, email string) (domain.User, error) {
    var user domain.User
    result := tx.Where("email = ?", email).First(&user)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return user, errors.New("user not found")
        }
        return user, result.Error
    }
    return user, nil
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB, r *http.Request) (*pkg.PaginationImpl, error) {
    var users []domain.User

	// Extract query parameters
    pagination := pkg.ExtractQueryParams(r)

	// Apply search filters
    tx = applyFilters(tx, r)

    // Apply pagination and retrieve users
    tx.Scopes(pkg.Paginate(users, &pagination, tx)).Find(&users)

    // Convert users to the response format
    pagination.Rows = helper.ToUserResponses(users)

    return &pagination, nil
}

func (repository *UserRepositoryImpl) ChangePassword(ctx context.Context, tx *gorm.DB, user domain.User) domain.User {
    // Encrypt password using helper function
    user.Password = helper.HashAndSalt([]byte(user.Password))
    tx.Save(&user)
    return user
}

func (repository *UserRepositoryImpl) ChangeProfilePicture(ctx context.Context, tx *gorm.DB, user domain.User) domain.User {
    tx.Model(&user).Update("profile_picture", user.ProfilePicture)
    return user
}

// Function to apply search filters to the query
func applyFilters(tx *gorm.DB, r *http.Request) *gorm.DB {
	query := r.URL.Query()

    searchName := query.Get("search_name")
    if searchName != "" {
        tx = tx.Where("name ILIKE ?", "%"+searchName+"%")
    }
    searchEmail := query.Get("search_email")
    if searchEmail != "" {
        tx = tx.Where("email ILIKE ?", "%"+searchEmail+"%")
    }
    return tx
}
