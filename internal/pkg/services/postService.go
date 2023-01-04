package services

import (
	"github.com/yosa12978/pngb/internal/pkg/models"
	"github.com/yosa12978/pngb/internal/pkg/repositories"
)

type PostService interface {
	Find() []models.Post
	FindByID(id string) (models.Post, error)
	Create(p models.Post) error
	Delete(id string) error
}

type postService struct {
	postRepo repositories.PostRepository
}

func NewPostService(postRepo repositories.PostRepository) PostService {
	return &postService{
		postRepo: postRepo,
	}
}

func (services *postService) Find() []models.Post {
	return services.postRepo.Find()
}

func (services *postService) FindByID(id string) (models.Post, error) {
	return services.postRepo.FindByID(id)
}

func (services *postService) Create(p models.Post) error {
	return services.postRepo.Create(p)
}

func (services *postService) Delete(id string) error {
	return services.postRepo.Delete(id)
}
