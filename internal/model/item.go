package model

import (
	"time"
)

type Item struct {
	Id         uint64    `gorm:"column:id" json:"id"`
	AuthorId   uint64    `gorm:"column:author_id" json:"author_id"`
	Title      string    `gorm:"column:title" json:"title"`
	Type       uint      `gorm:"column:type" json:"type"`
	CoverUrl   string    `gorm:"column:cover_url" json:"cover_url"`
	VideoUrl   string    `gorm:"column:video_url" json:"video_url"`
	PictureUrl string    `gorm:"column:picture_url" json:"picture_url"`
	Context    string    `gorm:"column:context" json:"context"`
	Status     uint      `gorm:"column:status" json:"status"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (m *Item) TableName() string {
	return "timeline.item"
}
