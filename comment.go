package clubhouse

import "time"

type ThreadedComment struct {
	AuthorID  string            `json:"author_id"`
	CreatedAt time.Time         `json:"created_at"`
	Deleted   bool              `json:"deleted"`
	ID        int64             `json:"id"`
	Text      string            `json:"text"`
	UpdatedAt time.Time         `json:"updated_at"`
	Comments  []ThreadedComment `json:"comments"`
}
