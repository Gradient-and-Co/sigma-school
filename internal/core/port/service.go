package port

import (
	"context"
	"github.com/Gradient-and-Co/sigma-school/internal/core/domain"
)

type IUserService interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	FindByID(ctx context.Context, userID domain.ID) (domain.User, error)
	FindByCredentials(ctx context.Context, credentials UserCredentials) (domain.User, error)
	FindUserInfo(ctx context.Context, userID domain.ID) (UserInfo, error)
	Create(ctx context.Context, param CreateUserParam) (domain.User, error)
	Update(ctx context.Context, userID domain.ID, param UpdateUserParam) (domain.User, error)
	Delete(ctx context.Context, userID domain.ID) error
}
