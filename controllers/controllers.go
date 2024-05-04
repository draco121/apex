package controllers

import (
	"github.com/draco121/common/models"
	"github.com/draco121/projectmanagerservice/core"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Controllers struct {
	service core.IProjectService
}

func NewControllers(service core.IProjectService) Controllers {
	c := Controllers{
		service: service,
	}
	return c
}

func (s *Controllers) CreateProject(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBind(&project); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	} else {
		res, err := s.service.CreateProject(c, &project)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(201, res)
		}
	}
}

func (s *Controllers) DeleteProject(c *gin.Context) {
	if projectId, ok := c.GetQuery("projectId"); !ok {
		c.JSON(400, gin.H{
			"message": "project id not provided",
		})
	} else {
		projectId, err := primitive.ObjectIDFromHex(projectId)
		if err != nil {
			c.JSON(400, err.Error())
		}
		res, err := s.service.DeleteProject(c, projectId)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(201, res)
		}
	}

}

func (s *Controllers) GetProject(c *gin.Context) {
	if projectId, ok := c.GetQuery("projectId"); ok {
		projectId, err := primitive.ObjectIDFromHex(projectId)
		if err != nil {
			c.JSON(400, err.Error())
		}
		res, err := s.service.GetProjectById(c, projectId)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, res)
		}

	} else {
		res, err := s.service.GetProjectsByUserId(c)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, res)
		}
	}
}
