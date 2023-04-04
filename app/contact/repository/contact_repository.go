package repository

import (
	"context"

	"github.com/VinncentWong/DiverBE/domain"
	"github.com/VinncentWong/DiverBE/infrastructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IContactRepository interface {
	CreateContact(contact domain.Contact) (domain.Contact, error)
	GetContacts() ([]domain.Contact, error)
	GetContact(id primitive.ObjectID) (domain.Contact, error)
	DeleteContact(id primitive.ObjectID) error
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

func (r *ContactRepository) GetContacts() ([]domain.Contact, error) {
	var container []domain.Contact
	coll := r.db.Database("jointscamp").Collection("contact")
	result, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		return []domain.Contact{}, err
	}
	err = result.All(context.Background(), &container)
	if err != nil {
		return []domain.Contact{}, err
	}
	return container, nil
}

func (r *ContactRepository) GetContact(id primitive.ObjectID) (domain.Contact, error) {
	filter := bson.D{
		{
			Key:   "_id",
			Value: id,
		},
	}
	coll := r.db.Database("jointscamp").Collection("contact")
	var container domain.Contact
	result := coll.FindOne(context.Background(), filter)
	if result.Err() != nil {
		return domain.Contact{}, result.Err()
	}
	err := result.Decode(&container)
	if err != nil {
		return domain.Contact{}, result.Err()
	}
	return container, nil
}

func (r *ContactRepository) DeleteContact(id primitive.ObjectID) error {
	filter := bson.D{
		{
			Key:   "_id",
			Value: id,
		},
	}
	coll := r.db.Database("jointscamp").Collection("contact")
	_, err := coll.DeleteOne(context.Background(), filter)
	return err
}
