package repository

import (
	"fmt"
	"github.com/pkg/errors"
	"nunu-project/internal/model"
)

type BlogRepository interface {
	FirstById(id int64) (*model.Blog, error)
	GetBlogList() ([]model.Blog, error)
}
type blogRepository struct {
	*Repository
}

func NewBlogRepository(repository *Repository) BlogRepository {
	return &blogRepository{
		Repository: repository,
	}
}

func (r *blogRepository) FirstById(id int64) (*model.Blog, error) {
	var blog model.Blog
	if err := r.db.First(&blog, id).Error; err != nil {

		return nil, errors.Wrap(err, "failed to get user by ID")
	}
	return &blog, nil
}

func (r *blogRepository) GetBlogList() ([]model.Blog, error) {
	var blogList []model.Blog
	result := r.db.Find(&blogList)
	r.logger.Info(fmt.Sprintf("blogList:%v", blogList))
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "failed to get list!")
	}

	return blogList, nil
}
