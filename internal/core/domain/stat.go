package domain

// LessonStat представляет статистику по уроку для пользователя.
type LessonStat struct {
	// ID уникальный идентификатор статистики по уроку.
	ID ID
	// LessonID уникальный идентификатор урока.
	LessonID ID
	// UserID уникальный идентификатор пользователя.
	UserID ID
	// Score балл, полученный пользователем за урок.
	Score int
	// TestStats список статистики по тестам, связанным с уроком.
	TestStats []TestStat
}

// TestStat представляет статистику по тесту для пользователя.
type TestStat struct {
	// ID уникальный идентификатор статистики по тесту.
	ID ID
	// TestID уникальный идентификатор теста.
	TestID ID
	// UserID уникальный идентификатор пользователя.
	UserID ID
	// Score балл, полученный пользователем за тест.
	Score int
}

