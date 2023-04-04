package usecase

import (
	"mime/multipart"
	"time"

	_photoRepository "github.com/VinncentWong/DiverBE/app/photo/repository"
	_storyRepository "github.com/VinncentWong/DiverBE/app/story/repository"
	"github.com/VinncentWong/DiverBE/domain"
	"github.com/VinncentWong/DiverBE/domain/dto/story"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IStoryUsecase interface {
	CreateStory(dto story.CreateStoryDto, file multipart.File) (domain.Story, error)
	DeleteStory(id string) error
	UpdateStory(id string, file *multipart.File, dto domain.Story) error
	GetStories() ([]domain.Story, error)
	GetStory(id primitive.ObjectID) (domain.Story, error)
}

type StoryUsecase struct {
	storyRepo _storyRepository.IStoryRepository
	photoRepo _photoRepository.IPhotoRepository
}

func NewStoryUsecase(storyRepo _storyRepository.IStoryRepository, photoRepo _photoRepository.IPhotoRepository) IStoryUsecase {
	return &StoryUsecase{
		storyRepo: storyRepo,
		photoRepo: photoRepo,
	}
}

func (u *StoryUsecase) CreateStory(dto story.CreateStoryDto, file multipart.File) (domain.Story, error) {
	secureUrl, publicId, err := u.photoRepo.UploadPhoto(file)
	if err != nil {
		return domain.Story{}, err
	}
	story := domain.Story{
		ID:          primitive.NewObjectID(),
		Title:       dto.Title,
		Description: dto.Description,
		CreatedAt:   time.Now(),
		Photo:       secureUrl,
		PublicID:    publicId,
	}
	result, err := u.storyRepo.CreateStory(story)
	if err != nil {
		return domain.Story{}, err
	}
	return result, nil
}

func (u *StoryUsecase) DeleteStory(id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := u.storyRepo.GetStory(objId)
	if err != nil {
		return err
	}
	err = u.photoRepo.DeletePhoto(result.PublicID)
	if err != nil {
		return err
	}
	err = u.storyRepo.DeleteStory(objId)
	return err
}

func (u *StoryUsecase) UpdateStory(id string, file *multipart.File, dto domain.Story) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := u.storyRepo.GetStory(objId)
	if err != nil {
		return err
	}
	var container domain.Story
	if file != nil {
		secureUrl, publicId, err := u.photoRepo.UploadPhoto(*file)
		if err != nil {
			return err
		}
		container = result
		container.PublicID = publicId
		container.Photo = secureUrl
	} else {
		container.Photo = result.Photo
		container.PublicID = result.PublicID
	}
	if dto.Title != "" {
		container.Title = dto.Title
	} else {
		container.Title = result.Title
	}
	if dto.Description != "" {
		container.Description = dto.Description
	} else {
		container.Description = result.Description
	}
	container.ID = objId
	container.CreatedAt = result.CreatedAt
	container.UpdatedAt = time.Now()
	err = u.storyRepo.UpdateStory(objId, container)
	return err
}

func (u *StoryUsecase) GetStories() ([]domain.Story, error) {
	result, err := u.storyRepo.GetStories()
	return result, err
}

func (u *StoryUsecase) GetStory(id primitive.ObjectID) (domain.Story, error) {
	result, err := u.storyRepo.GetStory(id)
	return result, err
}
