package domain

import (
	"github.com/Gradient-and-Co/sigma-school/internal/core/errs"
	"io"
)

// Url представляет строку, содержащую URL.
type Url string

// String возвращает строковое представление URL.
func (url Url) String() string {
	return string(url)
}

// File представляет файл с его метаданными.
type File struct {
	// Name - имя файла.
	// Не может быть пустым.
	Name string

	// Path - путь к файлу.
	// Не может быть пустым.
	Path string

	// Reader - объект, который предоставляет доступ к содержимому файла.
	// Не может быть nil.
	Reader io.Reader
}

// Validate проверяет корректность данных файла.
// Возвращает ошибку в случае, если:
// - имя файла пустое (ErrFilenameEmpty),
// - путь к файлу пустой (ErrFilepathEmpty),
// - отсутствует Reader для файла (ErrFileReaderEmpty).
func (f *File) Validate() error {
	if f.Name == "" {
		return errs.ErrFilenameEmpty
	}
	if f.Path == "" {
		return errs.ErrFilepathEmpty
	}
	if f.Reader == nil {
		return errs.ErrFileReaderEmpty
	}
	return nil
}

