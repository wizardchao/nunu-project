package repository

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"nunu-project/internal/model"
	"nunu-project/pkg/helper/localTime"
	"time"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id string) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	GetList(ctx context.Context, page int, pageSize int, options ...UserOptions) ([]*model.User, error)
	UpdateWhere(ctx context.Context, user *model.User, userMap map[string]interface{}) error
}

type userRepository struct {
	*Repository
}

type UserOptions func(r *gorm.DB) *gorm.DB

func WithUserNickname(nickname string) UserOptions {
	return func(r *gorm.DB) *gorm.DB {
		return r.Where("nickname  like ?", "%"+nickname+"%")
	}
}

func WithEmail(email string) UserOptions {
	return func(r *gorm.DB) *gorm.DB {
		return r.Where("email like ?", "%"+email+"%")
	}
}

func WithUserId(userId string) UserOptions {
	return func(r *gorm.DB) *gorm.DB {
		return r.Where("user_id like ?", "%"+userId+"%")
	}
}

func WithOrder(order string) UserOptions {
	return func(r *gorm.DB) *gorm.DB {
		return r.Order(order)
	}
}

func WithUsername(username string) UserOptions {
	return func(r *gorm.DB) *gorm.DB {
		return r.Where("username like ?", username+"%")
	}
}

func NewUserRepository(r *Repository) UserRepository {
	return &userRepository{
		Repository: r,
	}
}
func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	user.CreatedAt = localTime.LocalTime(time.Now())
	user.UpdatedAt = localTime.LocalTime(time.Now())
	if err := r.db.Create(user).Error; err != nil {
		return errors.Wrap(err, "failed to create user")
	}
	return nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	user.UpdatedAt = localTime.LocalTime(time.Now())
	if err := r.db.Save(user).Error; err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil
}

func (r *userRepository) UpdateWhere(ctx context.Context, user *model.User, userMap map[string]interface{}) error {
	re := r.db.Debug().Model(&user).Updates(userMap)
	if re.Error != nil {
		return errors.Wrap(re.Error, "failed to update user")
	}
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, userId string) (*model.User, error) {
	var user model.User
	if err := r.db.Debug().Where("user_id = ?", userId).First(&user).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get user by ID")
	}

	return &user, nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get user by username")
	}

	return &user, nil
}

func (r userRepository) GetList(ctx context.Context, page int, pageSize int, userOption ...UserOptions) (userList []*model.User, err error) {
	userModel := r.db
	for _, opt := range userOption {
		userModel = opt(userModel)
	}

	offset := (page - 1) * pageSize
	result := userModel.Debug().Offset(offset).Limit(pageSize).Find(&userList)
	err = result.Error
	return
}
