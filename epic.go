package clubhouse

import "time"

type CreateEpic struct {
	CreatedAt   time.Time `json:"created_at"`
	Deadline    time.Time `json:"deadline"`
	Description string    `json:"description"`
	ExternalID  string    `json:"external_id"`
	FollowerIds []string  `json:"follower_ids"`
	Name        string    `json:"name"`
	OwnerIds    []string  `json:"owner_ids"`
	State       string    `json:"state"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Epic struct {
	Archived    bool      `json:"archived"`
	Comments    []Comment `json:"comments"`
	CreatedAt   time.Time `json:"created_at"`
	Deadline    time.Time `json:"deadline"`
	Description string    `json:"description"`
	FollowerIds []string  `json:"follower_ids"`
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	OwnerIds    []string  `json:"owner_ids"`
	Position    int64     `json:"position"`
	State       string    `json:"state"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateEpic struct {
	AfterID     int64     `json:"after_id"`
	Archived    bool      `json:"archived"`
	BeforeID    int64     `json:"before_id"`
	Deadline    time.Time `json:"deadline"`
	Description string    `json:"description"`
	FollowerIds []string  `json:"follower_ids"`
	Name        string    `json:"name"`
	OwnerIds    []string  `json:"owner_ids"`
	State       string    `json:"state"`
}
