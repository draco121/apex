package core

import (
	"apex/repository"
	"context"
	"github.com/draco121/horizon/models"
	"github.com/draco121/horizon/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IProjectService interface {
	CreateProject(ctx context.Context, user *models.Project) (*models.Project, error)
	DeleteProject(ctx context.Context, id primitive.ObjectID) (*models.Project, error)
	GetProjectByName(ctx context.Context, name string) (*models.Project, error)
	GetProjectById(ctx context.Context, id primitive.ObjectID) (*models.Project, error)
	GetProjectsByUserId(ctx context.Context) (*[]models.Project, error)
}

type projectService struct {
	IProjectService
	repo   repository.IProjectRepository
	client *mongo.Client
}

func NewProjectService(client *mongo.Client, repository repository.IProjectRepository) IProjectService {
	us := projectService{
		repo:   repository,
		client: client,
	}
	return &us
}

func (s *projectService) CreateProject(ctx context.Context, project *models.Project) (*models.Project, error) {
	session, err := s.client.StartSession()
	if err != nil {
		utils.Logger.Error("failed to start mongo session", "error: ", err.Error())
		return nil, err
	}
	defer session.EndSession(ctx)
	err = session.StartTransaction()
	if err != nil {
		utils.Logger.Error("failed to start mongo transaction", "error: ", err.Error())
		return nil, err
	}
	project, err = s.repo.InsertOne(ctx, project)
	if err != nil {
		utils.Logger.Error("failed to insert project", "error: ", err.Error())
		return nil, err
	} else {
		_ = session.CommitTransaction(ctx)
		utils.Logger.Info("inserted project", "project", project)
		return project, nil
	}
}

func (s *projectService) DeleteProject(ctx context.Context, id primitive.ObjectID) (*models.Project, error) {
	session, err := s.client.StartSession()
	if err != nil {
		utils.Logger.Error("failed to start mongo session", "error: ", err.Error())
		return nil, err
	}
	defer session.EndSession(ctx)
	err = session.StartTransaction()
	if err != nil {
		utils.Logger.Error("failed to start mongo transaction", "error: ", err.Error())
		return nil, err
	}
	project, err := s.repo.DeleteOneById(ctx, id)
	if err != nil {
		utils.Logger.Error("failed to delete project", "error: ", err.Error())
		return nil, err
	} else {
		_ = session.CommitTransaction(ctx)
		utils.Logger.Info("deleted project", "project", project)
		return project, nil
	}
}

func (s *projectService) GetProjectByName(ctx context.Context, name string) (*models.Project, error) {
	session, err := s.client.StartSession()
	if err != nil {
		utils.Logger.Error("failed to start mongo session", "error: ", err.Error())
		return nil, err
	}
	defer session.EndSession(ctx)
	err = session.StartTransaction()
	if err != nil {
		utils.Logger.Error("failed to start mongo transaction", "error: ", err.Error())
		return nil, err
	}
	project, err := s.repo.FindOneByName(ctx, name)
	if err != nil {
		utils.Logger.Error("failed to find project", "error: ", err.Error())
		return nil, err
	} else {
		_ = session.CommitTransaction(ctx)
		utils.Logger.Info("found project", "project", project)
		return project, nil
	}
}

func (s *projectService) GetProjectById(ctx context.Context, id primitive.ObjectID) (*models.Project, error) {
	session, err := s.client.StartSession()
	if err != nil {
		utils.Logger.Error("failed to start mongo session", "error: ", err.Error())
		return nil, err
	}
	defer session.EndSession(ctx)
	err = session.StartTransaction()
	if err != nil {
		utils.Logger.Error("failed to start mongo transaction", "error: ", err.Error())
		return nil, err
	}
	project, err := s.repo.FindOneById(ctx, id)
	if err != nil {
		utils.Logger.Error("failed to find project", "error: ", err.Error())
		return nil, err
	} else {
		_ = session.CommitTransaction(ctx)
		utils.Logger.Info("found project", "project", project)
		return project, nil
	}
}

func (s *projectService) GetProjectsByUserId(ctx context.Context) (*[]models.Project, error) {
	session, err := s.client.StartSession()
	if err != nil {
		utils.Logger.Error("failed to start mongo session", "error: ", err.Error())
		return nil, err
	}
	defer session.EndSession(ctx)
	err = session.StartTransaction()
	if err != nil {
		utils.Logger.Error("failed to start mongo transaction", "error: ", err.Error())
		return nil, err
	}
	projects, err := s.repo.FindManyByUserId(ctx)
	if err != nil {
		utils.Logger.Error("failed to find projects", "error: ", err.Error())
		return nil, err
	} else {
		_ = session.CommitTransaction(ctx)
		utils.Logger.Info("found projects", "projects", projects)
		return projects, nil
	}

}
