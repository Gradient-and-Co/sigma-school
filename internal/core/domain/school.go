package domain

// School представляет учебное заведение (школу).
type School struct {
	// ID уникальный идентификатор школы.
	ID ID
	// OwnerID уникальный идентификатор владельца школы.
	OwnerID ID
	// Name название школы.
	Name string
	// Description описание школы.
	Description string
}

