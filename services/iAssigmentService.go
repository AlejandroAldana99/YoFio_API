package services

import "github.com/AlejandroAldana99/YoFio_API/models"

type IAssigmentService interface {
	GetAssigment(AssigmentID string) (models.AssigmentData, error)
	CreateAssigment(data models.AssigmentData) (models.ResponseData, error)
}
