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
	GetContacts() ([]domain.Contact, error)
	GetContact(id primitive.ObjectID) (domain.Contact, error)
	DeleteContact(id primitive.ObjectID) error
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

func (u *ContactUsecase) GetContacts() ([]domain.Contact, error) {
	result, err := u.r.GetContacts()
	return result, err
}

func (u *ContactUsecase) GetContact(id primitive.ObjectID) (domain.Contact, error) {
	result, err := u.r.GetContact(id)
	return result, err
}

func (u *ContactUsecase) DeleteContact(id primitive.ObjectID) error {
	return u.r.DeleteContact(id)
}
