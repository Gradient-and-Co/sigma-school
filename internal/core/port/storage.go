package port

import (
	"context"
	"github.com/Gradient-and-Co/sigma-school/internal/core/domain"
)

type IObjectStorage interface {
	SaveFile(ctx context.Context, file domain.File) (domain.Url, error)
}
