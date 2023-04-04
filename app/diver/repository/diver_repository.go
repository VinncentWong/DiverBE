package repository

import (
	"context"
	"errors"
	"os"

	"github.com/VinncentWong/DiverBE/domain"
	_diverDto "github.com/VinncentWong/DiverBE/domain/dto/diver"
	"github.com/VinncentWong/DiverBE/infrastructure"
	"github.com/VinncentWong/DiverBE/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IDiverRepository interface {
	CreateUsernameIndex() error
	CreateDiver() (domain.Diver, error)
	LoginDiver(dto _diverDto.LoginDto) (domain.Diver, error)
	GetDiver(username string) (domain.Diver, error)
}

type DiverRepository struct {
	db *mongo.Client
}

func NewDiverRepository() IDiverRepository {
	return &DiverRepository{
		db: infrastructure.GetClient(),
	}
}

func (r *DiverRepository) CreateUsernameIndex() error {
	index := mongo.IndexModel{
		Keys: bson.D{
			{Key: "username", Value: 1},
		},
	}
	_, err := r.db.Database("jointscamp").Collection("diver").Indexes().CreateOne(context.TODO(), index)
	return err
}

func (r *DiverRepository) CreateDiver() (domain.Diver, error) {
	hashedPassword, err := util.HashPassword(os.Getenv("USER_PASSWORD"))
	if err != nil {
		return domain.Diver{}, err
	}
	coll := r.db.Database("jointscamp").Collection("diver")
	_id := primitive.NewObjectID()
	_, err = coll.InsertOne(context.Background(), bson.M{
		"_id":      _id,
		"username": os.Getenv("USER_USERNAME"),
		"password": hashedPassword,
		"email":    os.Getenv("USER_EMAIL"),
	})
	if err != nil {
		return domain.Diver{}, err
	}
	return domain.Diver{
		ID:       _id,
		Email:    os.Getenv("USER_EMAIL"),
		Password: hashedPassword,
		Username: os.Getenv("USER_USERNAME"),
	}, nil
}

func (r *DiverRepository) LoginDiver(dto _diverDto.LoginDto) (domain.Diver, error) {
	filter := bson.D{
		{
			Key:   "username",
			Value: dto.Username,
		},
	}
	result := r.db.Database("jointscamp").Collection("diver").FindOne(context.TODO(), filter)
	if result.Err() != nil {
		return domain.Diver{}, errors.New("diver doesn't exist")
	}
	var diver domain.Diver
	err := result.Decode(&diver)
	if err != nil {
		return domain.Diver{}, err
	}
	return diver, nil
}

func (r *DiverRepository) GetDiver(username string) (domain.Diver, error) {
	filter := bson.D{
		{
			Key:   "username",
			Value: username,
		},
	}
	var container domain.Diver
	result := r.db.Database("jointscamp").Collection("diver").FindOne(context.TODO(), filter)
	if result.Err() != nil {
		return domain.Diver{}, errors.New("diver doesn't exist")
	}
	err := result.Decode(&container)
	if err != nil {
		return domain.Diver{}, err
	}
	return container, nil
}
