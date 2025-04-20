// Package errs содержит общие ошибки, используемые в приложении.
//
// Пакет определяет стандартные ошибки, которые могут быть использованы для обработки различных ситуаций,
// таких как проблемы с курсами, пользователями, платежами, записями в базе данных, файлами, а также ошибки аутентификации.
//
// Каждая ошибка описана с использованием стандартного пакета `errors` и предоставляет информативные сообщения,
// которые могут быть использованы для диагностики и обработки ошибок на различных уровнях приложения.
//
// Ошибки в этом пакете разделены на несколько категорий:
// 1. Ошибки, связанные с курсами.
// 2. Ошибки, связанные с пользователями и их действиями.
// 3. Ошибки, связанные с операциями записи в базу данных.
// 4. Ошибки, связанные с работой с файлами.
// 5. Ошибки, связанные с аутентификацией и авторизацией.
//
// Ошибки, определенные в пакете:
// - ErrCourseNotEnoughLessons: Ошибка, возникающая, если в курсе недостаточно уроков.
// - ErrCourseLessonInvalidScore: Ошибка, если балл за урок некорректен.
// - ErrCoursePracticeLessonEmptyTests: Ошибка, если в практическом уроке нет тестов.
// - ErrCourseInvalidLevel: Ошибка, если уровень курса некорректен.
// - ErrInvalidCredentials: Ошибка при неверных учетных данных пользователя.
package errs

import "errors"

var (
	// ErrCourseNotEnoughLessons возникает, если в курсе недостаточно уроков.
	// Курс должен содержать минимум один теоретический и один практический урок.
	ErrCourseNotEnoughLessons = errors.New("course must have at least 1 theory and 1 practice lessons")

	// ErrCourseLessonInvalidScore возникает, если балл за урок некорректен.
	// Балл должен быть больше нуля.
	ErrCourseLessonInvalidScore = errors.New("course lesson score must be > 0")

	// ErrCoursePracticeLessonEmptyTests возникает, если в практическом уроке нет тестов.
	ErrCoursePracticeLessonEmptyTests = errors.New("course practice lesson must contain at least 1 test")

	// ErrCoursePracticeLessonEmptyTestTaskUrl возникает, если у теста нет вопроса.
	ErrCoursePracticeLessonEmptyTestTaskUrl = errors.New("course practice lesson test has no question")

	// ErrCoursePracticeLessonEmptyTestOptions возникает, если у теста нет вариантов ответа.
	ErrCoursePracticeLessonEmptyTestOptions = errors.New("course practice lesson test has no options")

	// ErrCoursePracticeLessonInvalidTestScore возникает, если балл за тест некорректен.
	// Балл должен быть больше нуля.
	ErrCoursePracticeLessonInvalidTestScore = errors.New("course practice lesson test score must be > 0")

	// ErrCoursePracticeLessonInvalidTestLevel возникает, если уровень сложности теста некорректен.
	// Уровень должен быть больше нуля.
	ErrCoursePracticeLessonInvalidTestLevel = errors.New("course practice lesson test level must be > 0")

	// ErrCourseTheoryLessonEmptyUrl возникает, если у теоретического урока отсутствует URL.
	ErrCourseTheoryLessonEmptyUrl = errors.New("course theory lesson url is empty")

	// ErrCourseVideoLessonEmptyUrl возникает, если у видеоурока отсутствует URL.
	ErrCourseVideoLessonEmptyUrl = errors.New("course video lesson url is empty")

	// ErrCourseReadyState возникает, если курс не находится в состоянии "черновик" при попытке перевести его в "готов".
	ErrCourseReadyState = errors.New("course must be in draft state to make it ready")

	// ErrCoursePublishedState возникает, если курс не находится в состоянии "готов" при попытке опубликовать его.
	ErrCoursePublishedState = errors.New("course must be in ready state to publish it")

	// ErrCourseInvalidLevel возникает, если уровень курса некорректен.
	// Уровень должен быть больше нуля.
	ErrCourseInvalidLevel = errors.New("course level must be > 0")

	// ErrCourseInvalidPrice возникает, если цена курса некорректна.
	// Цена должна быть неотрицательной.
	ErrCourseInvalidPrice = errors.New("course price must be >= 0")
)

var (
	// ErrUserIsNotSchoolTeacher возникает, если пользователь не является преподавателем данной школы.
	ErrUserIsNotSchoolTeacher = errors.New("user is not a teacher in this school")

	// ErrUserIsAlreadyCourseStudent возникает, если пользователь уже записан на данный курс.
	ErrUserIsAlreadyCourseStudent = errors.New("user has already bought this course")

	// ErrInvalidPaymentSum возникает, если сумма оплаты некорректна.
	ErrInvalidPaymentSum = errors.New("received invalid payment")

	// ErrDecodePaymentKeyFailed возникает при ошибке декодирования данных платежа.
	ErrDecodePaymentKeyFailed = errors.New("failed to decode payment payload")
)

var (
	// ErrDuplicate возникает при попытке создать уже существующую запись.
	ErrDuplicate = errors.New("record already exists")

	// ErrNotExist возникает при попытке получить несуществующую запись.
	ErrNotExist = errors.New("record does not exist")

	// ErrUpdateFailed возникает при ошибке обновления записи.
	ErrUpdateFailed = errors.New("record update failed")

	// ErrDeleteFailed возникает при ошибке удаления записи.
	ErrDeleteFailed = errors.New("record delete failed")

	// ErrPersistenceFailed возникает при внутренней ошибке слоя хранения данных.
	ErrPersistenceFailed = errors.New("persistence internal error")

	// ErrEnumValueError возникает, если значение перечисления выходит за допустимые пределы.
	ErrEnumValueError = errors.New("enum value is out of scope")

	// ErrTransactionError возникает при ошибке транзакции.
	ErrTransactionError = errors.New("transaction error occurred")
)

var (
	// ErrFilenameEmpty возникает, если имя файла не указано.
	ErrFilenameEmpty = errors.New("validation filename is empty error")

	// ErrFilepathEmpty возникает, если путь к файлу не указан.
	ErrFilepathEmpty = errors.New("validation filepath is empty error")

	// ErrFileReaderEmpty возникает, если reader файла отсутствует.
	ErrFileReaderEmpty = errors.New("validation file reader is nil error")

	// ErrSaveFileError возникает при ошибке сохранения файла в объектное хранилище.
	ErrSaveFileError = errors.New("failed to save file to object storage")
)

var (
	// ErrInvalidCredentials возникает при неверных учетных данных.
	ErrInvalidCredentials = errors.New("invalid credentials")

	// ErrNotUniqueEmail возникает, если пользователь с таким email уже существует.
	ErrNotUniqueEmail = errors.New("user with such email already exists")

	// ErrAuthSessionIsNotPresent возникает, если сессия по указанному refresh-токену не найдена.
	ErrAuthSessionIsNotPresent = errors.New("session with such refresh token is not present")

	// ErrInvalidTokenSignMethod возникает, если метод подписи JWT некорректен.
	ErrInvalidTokenSignMethod = errors.New("invalid signing method")

	// ErrInvalidTokenClaims возникает, если claims JWT-токена недействительны.
	ErrInvalidTokenClaims = errors.New("invalid token claims")

	// ErrInvalidToken возникает, если JWT-токен недействителен.
	ErrInvalidToken = errors.New("invalid jwt token")

	// ErrInvalidFingerprint возникает, если отпечаток клиента недействителен.
	ErrInvalidFingerprint = errors.New("invalid client fingerprint")
)

