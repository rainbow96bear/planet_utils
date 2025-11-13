package model

import "time"

// Todo represents a single task in a calendar event
type Todo struct {
	ID      uint64 `gorm:"primaryKey;column:id" json:"id"`
	EventID uint64 `gorm:"column:event_id;not null" json:"eventId"` // FK to Calendar.EventID
	Content string `gorm:"type:varchar(255);not null" json:"content"`
	Done    bool   `gorm:"column:done;default:false" json:"done"`
}

// Calendar represents a calendar event
type Calendar struct {
	EventID     uint64    `gorm:"primaryKey;column:eventId" json:"eventId"` // DB PK는 그대로 유지
	UserUUID    string    `gorm:"column:user_id;type:char(36);not null" json:"userUuid"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Emoji       string    `gorm:"type:varchar(10);default:'📝'" json:"emoji"`
	StartAt     time.Time `gorm:"column:start_at;type:date;not null" json:"startAt"`
	EndAt       time.Time `gorm:"column:end_at;type:date;not null" json:"endAt"`
	Visibility  string    `gorm:"type:enum('public','friends','private');default:'public'" json:"visibility"`
	ImageURL    *string   `gorm:"column:image_url;type:text" json:"imageUrl,omitempty"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`

	Todos []Todo `gorm:"foreignKey:EventID;references:EventID" json:"todos"` // FK 매핑
}
