package controllers

import (
	"net/http"
	"strings"

	"github.com/AlejandroAldana99/YoFio_API/models"
	"github.com/AlejandroAldana99/YoFio_API/services"
	"github.com/labstack/echo/v4"
)

type ControllerData struct {
	Service services.IAssigmentService
}

func (controller ControllerData) GetAssigmentData(c echo.Context) error {
	AssigmentID := strings.ToLower(c.Param("id"))
	data, err := controller.Service.GetAssigment(AssigmentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, data)
}

func (controller ControllerData) CreateAssigmentData(c echo.Context) error {
	dto := c.Get("dto").(models.AssigmentData)
	response, err := controller.Service.CreateAssigment(dto)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, response)
}
