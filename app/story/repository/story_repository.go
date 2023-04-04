package repository

import (
	"context"

	"github.com/VinncentWong/DiverBE/domain"
	"github.com/VinncentWong/DiverBE/infrastructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IStoryRepository interface {
	CreateStory(story domain.Story) (domain.Story, error)
	GetStory(id primitive.ObjectID) (domain.Story, error)
	DeleteStory(id primitive.ObjectID) error
	UpdateStory(id primitive.ObjectID, story domain.Story) error
	GetStories() ([]domain.Story, error)
}

type StoryRepository struct {
	db *mongo.Client
}

func NewStoryRepository() IStoryRepository {
	return &StoryRepository{
		db: infrastructure.GetClient(),
	}
}

func (r *StoryRepository) CreateStory(story domain.Story) (domain.Story, error) {
	coll := r.db.Database("jointscamp").Collection("story")
	_, err := coll.InsertOne(context.Background(), story)
	if err != nil {
		return domain.Story{}, err
	}
	return story, nil
}

func (r *StoryRepository) DeleteStory(id primitive.ObjectID) error {
	filter := bson.D{
		{
			Key:   "_id",
			Value: id,
		},
	}
	coll := r.db.Database("jointscamp").Collection("story")
	_, err := coll.DeleteOne(context.Background(), filter)
	return err
}

func (r *StoryRepository) GetStory(id primitive.ObjectID) (domain.Story, error) {
	filter := bson.D{
		{
			Key:   "_id",
			Value: id,
		},
	}
	coll := r.db.Database("jointscamp").Collection("story")
	result := coll.FindOne(context.Background(), filter)
	if result.Err() != nil {
		return domain.Story{}, result.Err()
	}
	var container domain.Story
	err := result.Decode(&container)
	if err != nil {
		return domain.Story{}, err
	}
	return container, nil
}

func (r *StoryRepository) UpdateStory(id primitive.ObjectID, story domain.Story) error {
	filter := bson.D{
		{
			Key:   "_id",
			Value: id,
		},
	}
	coll := r.db.Database("jointscamp").Collection("story")
	_, err := coll.ReplaceOne(context.Background(), filter, story)
	if err != nil {
		return err
	}
	return nil
}

func (r *StoryRepository) GetStories() ([]domain.Story, error) {
	var container []domain.Story
	coll := r.db.Database("jointscamp").Collection("story")
	result, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		return []domain.Story{}, err
	}
	err = result.All(context.Background(), &container)
	if err != nil {
		return []domain.Story{}, err
	}
	return container, nil
}
