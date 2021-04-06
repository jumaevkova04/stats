package stats

import "github.com/jumaevkova04/bank/pkg/types"

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {

	money := types.Money(0)
	for _, payment := range payments {
		money += payment.Amount
	}
	average := money / types.Money(len(payments))

	return average
}

// TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {

	money := types.Money(0)
	for _, payment := range payments {
		if payment.Category == category {
			money += payment.Amount
		}
	}
	return money
}
