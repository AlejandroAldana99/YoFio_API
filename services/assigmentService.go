package services

import (
	"github.com/AlejandroAldana99/YoFio_API/errors"
	"github.com/AlejandroAldana99/YoFio_API/libs/logger"
	"github.com/AlejandroAldana99/YoFio_API/models"
	"github.com/AlejandroAldana99/YoFio_API/repositories"
)

const millisecondsEq = 1000000.0

type AssigmentService struct {
	Repository repositories.IAssigmentRepository
}

func (service AssigmentService) GetAssigment(AssigmentID string) (models.AssigmentData, error) {
	Assigment, err := service.Repository.GetAssigment(AssigmentID)
	if err != nil {
		logger.Error("services", "GetAssigment", err.Error())
		return Assigment, errors.HandleServiceError(err)
	}

	return Assigment, nil
}

func (service AssigmentService) CreateAssigment(data models.AssigmentData) (models.ResponseData, error) {
	response := models.ResponseData{
		Status: "Faild",
	}

	data.Combinations = calculateLoans(data.OriginalInvestment)

	err := service.Repository.CreateAssigment(data)
	if err != nil {
		logger.Error("services", "CreateAssigment", err.Error())
		return response, errors.HandleServiceError(err)
	}

	response.Status = "Success"
	return response, nil
}
