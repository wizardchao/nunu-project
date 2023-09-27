package service

import (
	"nunu-project/internal/model"
	"nunu-project/internal/repository"
)

type BlogService interface {
	GetBlogById(id int64) (*model.Blog, error)
	GetBlogList() ([]model.Blog, error)
}

type blogService struct {
	*Service
	blogRepository repository.BlogRepository
}

func NewBlogService(service *Service, blogRepository repository.BlogRepository) BlogService {
	return &blogService{
		Service:        service,
		blogRepository: blogRepository,
	}
}

func (s *blogService) GetBlogById(id int64) (*model.Blog, error) {
	return s.blogRepository.FirstById(id)
}

func (s *blogService) GetBlogList() ([]model.Blog, error) {
	return s.blogRepository.GetBlogList()
}
