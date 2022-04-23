package calculator

type DiscountCalculator struct {
	minimumPurchaseAmount int
	discountAmount        int
}

func NewDiscountCalculator(minimumPurchaseAmount int, discountAmount int) *DiscountCalculator {
	return &DiscountCalculator{
		minimumPurchaseAmount: minimumPurchaseAmount,
		discountAmount:        discountAmount,
	}
}

func (c *DiscountCalculator) Calculate(purchaseAmount int) int {

	if purchaseAmount > c.minimumPurchaseAmount {
		return purchaseAmount - c.discountAmount
	}

	return purchaseAmount
}
