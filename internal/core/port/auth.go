package port

import (
	"github.com/Gradient-and-Co/sigma-school/internal/core/domain"
	"github.com/guregu/null"
)

type SignInParam struct {
	Email       string
	Password    string
	Fingerprint string
}

type SignUpParam struct {
	Name      string
	Surname   string
	Email     string
	Password  string
	Phone     null.String
	City      null.String
	AvatarUrl null.String
}

type IAuthProvider interface {
	CreateJWTSession(payload domain.AuthPayload, fingerprint string) (domain.AuthDetails, error)
	RefreshJWTSession(refreshToken domain.Token, fingerprint string) (domain.AuthDetails, error)
	DeleteJWTSession(refreshToken domain.Token) error
	VerifyJWTToken(accessToken domain.Token) (domain.AuthPayload, error)
}
