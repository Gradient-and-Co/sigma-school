package errs

import "errors"

var (
	ErrCourseLessonInvalidScore             = errors.New("course lesson score must be > 0")
	ErrCoursePracticeLessonEmptyTestTaskUrl = errors.New("course practice lesson test has no question")
	ErrCoursePracticeLessonEmptyTestOptions = errors.New("course practice lesson test has no options")
	ErrCoursePracticeLessonInvalidTestScore = errors.New("course practice lesson test score must be > 0")
	ErrCoursePracticeLessonInvalidTestLevel = errors.New("course practice lesson test level must be > 0")
)

var (
	ErrFilenameEmpty   = errors.New("validation filename is empty error")
	ErrFilepathEmpty   = errors.New("validation filepath is empty error")
	ErrFileReaderEmpty = errors.New("validation file reader is nil error")
)
