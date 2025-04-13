package port

import (
	"context"
	"github.com/Gradient-and-Co/sigma-school/internal/core/domain"
	"net/url"
)

type IPaymentGateway interface {
	GetPaymentUrl(ctx context.Context, payload domain.PaymentPayload) (url.URL, error)
	ProcessPayment(ctx context.Context, key string) (domain.PaymentPayload, error)
}
