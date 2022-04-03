package delivery

import (
	"time"

	"github.com/google/uuid"
)

type (
	PermissionDto struct {
		ID               uuid.UUID
		Name             string
		CreationDate     time.Time
		LastUpdateDate   time.Time
		MarkAsDeleteDate *time.Time
	}
)
