package util

import (
	"time"

	"github.com/gofrs/uuid"
)

type Util struct {
	now  *time.Time
	uuid *uuid.UUID
}

func NewUtil(now *time.Time, uuid *uuid.UUID) *Util {
	return &Util{
		now:  now,
		uuid: uuid,
	}
}

func (u *Util) GetCurrentTime() time.Time {
	if u.now != nil {
		return *u.now
	}

	return time.Now()
}

func (u *Util) GetNewUUID() uuid.UUID {
	if u.uuid != nil {
		return *u.uuid
	}

	return *u.uuid
}
