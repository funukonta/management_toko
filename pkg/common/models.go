package common

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func generateNewUUIDWithPrefix(prefix string) string {
	return prefix + "_" + uuid.New().String()
}

type ModelID struct {
	ID         string         `json:"id"`
	Created_at time.Time      `json:"created_at"`
	Updated_at time.Time      `json:"updated_at"`
	Deleted_at gorm.DeletedAt `json:"deleted_at"`
}

func (m *ModelID) BeforeCreate(db *gorm.DB) error {
	m.Created_at = time.Now()
	return nil
}

func (m *ModelID) SetUUID(prefix string) {
	m.ID = generateNewUUIDWithPrefix(prefix)
}
