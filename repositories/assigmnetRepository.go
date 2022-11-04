package repositories

import (
	"context"
	"time"

	"github.com/AlejandroAldana99/YoFio_API/config"
	"github.com/AlejandroAldana99/YoFio_API/errors"
	"github.com/AlejandroAldana99/YoFio_API/libs/logger"
	"github.com/AlejandroAldana99/YoFio_API/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AssigmentRepository struct {
	Config  config.Configuration
	MongoDB *mongo.Database
}

func (repo AssigmentRepository) GetAssigment(AssigmentID string) (models.AssigmentData, error) {
	t := time.Now()
	var Assigment models.AssigmentData
	objectId, oErr := primitive.ObjectIDFromHex(AssigmentID)
	if oErr != nil {
		logger.Error("repositories", "GetAssigment", oErr.Error())
		return models.AssigmentData{}, errors.HandleServiceError(oErr)
	}
	err := repo.MongoDB.Collection("Assigments").FindOne(
		context.TODO(),
		bson.D{{Key: "_id", Value: objectId}},
	).Decode(&Assigment)

	if err != nil {
		logger.Error("repositories", "GetAssigmentData", err.Error())
		return models.AssigmentData{}, errors.HandleServiceError(err)
	}

	logger.Performance("repository", "GetAssigment", t)

	return Assigment, nil
}

func (repo AssigmentRepository) CreateAssigment(data models.AssigmentData) error {

	t := time.Now()
	_, err := repo.MongoDB.Collection("Assigments").InsertOne(context.TODO(), data)
	if err != nil {
		logger.Error("repositories", "CreateAssigment", err.Error())
		return errors.HandleServiceError(err)
	}

	logger.Performance("repository", "CreateAssigment", t)

	return nil
}
