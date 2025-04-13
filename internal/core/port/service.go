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
