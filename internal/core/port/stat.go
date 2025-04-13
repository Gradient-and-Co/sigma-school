package port

import (
	"github.com/Gradient-and-Co/sigma-school/internal/core/domain"
	"github.com/guregu/null"
)

type UpdateLessonStatParam struct {
	Score     null.Int
	TestStats []UpdateTestStatParam
}

type UpdateTestStatParam struct {
	TestID domain.ID
	Score  int
}
