package domain

// Token представляет собой строковый тип, который используется для хранения аутентификационных токенов.
type Token string

// String возвращает строковое представление токена.
func (t Token) String() string {
    return string(t)
}

// AuthDetails представляет данные аутентификации пользователя, включая access и refresh токены.
type AuthDetails struct {
    // AccessToken — токен доступа, который используется для аутентификации пользователя в запросах.
    AccessToken Token

    // RefreshToken — токен обновления, который используется для получения нового access токена.
    RefreshToken Token
}

// AuthPayload представляет полезную нагрузку для аутентификационных запросов.
// Включает ID пользователя, который авторизован в системе.
type AuthPayload struct {
    // UserID — уникальный идентификатор пользователя, для которого генерируются токены.
    UserID ID
}
