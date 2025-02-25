package api

import (
	"github.com/Zulhaidir/microservice/mantap/util"
	"github.com/go-playground/validator/v10"
)

/* Membuat validator untuk currency sebelum diterapkan pada transfers */
var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
