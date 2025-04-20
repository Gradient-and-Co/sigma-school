package port

import (
	"context"
	"github.com/Gradient-and-Co/sigma-school/internal/core/domain"
	"net/url"
)

// IPaymentGateway описывает интерфейс для работы с платежным шлюзом.
type IPaymentGateway interface {
	// GetPaymentUrl генерирует URL для оплаты на основании переданных данных.
	// Возвращает сформированный URL для платежа и ошибку, если операция не удалась.
	GetPaymentUrl(ctx context.Context, payload domain.PaymentPayload) (url.URL, error)

	// ProcessPayment обрабатывает платеж по предоставленному ключу.
	// Возвращает объект PaymentPayload, содержащий информацию о платеже, и ошибку, если обработка не удалась.
	ProcessPayment(ctx context.Context, key string) (domain.PaymentPayload, error)
}

