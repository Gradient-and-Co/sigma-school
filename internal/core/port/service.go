package port

import (
	"context"
	"github.com/Gradient-and-Co/sigma-school/internal/core/domain"
)

type IUserService interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	FindByID(ctx context.Context, userID domain.ID) (domain.User, error)
	FindByCredentials(ctx context.Context, credentials UserCredentials) (domain.User, error)
	FindUserInfo(ctx context.Context, userID domain.ID) (UserInfo, error)
	Create(ctx context.Context, param CreateUserParam) (domain.User, error)
	Update(ctx context.Context, userID domain.ID, param UpdateUserParam) (domain.User, error)
	Delete(ctx context.Context, userID domain.ID) error
}

type ICourseService interface {
	FindAll(ctx context.Context) ([]domain.Course, error)
	FindByID(ctx context.Context, courseID domain.ID) (domain.Course, error)
	FindStudentCourses(ctx context.Context, studentID domain.ID) ([]domain.Course, error)
	FindTeacherCourses(ctx context.Context, teacherID domain.ID) ([]domain.Course, error)
	FindCourseTeachers(ctx context.Context, courseID domain.ID) ([]domain.User, error)
	IsCourseStudent(ctx context.Context, studentID, courseID domain.ID) (bool, error)
	IsCourseTeacher(ctx context.Context, teacherID, courseID domain.ID) (bool, error)
	AddCourseStudent(ctx context.Context, studentID, courseID domain.ID) error
	AddCourseTeacher(ctx context.Context, teacherID, courseID domain.ID) error
	ConfirmDraftCourse(ctx context.Context, courseID domain.ID) []error
	PublishReadyCourse(ctx context.Context, courseID domain.ID) error
	CreateSchoolCourse(ctx context.Context, schoolID domain.ID,
		param CreateCourseParam) (domain.Course, error)
	Update(ctx context.Context, courseID domain.ID,
		param UpdateCourseParam) (domain.Course, error)
	Delete(ctx context.Context, courseID domain.ID) error
}

type ILessonService interface {
	FindAll(ctx context.Context) ([]domain.Lesson, error)
	FindByID(ctx context.Context, lessonID domain.ID) (domain.Lesson, error)
	FindCourseLessons(ctx context.Context, courseID domain.ID) ([]domain.Lesson, error)
	CreateTheoryLesson(ctx context.Context, courseID domain.ID,
		param CreateTheoryParam) (domain.Lesson, error)
	CreateVideoLesson(ctx context.Context, courseID domain.ID,
		param CreateVideoParam) (domain.Lesson, error)
	CreatePracticeLesson(ctx context.Context, courseID domain.ID,
		param CreatePracticeParam) (domain.Lesson, error)
	UpdateTheoryLesson(ctx context.Context, lessonID domain.ID,
		param UpdateTheoryParam) (domain.Lesson, error)
	UpdateVideoLesson(ctx context.Context, lessonID domain.ID,
		param UpdateVideoParam) (domain.Lesson, error)
	UpdatePracticeLesson(ctx context.Context, lessonID domain.ID,
		param UpdatePracticeParam) (domain.Lesson, error)
	Delete(ctx context.Context, lessonID domain.ID) error
}

type IStatService interface {
	FindLessonStat(ctx context.Context, userID, lessonID domain.ID) (domain.LessonStat, error)
	CreateLessonStat(ctx context.Context, userID, lessonID domain.ID) error
	UpdateLessonStat(ctx context.Context, userID, lessonID domain.ID,
		param UpdateLessonStatParam) error
}
