package domain

import "github.com/google/uuid"

// ID представляет уникальный идентификатор в виде строки.
type ID string

// String возвращает строковое представление идентификатора.
func (id ID) String() string {
	return string(id)
}

// NewID генерирует новый уникальный идентификатор с использованием UUID.
func NewID() ID {
	return ID(uuid.New().String())
}

// RandomID генерирует новый случайный идентификатор.
// Это псевдоним для функции NewID.
func RandomID() ID {
	return NewID()
}

// ParseID парсит строковое значение в тип ID.
// Возвращает ошибку, если строка не является допустимым UUID.
func ParseID(value string) (ID, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return ID(uuid.UUID{}.String()), err
	}
	return ID(id.String()), nil
}

