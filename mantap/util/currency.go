package util

/* Dapat menambahkan currency kedepannya, saat ini 3 saja dulu */
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

/* isSupportedCurrency me-return true jika mensupport currency */
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}
