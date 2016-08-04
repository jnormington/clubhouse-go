package clubhouse

import "time"

type CreateLabel struct {
	ExternalID string `json:"external_id"`
	Name       string `json:"name"`
}

type LabelWithCounts struct {
	CreatedAt            time.Time `json:"created_at"`
	ID                   int64     `json:"id"`
	Name                 string    `json:"name"`
	NumStoriesCompleted  int64     `json:"num_stories_completed"`
	NumStoriesInProgress int64     `json:"num_stories_in_progress"`
	NumStoriesTotal      int64     `json:"num_stories_total"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type UpdateLabel struct {
	Name string `json:"name"`
}
