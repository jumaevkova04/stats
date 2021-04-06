package stats

import (
	"fmt"

	"github.com/jumaevkova04/bank/pkg/types"
)

// ExampleAvg ...
func ExampleAvg() {
	payments := []types.Payment{{
		ID:       1,
		Amount:   10_000_000,
		Category: "авто",
	},
		{
			ID:       2,
			Amount:   10_000,
			Category: "рестораны",
		},
		{
			ID:       3,
			Amount:   10_000,
			Category: "рестораны",
		},
	}

	average := Avg(payments)
	fmt.Println(average)
	result := TotalInCategory(payments, "рестораны")
	fmt.Println(result)

	// Output:
	// 3340000
	// 20000
}
