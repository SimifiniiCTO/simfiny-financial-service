package consts

import (
	"github.com/plaid/plaid-go/plaid"
)

var (
	PlaidClientName = "simfiny"
	PlaidLanguage   = "en"
	PlaidCountries  = []plaid.CountryCode{
		plaid.COUNTRYCODE_US,
	}
	PlaidProducts = []plaid.Products{
		plaid.PRODUCTS_TRANSACTIONS,
		plaid.PRODUCTS_INVESTMENTS,
		plaid.PRODUCTS_LIABILITIES,
		plaid.PRODUCTS_BALANCE,
	}
)

func PlaidProductStrings() []string {
	items := make([]string, len(PlaidProducts))
	for i, product := range PlaidProducts {
		items[i] = string(product)
	}

	return items
}
