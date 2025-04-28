// Package port содержит интерфейсы, описывающие операции с репозиториями данных для работы с основными сущностями системы,
// такими как пользователи, курсы, уроки, школы, статистика и отзывы. Эти интерфейсы служат для взаимодействия с хранилищем данных,
// предоставляя абстракцию для выполнения CRUD операций (создание, чтение, обновление, удаление) с различными типами данных.
//
// Пакет включает в себя репозитории для следующих сущностей:
// - Пользователи (IUserRepository)
// - Курсы (ICourseRepository)
// - Уроки (ILessonRepository)
// - Школы (ISchoolRepository)
// - Статистика (IStatRepository)
// - Отзывы (IReviewRepository)
//
// Каждый репозиторий определяет методы для работы с соответствующими сущностями, а также методы для выполнения специфических операций,
// таких как проверка роли пользователя в курсе или школы, или создание статистики для прохождения урока.
//
// Основные компоненты пакета:
// - Интерфейсы репозиториев, которые описывают операции, доступные для взаимодействия с данными.
// - Методы для получения информации, добавления, обновления и удаления данных для каждого типа сущности.
// - Описание возможных ошибок, которые могут возникнуть при выполнении операций.
//
// Пакет служит для инкапсуляции логики работы с данными и позволяет легко заменять конкретные реализации репозиториев,
// не затрагивая логику бизнес-слоя.
package port

import (
	"context"
	"github.com/Gradient-and-Co/sigma-school/internal/core/domain"
	"io"
	"net/url"
)

// Пакет port описывает интерфейсы и методы для работы с различными сервисами в системе, включая управление пользователями, курсами, уроками, статистикой, школами, отзывами, оплатами и медиафайлами.
// Эти интерфейсы определяют операции, которые могут быть выполнены в рамках каждого сервиса, включая создание, обновление, удаление и поиск данных, а также обработку различных бизнес-операций.


// IUserService предоставляет методы для управления пользователями системы.
type IUserService interface {
	// FindAll возвращает список всех пользователей.
	FindAll(ctx context.Context) ([]domain.User, error)

	// FindByID возвращает пользователя по его идентификатору.
	FindByID(ctx context.Context, userID domain.ID) (domain.User, error)

	// FindByCredentials возвращает пользователя по учетным данным (email/пароль).
	FindByCredentials(ctx context.Context, credentials UserCredentials) (domain.User, error)

	// FindUserInfo возвращает агрегированную информацию о пользователе.
	FindUserInfo(ctx context.Context, userID domain.ID) (UserInfo, error)

	// Create создает нового пользователя на основе переданных параметров.
	Create(ctx context.Context, param CreateUserParam) (domain.User, error)

	// Update обновляет данные существующего пользователя.
	Update(ctx context.Context, userID domain.ID, param UpdateUserParam) (domain.User, error)

	// Delete удаляет пользователя по его идентификатору.
	Delete(ctx context.Context, userID domain.ID) error
}

// ICourseService предоставляет методы для управления курсами.
type ICourseService interface {
	// FindAll возвращает список всех курсов.
	FindAll(ctx context.Context) ([]domain.Course, error)

	// FindByID возвращает курс по его идентификатору.
	FindByID(ctx context.Context, courseID domain.ID) (domain.Course, error)

	// FindStudentCourses возвращает курсы, в которых обучается студент.
	FindStudentCourses(ctx context.Context, studentID domain.ID) ([]domain.Course, error)

	// FindTeacherCourses возвращает курсы, которые ведет преподаватель.
	FindTeacherCourses(ctx context.Context, teacherID domain.ID) ([]domain.Course, error)

	// FindCourseTeachers возвращает список преподавателей курса.
	FindCourseTeachers(ctx context.Context, courseID domain.ID) ([]domain.User, error)

	// IsCourseStudent проверяет, является ли пользователь студентом курса.
	IsCourseStudent(ctx context.Context, studentID, courseID domain.ID) (bool, error)

	// IsCourseTeacher проверяет, является ли пользователь преподавателем курса.
	IsCourseTeacher(ctx context.Context, teacherID, courseID domain.ID) (bool, error)

	// AddCourseStudent добавляет студента в курс.
	AddCourseStudent(ctx context.Context, studentID, courseID domain.ID) error

	// AddCourseTeacher добавляет преподавателя в курс.
	AddCourseTeacher(ctx context.Context, teacherID, courseID domain.ID) error

	// ConfirmDraftCourse проверяет черновой курс на наличие ошибок перед публикацией.
	ConfirmDraftCourse(ctx context.Context, courseID domain.ID) []error

	// PublishReadyCourse публикует курс, прошедший модерацию.
	PublishReadyCourse(ctx context.Context, courseID domain.ID) error

	// CreateSchoolCourse создает новый курс внутри указанной школы.
	CreateSchoolCourse(ctx context.Context, schoolID domain.ID, param CreateCourseParam) (domain.Course, error)

	// Update обновляет курс по его идентификатору.
	Update(ctx context.Context, courseID domain.ID, param UpdateCourseParam) (domain.Course, error)

	// Delete удаляет курс по его идентификатору.
	Delete(ctx context.Context, courseID domain.ID) error
}

// ILessonService предоставляет методы для работы с уроками.
type ILessonService interface {
	// FindAll возвращает список всех уроков.
	FindAll(ctx context.Context) ([]domain.Lesson, error)

	// FindByID возвращает урок по идентификатору.
	FindByID(ctx context.Context, lessonID domain.ID) (domain.Lesson, error)

	// FindCourseLessons возвращает список уроков курса.
	FindCourseLessons(ctx context.Context, courseID domain.ID) ([]domain.Lesson, error)

	// CreateTheoryLesson создает теоретический урок в курсе.
	CreateTheoryLesson(ctx context.Context, courseID domain.ID, param CreateTheoryParam) (domain.Lesson, error)

	// CreateVideoLesson создает видеоурок в курсе.
	CreateVideoLesson(ctx context.Context, courseID domain.ID, param CreateVideoParam) (domain.Lesson, error)

	// CreatePracticeLesson создает практический урок в курсе.
	CreatePracticeLesson(ctx context.Context, courseID domain.ID, param CreatePracticeParam) (domain.Lesson, error)

	// UpdateTheoryLesson обновляет теоретический урок.
	UpdateTheoryLesson(ctx context.Context, lessonID domain.ID, param UpdateTheoryParam) (domain.Lesson, error)

	// UpdateVideoLesson обновляет видеоурок.
	UpdateVideoLesson(ctx context.Context, lessonID domain.ID, param UpdateVideoParam) (domain.Lesson, error)

	// UpdatePracticeLesson обновляет практический урок.
	UpdatePracticeLesson(ctx context.Context, lessonID domain.ID, param UpdatePracticeParam) (domain.Lesson, error)

	// Delete удаляет урок по его идентификатору.
	Delete(ctx context.Context, lessonID domain.ID) error
}

// IStatService предоставляет методы для управления статистикой по урокам.
type IStatService interface {
	// FindLessonStat возвращает статистику прохождения урока пользователем.
	FindLessonStat(ctx context.Context, userID, lessonID domain.ID) (domain.LessonStat, error)

	// CreateLessonStat создает новую запись о прохождении урока.
	CreateLessonStat(ctx context.Context, userID, lessonID domain.ID) error

	// UpdateLessonStat обновляет данные о прохождении урока.
	UpdateLessonStat(ctx context.Context, userID, lessonID domain.ID, param UpdateLessonStatParam) error
}

// ISchoolService предоставляет методы для работы со школами.
type ISchoolService interface {
	// FindAll возвращает список всех школ.
	FindAll(ctx context.Context) ([]domain.School, error)

	// FindByID возвращает школу по ее идентификатору.
	FindByID(ctx context.Context, schoolID domain.ID) (domain.School, error)

	// FindUserSchools возвращает список школ, в которых состоит пользователь.
	FindUserSchools(ctx context.Context, userID domain.ID) ([]domain.School, error)

	// FindSchoolCourses возвращает курсы, относящиеся к школе.
	FindSchoolCourses(ctx context.Context, schoolID domain.ID) ([]domain.Course, error)

	// FindSchoolTeachers возвращает преподавателей, работающих в школе.
	FindSchoolTeachers(ctx context.Context, schoolID domain.ID) ([]domain.User, error)

	// IsSchoolTeacher проверяет, является ли пользователь преподавателем школы.
	IsSchoolTeacher(ctx context.Context, schoolID, teacherID domain.ID) (bool, error)

	// AddSchoolTeacher добавляет преподавателя в школу.
	AddSchoolTeacher(ctx context.Context, schoolID, teacherID domain.ID) error

	// CreateUserSchool создает новую школу от имени пользователя.
	CreateUserSchool(ctx context.Context, userID domain.ID, param CreateSchoolParam) (domain.School, error)

	// Update обновляет информацию о школе.
	Update(ctx context.Context, schoolID domain.ID, param UpdateSchoolParam) (domain.School, error)

	// Delete удаляет школу.
	Delete(ctx context.Context, schoolID domain.ID) error
} 

// IReviewService предоставляет методы для управления отзывами.
type IReviewService interface {
	// FindAll возвращает все отзывы.
	FindAll(ctx context.Context) ([]domain.Review, error)

	// FindByID возвращает отзыв по его идентификатору.
	FindByID(ctx context.Context, reviewID domain.ID) (domain.Review, error)

	// FindUserReviews возвращает отзывы, оставленные пользователем.
	FindUserReviews(ctx context.Context, userID domain.ID) ([]domain.Review, error)

	// FindCourseReviews возвращает отзывы, относящиеся к курсу.
	FindCourseReviews(ctx context.Context, courseID domain.ID) ([]domain.Review, error)

	// CreateCourseReview создает отзыв к курсу от имени пользователя.
	CreateCourseReview(ctx context.Context, courseID, userID domain.ID, param CreateReviewParam) (domain.Review, error)

	// Delete удаляет отзыв по идентификатору.
	Delete(ctx context.Context, reviewID domain.ID) error
}

// IPaymentService предоставляет методы для работы с оплатами курсов.
type IPaymentService interface {
	// GetCoursePaymentUrl генерирует URL для оплаты курса пользователем.
	GetCoursePaymentUrl(ctx context.Context, userID, courseID domain.ID) (url.URL, error)

	// ProcessCoursePayment обрабатывает поступивший платеж по метке.
	ProcessCoursePayment(ctx context.Context, label string, paid int64) (domain.PaymentPayload, error)
}

// IMediaService предоставляет методы для загрузки и хранения медиафайлов.
type IMediaService interface {
	// SaveMediaFile сохраняет произвольный медиафайл.
	SaveMediaFile(ctx context.Context, file domain.File) (domain.Url, error)

	// SaveUserAvatar сохраняет аватар пользователя.
	SaveUserAvatar(ctx context.Context, userID domain.ID, file domain.File) (domain.Url, error)

	// SaveLessonTheory сохраняет теоретический материал урока в формате PDF или HTML.
	SaveLessonTheory(ctx context.Context, lessonID domain.ID, reader io.Reader) (domain.Url, error)

	// SaveTestQuestion сохраняет изображение к тестовому вопросу.
	SaveTestQuestion(ctx context.Context, testID domain.ID, reader io.Reader) (domain.Url, error)
}

// IAuthTokenService предоставляет методы для работы с JWT токенами.
type IAuthTokenService interface {
	// SignIn выполняет вход пользователя в систему.
	SignIn(ctx context.Context, param SignInParam) (domain.AuthDetails, error)

	// SignUp регистрирует нового пользователя.
	SignUp(ctx context.Context, param SignUpParam) error

	// LogOut выполняет выход пользователя, аннулируя refresh токен.
	LogOut(ctx context.Context, refreshToken domain.Token) error

	// Refresh обновляет пару токенов по refresh токену и отпечатку браузера.
	Refresh(ctx context.Context, refreshToken domain.Token, fingerprint string) (domain.AuthDetails, error)

	// Verify проверяет валидность access токена.
	Verify(ctx context.Context, accessToken domain.Token) error

	// Payload извлекает полезную нагрузку из access токена.
	Payload(ctx context.Context, accessToken domain.Token) (domain.AuthPayload, error)
}

