package util

const (
	USD = "USD"
	NGN = "NGN"
	EUR = "EUR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, NGN, EUR:
		return true
	}
	return false
}