package domain

import (
	"github.com/Gradient-and-Co/sigma-school/internal/core/errs"
	"github.com/guregu/null"
)

// LessonType представляет тип урока (теоретический, практический или видеоурок).
type LessonType int

const (
	// TheoryLesson — теоретический урок.
	TheoryLesson LessonType = iota
	// PracticeLesson — практический урок.
	PracticeLesson
	// VideoLesson — видеоурок.
	VideoLesson
)

// Lesson представляет урок в курсе, который может быть теоретическим, практическим или видеоуроком.
type Lesson struct {
	// ID уникальный идентификатор урока.
	ID ID
	// CourseID идентификатор курса, к которому относится урок.
	CourseID ID
	// Title название урока.
	Title string
	// Score оценка, которую можно получить за урок.
	Score int
	// Type тип урока (теоретический, практический, видеоурок).
	Type LessonType
	// TheoryUrl URL теоретического урока (если применимо).
	TheoryUrl null.String
	// VideoUrl URL видеоурока (если применимо).
	VideoUrl null.String
	// Tests список тестов, связанных с практическим уроком.
	Tests []Test
}

// Test представляет тест, связанный с уроком.
type Test struct {
	// ID уникальный идентификатор теста.
	ID ID
	// LessonID идентификатор урока, к которому привязан тест.
	LessonID ID
	// TaskUrl URL задания теста.
	TaskUrl string
	// Options варианты ответов на тест.
	Options []string
	// Answer правильный ответ.
	Answer string
	// Level уровень сложности теста.
	Level int
	// Score баллы, которые можно получить за тест.
	Score int
}

// Validate проверяет корректность данных урока.
func (l *Lesson) Validate() error {
	// Проверка, что оценка урока больше 0.
	if l.Score <= 0 {
		return errs.ErrCourseLessonInvalidScore
	}
	// Валидация для различных типов уроков.
	switch l.Type {
	case PracticeLesson:
		// Для практического урока проверяются тесты.
		for _, test := range l.Tests {
			// Проверка наличия задания.
			if test.TaskUrl == "" {
				return errs.ErrCoursePracticeLessonEmptyTestTaskUrl
			}
			// Проверка наличия вариантов ответов.
			if len(test.Options) == 0 {
				return errs.ErrCoursePracticeLessonEmptyTestOptions
			}
			// Проверка уровня сложности теста.
			if test.Level < 0 {
				return errs.ErrCoursePracticeLessonInvalidTestLevel
			}
			// Проверка баллов за тест.
			if test.Score <= 0 {
				return errs.ErrCoursePracticeLessonInvalidTestScore
			}
		}
	case TheoryLesson:
		// Теоретический урок не требует дополнительных проверок.
	case VideoLesson:
		// Видеоурок также не требует дополнительных проверок.
	}

	return nil
}

