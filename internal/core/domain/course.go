package domain

// CourseStatus представляет возможные статусы курса.
type CourseStatus int

const (
	// CourseDraft - статус "Черновик", когда курс еще не готов.
	CourseDraft CourseStatus = iota
	// CourseReady - статус "Готов", когда курс готов к публикации.
	CourseReady
	// CoursePublished - статус "Опубликован", когда курс доступен для студентов.
	CoursePublished
)

// Course представляет курс в системе.
type Course struct {
	// ID - уникальный идентификатор курса.
	ID ID

	// SchoolID - идентификатор учебного заведения, к которому относится курс.
	SchoolID ID

	// Name - название курса.
	Name string

	// Description - описание курса.
	Description string

	// Level - уровень сложности курса.
	// Значение должно быть больше нуля.
	Level int

	// Price - цена курса в виде целого числа (в минимальных единицах валюты, например, в копейках).
	Price int64

	// Language - язык, на котором преподается курс.
	Language string

	// ImageUrl - URL изображения для курса.
	ImageUrl string

	// Status - статус курса.
	// Может быть одним из: CourseDraft, CourseReady, CoursePublished.
	Status CourseStatus
}

