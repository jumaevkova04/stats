package stats

import (
	"github.com/jumaevkova04/bank/v2/pkg/types"
)

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {

	len := types.Money(len(payments))

	money := types.Money(0)
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			payment.Amount = 0
			len--
		}
		money += payment.Amount
	}
	average := money / len

	return average
}

// TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {

	money := types.Money(0)
	for _, payment := range payments {
		if payment.Category == category {
			if payment.Status == types.StatusFail {
				payment.Amount = 0
			}
			money += payment.Amount
		}
	}
	return money
}
