package util

// Constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	RUB = "RUB"
	CAD = "CAD"
)

// InSupportedCurrency return true if the currency is supported
func InSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, RUB, CAD:
		return true
	}
	return false
}
