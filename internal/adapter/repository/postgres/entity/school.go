package entity

import (
	"github.com/Gradient-and-Co/sigma-school/internal/core/domain"
	"github.com/google/uuid"
)

type PgSchool struct {
	ID          uuid.UUID `db:"id"`
	OwnerID     uuid.UUID `db:"owner_id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
}

func (s *PgSchool) ToDomain() domain.School {
	return domain.School{
		ID:          domain.ID(s.ID.String()),
		OwnerID:     domain.ID(s.OwnerID.String()),
		Name:        s.Name,
		Description: s.Description,
	}
}

func NewPgSchool(school domain.School) PgSchool {
	id, _ := uuid.Parse(school.ID.String())
	ownerID, _ := uuid.Parse(school.OwnerID.String())
	return PgSchool{
		ID:          id,
		OwnerID:     ownerID,
		Name:        school.Name,
		Description: school.Description,
	}
}
