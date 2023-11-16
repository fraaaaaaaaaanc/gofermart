package utilsmodels

import "errors"

var IsNotLuhnValid = errors.New("the number did not pass the algorithm Luhn")

var IsNotNumber = errors.New("the string contains characters other than numbers")

var IsNilOrderNumber = errors.New("an empty order number has been sent")
