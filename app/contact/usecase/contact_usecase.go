package usecase

import (
	"time"

	_contactRepository "github.com/VinncentWong/DiverBE/app/contact/repository"
	"github.com/VinncentWong/DiverBE/domain"
	_contactDto "github.com/VinncentWong/DiverBE/domain/dto/contact"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IContactUsecase interface {
	CreateContact(dto _contactDto.CreateContactDto) (domain.Contact, error)
}

type ContactUsecase struct {
	r _contactRepository.IContactRepository
}

func NewContactUsecase(r _contactRepository.IContactRepository) IContactUsecase {
	return &ContactUsecase{
		r: r,
	}
}

func (u *ContactUsecase) CreateContact(dto _contactDto.CreateContactDto) (domain.Contact, error) {
	contact := domain.Contact{
		ID:          primitive.NewObjectID(),
		Name:        dto.Name,
		Description: dto.Message,
		Email:       dto.Email,
		CreatedAt:   time.Now(),
	}
	result, err := u.r.CreateContact(contact)
	if err != nil {
		return domain.Contact{}, err
	}
	return result, nil
}
