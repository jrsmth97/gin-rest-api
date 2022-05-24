package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ArticleEntity struct {
	ID        string `gorm:"primaryKey;"`
	Title     string `gorm:"type:varchar(50);not null"`
	Content   string `gorm:"type:text;"`
	AuthorId  string `gorm:"type:int;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (entity *ArticleEntity) BeforeCreate(db *gorm.DB) error {
	entity.ID = uuid.New().String()
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *ArticleEntity) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
