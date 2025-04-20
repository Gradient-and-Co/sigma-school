package port

import (
	"context"
	"github.com/Gradient-and-Co/sigma-school/internal/core/domain"
)

// IUserRepository описывает репозиторий для работы с пользователями.
type IUserRepository interface {
	// FindAll возвращает список всех пользователей в системе.
	// Если запрос не удался, возвращается ошибка.
	FindAll(ctx context.Context) ([]domain.User, error)

	// FindByID находит пользователя по его уникальному идентификатору.
	// Если пользователь не найден, возвращается ошибка.
	FindByID(ctx context.Context, userID domain.ID) (domain.User, error)

	// FindByEmail находит пользователя по его email адресу.
	// Если пользователь с таким email не найден, возвращается ошибка.
	FindByEmail(ctx context.Context, email string) (domain.User, error)

	// FindByCredentials находит пользователя по email и паролю.
	// Если данные неверны, возвращается ошибка.
	FindByCredentials(ctx context.Context, email string, password string) (domain.User, error)

	// FindUserInfo получает дополнительную информацию о пользователе по его ID.
	// Возвращается структура UserInfo с дополнительной информацией о пользователе.
	FindUserInfo(ctx context.Context, userID domain.ID) (UserInfo, error)

	// Create создает нового пользователя в системе.
	// Возвращается созданный пользователь.
	Create(ctx context.Context, user domain.User) (domain.User, error)

	// Update обновляет информацию о пользователе.
	// Возвращается обновленный пользователь.
	Update(ctx context.Context, user domain.User) (domain.User, error)

	// Delete удаляет пользователя по его ID.
	// Возвращается ошибка, если удаление не удалось.
	Delete(ctx context.Context, userID domain.ID) error
}

// ICourseRepository описывает репозиторий для работы с курсами.
type ICourseRepository interface {
	// FindAll возвращает список всех курсов в системе.
	// Если запрос не удался, возвращается ошибка.
	FindAll(ctx context.Context) ([]domain.Course, error)

	// FindByID находит курс по его уникальному идентификатору.
	// Если курс не найден, возвращается ошибка.
	FindByID(ctx context.Context, courseID domain.ID) (domain.Course, error)

	// FindStudentCourses возвращает список курсов, на которых обучается студент.
	// Если запрос не удался, возвращается ошибка.
	FindStudentCourses(ctx context.Context, studentID domain.ID) ([]domain.Course, error)

	// FindTeacherCourses возвращает список курсов, которые преподает преподаватель.
	// Если запрос не удался, возвращается ошибка.
	FindTeacherCourses(ctx context.Context, teacherID domain.ID) ([]domain.Course, error)

	// FindCourseTeachers находит преподавателей, которые ведут курс.
	// Если запрос не удался, возвращается ошибка.
	FindCourseTeachers(ctx context.Context, courseID domain.ID) ([]domain.User, error)

	// IsCourseStudent проверяет, является ли студент участником курса.
	// Возвращает true, если студент является участником курса, и ошибку, если запрос не удался.
	IsCourseStudent(ctx context.Context, studentID, courseID domain.ID) (bool, error)

	// IsCourseTeacher проверяет, является ли преподаватель преподавателем курса.
	// Возвращает true, если преподаватель ведет курс, и ошибку, если запрос не удался.
	IsCourseTeacher(ctx context.Context, teacherID, courseID domain.ID) (bool, error)

	// AddCourseStudent добавляет студента в курс.
	// Возвращается ошибка, если операция не удалась.
	AddCourseStudent(ctx context.Context, studentID, courseID domain.ID) error

	// AddCourseTeacher добавляет преподавателя в курс.
	// Возвращается ошибка, если операция не удалась.
	AddCourseTeacher(ctx context.Context, teacherID, courseID domain.ID) error

	// Create создает новый курс.
	// Возвращается созданный курс.
	Create(ctx context.Context, course domain.Course) (domain.Course, error)

	// Update обновляет информацию о курсе.
	// Возвращается обновленный курс.
	Update(ctx context.Context, course domain.Course) (domain.Course, error)

	// UpdateStatus обновляет статус курса.
	// Возвращается ошибка, если запрос не удался.
	UpdateStatus(ctx context.Context, courseID domain.ID, status domain.CourseStatus) error

	// Delete удаляет курс по его ID.
	// Возвращается ошибка, если удаление не удалось.
	Delete(ctx context.Context, courseID domain.ID) error
}

// ILessonRepository описывает репозиторий для работы с уроками.
type ILessonRepository interface {
	// FindAll возвращает список всех уроков в системе.
	// Если запрос не удался, возвращается ошибка.
	FindAll(ctx context.Context) ([]domain.Lesson, error)

	// FindByID находит урок по его уникальному идентификатору.
	// Если урок не найден, возвращается ошибка.
	FindByID(ctx context.Context, lessonID domain.ID) (domain.Lesson, error)

	// FindCourseLessons возвращает список уроков для указанного курса.
	// Если запрос не удался, возвращается ошибка.
	FindCourseLessons(ctx context.Context, courseID domain.ID) ([]domain.Lesson, error)

	// FindLessonTests находит тесты, связанные с уроком.
	// Если запрос не удался, возвращается ошибка.
	FindLessonTests(ctx context.Context, lessonID domain.ID) ([]domain.Test, error)

	// Create создает новый урок.
	// Возвращается созданный урок.
	Create(ctx context.Context, lesson domain.Lesson) (domain.Lesson, error)

	// Update обновляет информацию о уроке.
	// Возвращается обновленный урок.
	Update(ctx context.Context, lesson domain.Lesson) (domain.Lesson, error)

	// Delete удаляет урок по его ID.
	// Возвращается ошибка, если удаление не удалось.
	Delete(ctx context.Context, lessonID domain.ID) error
}

// ISchoolRepository описывает репозиторий для работы с учебными заведениями.
type ISchoolRepository interface {
	// FindAll возвращает список всех школ.
	// Если запрос не удался, возвращается ошибка.
	FindAll(ctx context.Context) ([]domain.School, error)

	// FindByID находит школу по ее уникальному идентификатору.
	// Если школа не найдена, возвращается ошибка.
	FindByID(ctx context.Context, schoolID domain.ID) (domain.School, error)

	// FindUserSchools возвращает список школ, в которых состоит пользователь.
	// Если запрос не удался, возвращается ошибка.
	FindUserSchools(ctx context.Context, userID domain.ID) ([]domain.School, error)

	// FindSchoolCourses возвращает список курсов, доступных в школе.
	// Если запрос не удался, возвращается ошибка.
	FindSchoolCourses(ctx context.Context, schoolID domain.ID) ([]domain.Course, error)

	// FindSchoolTeachers возвращает список преподавателей, работающих в школе.
	// Если запрос не удался, возвращается ошибка.
	FindSchoolTeachers(ctx context.Context, schoolID domain.ID) ([]domain.User, error)

	// IsSchoolTeacher проверяет, является ли преподаватель преподавателем данной школы.
	// Возвращает true, если преподаватель работает в школе, и ошибку, если запрос не удался.
	IsSchoolTeacher(ctx context.Context, schoolID, teacherID domain.ID) (bool, error)

	// AddSchoolTeacher добавляет преподавателя в школу.
	// Возвращается ошибка, если добавление не удалось.
	AddSchoolTeacher(ctx context.Context, schoolID, teacherID domain.ID) error

	// Create создает новую школу.
	// Возвращается созданная школа.
	Create(ctx context.Context, school domain.School) (domain.School, error)

	// Update обновляет информацию о школе.
	// Возвращается обновленная школа.
	Update(ctx context.Context, school domain.School) (domain.School, error)

	// Delete удаляет школу по ее ID.
	// Возвращается ошибка, если удаление не удалось.
	Delete(ctx context.Context, schoolID domain.ID) error
}

// IStatRepository описывает репозиторий для работы со статистикой уроков.
type IStatRepository interface {
	// FindLessonStat находит статистику для урока.
	// Возвращается статистика по уроку, если она найдена.
	FindLessonStat(ctx context.Context, userID, lessonID domain.ID) (domain.LessonStat, error)

	// CreateLessonStat создает статистику для урока.
	// Возвращается ошибка, если создание не удалось.
	CreateLessonStat(ctx context.Context, stat domain.LessonStat) error

	// UpdateLessonStat обновляет статистику для урока.
	// Возвращается ошибка, если обновление не удалось.
	UpdateLessonStat(ctx context.Context, stat domain.LessonStat) error
}

// IReviewRepository описывает репозиторий для работы с отзывами.
type IReviewRepository interface {
	// FindAll возвращает список всех отзывов.
	// Если запрос не удался, возвращается ошибка.
	FindAll(ctx context.Context) ([]domain.Review, error)

	// FindByID находит отзыв по его уникальному идентификатору.
	// Если отзыв не найден, возвращается ошибка.
	FindByID(ctx context.Context, reviewID domain.ID) (domain.Review, error)

	// FindUserReviews возвращает отзывы, оставленные пользователем.
	// Если запрос не удался, возвращается ошибка.
	FindUserReviews(ctx context.Context, userID domain.ID) ([]domain.Review, error)

	// FindCourseReviews возвращает отзывы, оставленные на курс.
	// Если запрос не удался, возвращается ошибка.
	FindCourseReviews(ctx context.Context, courseID domain.ID) ([]domain.Review, error)

	// Create создает новый отзыв.
	// Возвращается созданный отзыв.
	Create(ctx context.Context, review domain.Review) (domain.Review, error)

	// Delete удаляет отзыв по его ID.
	// Возвращается ошибка, если удаление не удалось.
	Delete(ctx context.Context, reviewID domain.ID) error
}

