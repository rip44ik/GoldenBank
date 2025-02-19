package api

import (
	"github.com/go-playground/validator/v10"
	"gitlab.com/xfx1/goldbank/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check currency is supported
		return util.InSupportedCurrency(currency)
	}
	return false
}
