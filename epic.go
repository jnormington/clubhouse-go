package clubhouse

import (
	"encoding/json"
	"time"
)

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
	Archived    bool              `json:"archived"`
	Comments    []ThreadedComment `json:"comments"`
	CreatedAt   time.Time         `json:"created_at"`
	Deadline    time.Time         `json:"deadline"`
	Description string            `json:"description"`
	FollowerIds []string          `json:"follower_ids"`
	ID          int64             `json:"id"`
	Name        string            `json:"name"`
	OwnerIds    []string          `json:"owner_ids"`
	Position    int64             `json:"position"`
	State       string            `json:"state"`
	UpdatedAt   time.Time         `json:"updated_at"`
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

func (ch *Clubhouse) GetEpic(epicID int64) (Epic, error) {
	body, err := ch.getResource("epics", epicID)
	if err != nil {
		return Epic{}, err
	}
	epic := Epic{}
	json.Unmarshal(body, &epic)
	return epic, nil
}

func (ch *Clubhouse) UpdateEpic(updatedEpic UpdateEpic, epicID int64) (Epic, error) {
	jsonStr, _ := json.Marshal(updatedEpic)
	body, err := ch.updateResource("epics", epicID, jsonStr)
	if err != nil {
		return Epic{}, err
	}
	epic := Epic{}
	json.Unmarshal(body, &epic)
	return epic, nil
}

func (ch *Clubhouse) DeleteEpic(epicID int64) error {
	return ch.deleteResource("epics", epicID)
}

func (ch *Clubhouse) ListEpics() ([]Epic, error) {
	body, err := ch.listResources("epics")
	if err != nil {
		return []Epic{}, err
	}
	epics := []Epic{}
	json.Unmarshal(body, &epics)
	return epics, nil
}

func (ch *Clubhouse) CreateEpic(newEpic CreateEpic) (Epic, error) {
	jsonStr, _ := json.Marshal(newEpic)
	body, err := ch.createObject("epics", jsonStr)
	if err != nil {
		return Epic{}, err
	}
	epic := Epic{}
	json.Unmarshal(body, &epic)
	return epic, nil
}
