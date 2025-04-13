package entity

import (
	"github.com/Gradient-and-Co/sigma-school/internal/core/domain"
	"github.com/google/uuid"
)

type PgReview struct {
	ID       uuid.UUID `db:"id"`
	UserID   uuid.UUID `db:"user_id"`
	CourseID uuid.UUID `db:"course_id"`
	Text     string    `db:"text"`
}

func (r *PgReview) ToDomain() domain.Review {
	return domain.Review{
		ID:       domain.ID(r.ID.String()),
		UserID:   domain.ID(r.UserID.String()),
		CourseID: domain.ID(r.CourseID.String()),
		Text:     r.Text,
	}
}

func NewPgReview(review domain.Review) PgReview {
	id, _ := uuid.Parse(review.ID.String())
	userID, _ := uuid.Parse(review.UserID.String())
	courseID, _ := uuid.Parse(review.CourseID.String())
	return PgReview{
		ID:       id,
		UserID:   userID,
		CourseID: courseID,
		Text:     review.Text,
	}
}
