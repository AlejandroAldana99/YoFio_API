package repositories

import "github.com/AlejandroAldana99/YoFio_API/models"

type IAssigmentRepository interface {
	GetAssigment(assigmentID string) (models.AssigmentData, error)
	CreateAssigment(data models.AssigmentData) error
	GetStatistics() (models.StatisticsData, error)
}
