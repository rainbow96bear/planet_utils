package model

import "time"

// Event : 사용자의 일정 정보 구조체
type Event struct {
	ID          uint64    `json:"id" db:"id"`                         // AUTO_INCREMENT 기본키
	UserUuid    string    `json:"user_uuid" db:"user_uuid"`           // User UUID 참조 (FK 없이)
	Title       string    `json:"title" db:"title"`                   // 일정 제목
	Description string    `json:"description" db:"description"`       // 일정 설명
	Emoji       string    `json:"emoji" db:"emoji"`                   // 이모지 (기본값: 📝)
	StartDate   time.Time `json:"start_date" db:"start_date"`         // 시작일
	EndDate     time.Time `json:"end_date" db:"end_date"`             // 종료일
	Visibility  string    `json:"visibility" db:"visibility"`         // public/friends/private (기본값: public)
	ImageUrl    string    `json:"image_url,omitempty" db:"image_url"` // 첨부 이미지 URL (선택)
	CreatedAt   time.Time `json:"created_at" db:"created_at"`         // 생성일
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`         // 수정일
}

// EventTodo : 일정에 포함된 할 일 목록
type EventTodo struct {
	ID        uint64    `json:"id" db:"id"`                 // AUTO_INCREMENT 기본키
	EventID   uint64    `json:"event_id" db:"event_id"`     // Event ID 참조
	TodoText  string    `json:"todo_text" db:"todo_text"`   // 할 일 내용
	Completed bool      `json:"completed" db:"completed"`   // 완료 여부 (기본값: false)
	OrderNum  int       `json:"order_num" db:"order_num"`   // 정렬 순서
	CreatedAt time.Time `json:"created_at" db:"created_at"` // 생성일
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"` // 수정일
}

// EventLike : 일정 좋아요
type EventLike struct {
	ID        uint64    `json:"id" db:"id"`                 // AUTO_INCREMENT 기본키
	EventID   uint64    `json:"event_id" db:"event_id"`     // Event ID 참조
	UserUuid  string    `json:"user_uuid" db:"user_uuid"`   // User UUID 참조
	CreatedAt time.Time `json:"created_at" db:"created_at"` // 생성일
}

// EventComment : 일정 댓글
type EventComment struct {
	ID        uint64    `json:"id" db:"id"`                 // AUTO_INCREMENT 기본키
	EventID   uint64    `json:"event_id" db:"event_id"`     // Event ID 참조
	UserUuid  string    `json:"user_uuid" db:"user_uuid"`   // 작성자 UUID
	Content   string    `json:"content" db:"content"`       // 댓글 내용
	CreatedAt time.Time `json:"created_at" db:"created_at"` // 생성일
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"` // 수정일
}

// EventBookmark : 일정 북마크
type EventBookmark struct {
	ID        uint64    `json:"id" db:"id"`                 // AUTO_INCREMENT 기본키
	EventID   uint64    `json:"event_id" db:"event_id"`     // Event ID 참조
	UserUuid  string    `json:"user_uuid" db:"user_uuid"`   // User UUID 참조
	CreatedAt time.Time `json:"created_at" db:"created_at"` // 생성일
}

// EventView : 일정 조회 수 (선택)
type EventView struct {
	ID       uint64    `json:"id" db:"id"`                         // AUTO_INCREMENT 기본키
	EventID  uint64    `json:"event_id" db:"event_id"`             // Event ID 참조
	UserUuid string    `json:"user_uuid,omitempty" db:"user_uuid"` // 조회한 User UUID (선택)
	ViewedAt time.Time `json:"viewed_at" db:"viewed_at"`           // 조회 시간
}

// EventWithDetails : Event + 관련 정보 (조회용)
type EventWithDetails struct {
	Event
	Todos        []EventTodo `json:"todos"`         // 할 일 목록
	LikeCount    int         `json:"like_count"`    // 좋아요 수
	CommentCount int         `json:"comment_count"` // 댓글 수
	IsLiked      bool        `json:"is_liked"`      // 현재 사용자의 좋아요 여부
	IsBookmarked bool        `json:"is_bookmarked"` // 현재 사용자의 북마크 여부
	ViewCount    int         `json:"view_count"`    // 조회 수
	AuthorName   string      `json:"author_name"`   // 작성자 이름
	AuthorHandle string      `json:"author_handle"` // 작성자 핸들
	AuthorAvatar string      `json:"author_avatar"` // 작성자 아바타
}

// EventFeed : 피드용 간단한 Event 정보
type EventFeed struct {
	ID           uint64      `json:"id"`
	UserUuid     string      `json:"user_uuid"`
	Title        string      `json:"title"`
	Emoji        string      `json:"emoji"`
	StartDate    time.Time   `json:"start_date"`
	EndDate      time.Time   `json:"end_date"`
	Visibility   string      `json:"visibility"`
	ImageUrl     string      `json:"image_url,omitempty"`
	Todos        []EventTodo `json:"todos"`
	LikeCount    int         `json:"like_count"`
	CommentCount int         `json:"comment_count"`
	IsLiked      bool        `json:"is_liked"`
	CreatedAt    time.Time   `json:"created_at"`
	AuthorName   string      `json:"author_name"`
	AuthorHandle string      `json:"author_handle"`
	AuthorAvatar string      `json:"author_avatar"`
}
