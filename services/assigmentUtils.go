package services

import (
	"github.com/AlejandroAldana99/YoFio_API/constants"
	"github.com/AlejandroAldana99/YoFio_API/models"
)

func calculateLoans(invest float64) []models.CombinationsData {
	combinatios := []models.CombinationsData{}

	combinatios = append(combinatios, calculateFromHighest(invest))
	combinatios = append(combinatios, calculateFromLowest(invest))

	return combinatios
}

func calculateFromHighest(invest float64) models.CombinationsData {
	fromHighest := models.CombinationsData{}

	for {
		if doRest(invest, constants.Amount700) {
			invest = invest - constants.Amount700
			fromHighest.CreditType700++
		} else if doRest(invest, constants.Amount500) {
			invest = invest - constants.Amount500
			fromHighest.CreditType500++
		} else if doRest(invest, constants.Amount300) {
			invest = invest - constants.Amount300
			fromHighest.CreditType300++
		} else {
			break
		}
	}

	return fromHighest
}

func calculateFromLowest(invest float64) models.CombinationsData {
	fromLowest := models.CombinationsData{}

	for {
		if doRest(invest, constants.Amount300) {
			invest = invest - constants.Amount300
			fromLowest.CreditType300++
		} else if doRest(invest, constants.Amount500) {
			invest = invest - constants.Amount500
			fromLowest.CreditType500++
		} else if doRest(invest, constants.Amount700) {
			invest = invest - constants.Amount700
			fromLowest.CreditType700++
		} else {
			break
		}
	}

	return fromLowest
}

func doRest(num1, num2 float64) bool {
	result := (num1 - num2)
	return result >= 0.0
}
