package repositories

import (
	"go-commerce/schemas"

	"gorm.io/gorm"
)

type UserRepository interface {
    Create(user *schemas.User) error
    Delete(user *schemas.User) error
    List() ([]schemas.User, error)
    FindByID(id string) (schemas.User, error)
    Update(user *schemas.User) error
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) Create(user *schemas.User) error {
    return r.db.Create(user).Error
}

func (r *userRepository) Delete(user *schemas.User) error {
    return r.db.Delete(user).Error
}

func (r *userRepository) List() ([]schemas.User, error) {
    var users []schemas.User
    err := r.db.Find(&users).Error
    return users, err
}

func (r *userRepository) FindByID(id string) (schemas.User, error) {
    var user schemas.User
    err := r.db.First(&user, id).Error
    return user, err
}

func (r *userRepository) Update(user *schemas.User) error {
    return r.db.Save(user).Error
}
