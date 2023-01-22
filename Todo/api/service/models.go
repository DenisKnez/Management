package service

import "github.com/gofrs/uuid"

type Todo struct {
	ID   uuid.UUID
	Text string
}
