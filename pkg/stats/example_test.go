package stats

import (
	"fmt"

	"github.com/jumaevkova04/bank/v2/pkg/types"
)

// ExampleAvg ...
func ExampleAvg() {
	payments := []types.Payment{types.Payment{
		ID:       1,
		Amount:   10_000_000,
		Category: "авто",
		Status:   "OK",
	},
		types.Payment{
			ID:       2,
			Amount:   10_000,
			Category: "рестораны",
			Status:   types.StatusFail,
		},
		types.Payment{
			ID:       3,
			Amount:   10_000,
			Category: "рестораны",
			Status:   "OK",
		},
	}

	average := Avg(payments)
	fmt.Println(average)
	result := TotalInCategory(payments, "рестораны")
	fmt.Println(result)

	// Output:
	// 5005000
	// 10000
}
