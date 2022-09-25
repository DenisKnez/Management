package repository

import (
	"time"

	"github.com/gofrs/uuid"
)

type TodoEntity struct {
	ID        uuid.UUID
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   bool
}
