package biling

import (
	"diploma/internal/domain"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

const (
	maskCreateCustomer = iota
	maskPurchase
	maskPayout
	maskRecurring
	maskFraudControl
	maskCheckoutPage
)

func ParseData(path string) *domain.BillingData {
	r, err := os.Open(path)
	if err != nil {
		log.Println(err)

		return nil
	}
	reader, err := io.ReadAll(r)
	if err != nil {
		log.Println(err)

		return nil
	}
	var num float64
	RightByte := 0
	for i := 5; i >= 0; i-- {
		if string(reader[i]) == "1" {
			num += math.Pow(2, float64(RightByte))
		}
		RightByte++
	}
	n := uint8(num)
	fmt.Println(n)
	return &domain.BillingData{
		CreateCustomer: n&(1<<maskCreateCustomer) > 0,
		Purchase:       n&(1<<maskPurchase) > 0,
		Payout:         n&(1<<maskPayout) > 0,
		Recurring:      n&(1<<maskRecurring) > 0,
		FraudControl:   n&(1<<maskFraudControl) > 0,
		CheckoutPage:   n&(1<<maskCheckoutPage) > 0,
	}
}
