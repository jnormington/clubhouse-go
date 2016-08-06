package clubhouse

import (
	"encoding/json"
	"fmt"
	"time"
)

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

func (ch *Clubhouse) ListLabels() ([]LabelWithCounts, error) {
	body, err := ch.listResources("labels")
	if err != nil {
		return []LabelWithCounts{}, err
	}
	labels := []LabelWithCounts{}
	json.Unmarshal(body, &labels)
	return labels, nil
}

func (ch *Clubhouse) CreateLabel(newLabel CreateLabel) (LabelWithCounts, error) {
	jsonStr, _ := json.Marshal(newLabel)

	body, err := ch.createObject("labels", jsonStr)
	if err != nil {
		return LabelWithCounts{}, err
	}
	label := LabelWithCounts{}
	json.Unmarshal(body, &label)
	return label, nil
}

func (ch *Clubhouse) UpdateLabel(updatedLabel UpdateLabel, labelID int64) (LabelWithCounts, error) {
	jsonStr, _ := json.Marshal(updatedLabel)
	body, err := ch.updateResource(fmt.Sprintf("%s/%d", "labels", labelID), jsonStr)
	if err != nil {
		return LabelWithCounts{}, err
	}
	label := LabelWithCounts{}
	json.Unmarshal(body, &label)
	return label, nil
}

func (ch *Clubhouse) DeleteLabel(labelID int64) error {
	return ch.deleteResource(fmt.Sprintf("%s/%d", "labels", labelID))
}
