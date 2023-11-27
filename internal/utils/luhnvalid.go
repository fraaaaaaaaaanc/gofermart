package utils

import (
	"gofermart/internal/models/utils_models"
	"strconv"
)

func IsLuhnValid(orderNumber string) error {
	sum := 0
	lenOrderNumber := len(orderNumber) - 1
	if lenOrderNumber < 0 {
		return utilsmodels.ErrIsNilOrderNumber
	}

	for i := lenOrderNumber; i >= 0; i-- {
		digit, err := strconv.Atoi(string(orderNumber[i]))
		if err != nil {
			return utilsmodels.ErrIsNotNumber
		}

		if (lenOrderNumber-i)%2 != 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
	}
	if sum%10 != 0 {
		return utilsmodels.ErrIsNotLuhnValid
	}
	return nil
}
