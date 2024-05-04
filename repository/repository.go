package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/draco121/common/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IProjectRepository interface {
	InsertOne(ctx context.Context, project *models.Project) (*models.Project, error)
	FindOneById(ctx context.Context, id primitive.ObjectID) (*models.Project, error)
	FindOneByName(ctx context.Context, name string) (*models.Project, error)
	DeleteOneById(ctx context.Context, id primitive.ObjectID) (*models.Project, error)
	FindManyByUserId(ctx context.Context) (*[]models.Project, error)
}

type projectRepository struct {
	IProjectRepository
	db *mongo.Database
}

func NewProjectRepository(database *mongo.Database) IProjectRepository {
	return &projectRepository{db: database}
}

func (ur *projectRepository) InsertOne(ctx context.Context, project *models.Project) (*models.Project, error) {
	userId := ctx.Value("UserId").(primitive.ObjectID)
	project.Owner = userId
	result, _ := ur.FindOneByName(ctx, project.Name)
	if result != nil {
		return nil, fmt.Errorf("record exists")
	} else {
		project.ID = primitive.NewObjectID()
		_, err := ur.db.Collection("projects").InsertOne(ctx, project)
		if err != nil {
			return nil, err
		}
	}
	return project, nil

}

func (ur *projectRepository) FindOneById(ctx context.Context, id primitive.ObjectID) (*models.Project, error) {
	filter := bson.D{{Key: "_id", Value: id}, {Key: "Owner"}}
	result := models.Project{}
	err := ur.db.Collection("projects").FindOne(ctx, filter).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else {
		return &result, nil
	}
}

func (ur *projectRepository) FindOneByName(ctx context.Context, name string) (*models.Project, error) {
	userId := ctx.Value("UserId").(primitive.ObjectID)
	filter := bson.D{{Key: "name", Value: name}, {Key: "owner", Value: userId}}
	result := models.Project{}
	err := ur.db.Collection("projects").FindOne(ctx, filter).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else {
		return &result, nil
	}
}

func (ur *projectRepository) DeleteOneById(ctx context.Context, id primitive.ObjectID) (*models.Project, error) {
	userId := ctx.Value("UserId").(primitive.ObjectID)
	filter := bson.D{{Key: "_id", Value: id}, {Key: "owner", Value: userId}}
	result := models.Project{}
	err := ur.db.Collection("projects").FindOneAndDelete(ctx, filter).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else {
		return &result, nil
	}
}

func (ur *projectRepository) FindManyByUserId(ctx context.Context) (*[]models.Project, error) {
	userId := ctx.Value("UserId").(primitive.ObjectID)
	filter := bson.D{{Key: "owner", Value: userId}}
	var projects []models.Project
	cur, err := ur.db.Collection("projects").Find(ctx, filter)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else {
		if err = cur.All(context.TODO(), &projects); err != nil {
			return nil, err
		}
		return &projects, nil
	}
}
