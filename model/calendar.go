package model

import "time"

type Todo struct {
	ID      uint64 `gorm:"primaryKey;column:id" json:"id"`
	EventID uint64 `gorm:"column:event_id;not null" json:"event_id"` // calendar_events.id 참조
	Content string `gorm:"type:varchar(255);not null" json:"content"`
	Done    bool   `gorm:"column:done;default:false" json:"done"`
}

type Calendar struct {
	ID          uint64    `gorm:"primaryKey;column:id" json:"id"` // 실제 DB 컬럼명
	UserUUID    string    `gorm:"column:user_id;type:char(36);not null" json:"user_uuid"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Emoji       string    `gorm:"type:varchar(10);default:'📝'" json:"emoji"`
	StartAt     time.Time `gorm:"column:start_at;type:date;not null" json:"start_at"`
	EndAt       time.Time `gorm:"column:end_at;type:date;not null" json:"end_at"`
	Visibility  string    `gorm:"type:enum('public','friends','private');default:'public'" json:"visibility"`
	ImageURL    *string   `gorm:"column:image_url;type:text" json:"image_url,omitempty"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	Todos []Todo `gorm:"foreignKey:EventID;references:ID" json:"todos"` // event_id FK
}
