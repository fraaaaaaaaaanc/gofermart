package utils_models

import "errors"

var ErrIsNotLuhnValid = errors.New("the number did not pass the algorithm Luhn")

var ErrIsNotNumber = errors.New("the string contains characters other than numbers")

var ErrIsNilOrderNumber = errors.New("an empty order number has been sent")
