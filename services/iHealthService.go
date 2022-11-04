package services

import "github.com/AlejandroAldana99/YoFio_API/models"

type IHealthService interface {
	CheckPod(chanHealth chan models.HealthComponentDetail)
}
