package paymentsvc

import "context"

type Payment interface {
	ProcessPayment(ctx context.Context, transactionId string, userName string) (isSuccessful bool)
}
