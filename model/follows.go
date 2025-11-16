package model

import "time"

type Follows struct {
	ID           uint64 `gorm:"primaryKey;autoIncrement"`
	FollowerUUID string `gorm:"size:36;index"`
	FolloweeUUID string `gorm:"size:36;index"`
	CreatedAt    time.Time
}
