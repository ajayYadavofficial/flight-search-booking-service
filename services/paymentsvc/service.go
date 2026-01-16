package paymentsvc

import (
	"context"
	"crypto/rand"
)

type Payment interface {
	ProcessPayment(ctx context.Context, transactionId string, userName string) (isSuccessful bool)
}

type demoPayment struct{}

func NewPaymentService() Payment {
	return &demoPayment{}
}

func (p *demoPayment) ProcessPayment(ctx context.Context, transactionId string, userName string) (isSuccessful bool) {
	// 80% success rate for demo purposes
	// generate a random number between 1 and 10

	randBytes := make([]byte, 1)
	_, err := rand.Read(randBytes)
	if err != nil {
		return false
	}
	randomNum := int(randBytes[0]%10) + 1
	if randomNum <= 8 {
		return true
	}

	return false
}
