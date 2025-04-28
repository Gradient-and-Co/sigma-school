package port

import (
	"github.com/Gradient-and-Co/sigma-school/internal/core/domain"
	"github.com/guregu/null"
)

// SignInParam представляет параметры, необходимые для аутентификации пользователя.
//
// Используется для входа пользователя в онлайн-школу.
type SignInParam struct {
	Email       string // Электронная почта пользователя
	Password    string // Пароль пользователя
	Fingerprint string // Уникальный отпечаток устройства для идентификации сессии
}

// SignUpParam представляет параметры, необходимые для регистрации нового пользователя.
//
// Используется при создании аккаунта в онлайн-школе.
type SignUpParam struct {
	Name      string      // Имя пользователя
	Surname   string      // Фамилия пользователя
	Email     string      // Электронная почта пользователя
	Password  string      // Пароль для входа
	Phone     null.String // Номер телефона (необязательный)
	City      null.String // Город проживания (необязательный)
	AvatarUrl null.String // URL изображения аватара пользователя (необязательный)
}

// IAuthProvider определяет интерфейс для управления аутентификацией и сессиями пользователей.
//
// Интерфейс включает в себя создание и обновление JWT-сессий, а также верификацию и удаление токенов.
// Используется для обеспечения безопасного доступа к функционалу онлайн-платформы.
type IAuthProvider interface {

	// CreateJWTSession создает новую JWT-сессию для пользователя.
	//
	// Принимает информацию об аутентификации и отпечаток устройства.
	// Возвращает детали аутентификации, включая access и refresh токены.
	CreateJWTSession(payload domain.AuthPayload, fingerprint string) (domain.AuthDetails, error)

	// RefreshJWTSession обновляет существующую JWT-сессию.
	//
	// Принимает refresh токен и отпечаток устройства.
	// Возвращает новые детали аутентификации.
	RefreshJWTSession(refreshToken domain.Token, fingerprint string) (domain.AuthDetails, error)

	// DeleteJWTSession удаляет сессию по refresh токену.
	//
	// Используется при выходе пользователя из системы.
	DeleteJWTSession(refreshToken domain.Token) error

	// VerifyJWTToken проверяет корректность и актуальность access токена.
	//
	// Возвращает полезную нагрузку токена при успешной проверке.
	VerifyJWTToken(accessToken domain.Token) (domain.AuthPayload, error)
}

