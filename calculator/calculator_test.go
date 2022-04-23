package calculator

import "testing"

// TestDiscountCalculator - table driven testing approach
func TestDiscountCalculator(t *testing.T) {
	type testCase struct {
		minimumPurchaseAmount int
		purchaseAmount        int
		discount              int
		expectedAmount        int
	}

	testCases := []testCase{
		{
			minimumPurchaseAmount: 100,
			purchaseAmount:        150,
			discount:              20,
			expectedAmount:        130,
		},
		{
			minimumPurchaseAmount: 100,
			purchaseAmount:        200,
			discount:              20,
			expectedAmount:        180,
		},
		{
			minimumPurchaseAmount: 100,
			purchaseAmount:        350,
			discount:              20,
			expectedAmount:        330,
		},
		{
			minimumPurchaseAmount: 100,
			purchaseAmount:        50,
			discount:              20,
			expectedAmount:        50,
		},
		{
			minimumPurchaseAmount: 100,
			purchaseAmount:        50,
			discount:              0,
			expectedAmount:        50,
		},
	}

	for _, tc := range testCases {
		calculator := NewDiscountCalculator(tc.minimumPurchaseAmount, tc.discount)
		amount := calculator.Calculate(tc.purchaseAmount)

		if amount != tc.expectedAmount {
			t.Errorf("expected %v, got %v", tc.expectedAmount, amount)
		}
	}
}
