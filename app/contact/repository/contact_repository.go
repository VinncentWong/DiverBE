package repository

import (
	"context"

	"github.com/VinncentWong/DiverBE/domain"
	"github.com/VinncentWong/DiverBE/infrastructure"
	"go.mongodb.org/mongo-driver/mongo"
)

type IContactRepository interface {
	CreateContact(contact domain.Contact) (domain.Contact, error)
}

type ContactRepository struct {
	db *mongo.Client
}

func NewContactRepository() IContactRepository {
	return &ContactRepository{
		db: infrastructure.GetClient(),
	}
}

func (r *ContactRepository) CreateContact(contact domain.Contact) (domain.Contact, error) {
	coll := r.db.Database("jointscamp").Collection("contact")
	_, err := coll.InsertOne(context.Background(), contact)
	if err != nil {
		return domain.Contact{}, err
	}
	return contact, nil
}
