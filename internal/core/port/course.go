package port

import "github.com/guregu/null"

// CourseInfo представляет информацию о курсе.
type CourseInfo struct {
	// Name — название курса.
	Name string

	// Level — уровень сложности курса.
	Level int

	// Price — цена курса в выбранной валюте.
	Price int64

	// Language — язык, на котором проводится курс.
	Language string
}

// CreateCourseParam содержит параметры для создания нового курса.
type CreateCourseParam struct {
	// Name — название курса.
	Name string

	// Description — описание курса.
	Description string

	// Level — уровень сложности курса.
	Level int

	// Price — цена курса в выбранной валюте.
	Price int64

	// Language — язык курса.
	Language string

	// ImageUrl — URL изображения, ассоциированного с курсом.
	ImageUrl string
}

// UpdateCourseParam содержит параметры для обновления информации о курсе.
// Поля могут быть пустыми (null), если они не требуются для обновления.
type UpdateCourseParam struct {
	// Name — новое название курса, если требуется обновление.
	// Может быть пустым (null), если имя не изменяется.
	Name null.String

	// Description — новое описание курса, если требуется обновление.
	// Может быть пустым (null), если описание не изменяется.
	Description null.String

	// Level — новый уровень сложности курса, если требуется обновление.
	// Может быть пустым (null), если уровень не изменяется.
	Level null.Int

	// Price — новая цена курса, если требуется обновление.
	// Может быть пустым (null), если цена не изменяется.
	Price null.Int

	// Language — новый язык курса, если требуется обновление.
	// Может быть пустым (null), если язык не изменяется.
	Language null.String

	// ImageUrl — новый URL изображения для курса, если требуется обновление.
	// Может быть пустым (null), если изображение не изменяется.
	ImageUrl null.String
}
