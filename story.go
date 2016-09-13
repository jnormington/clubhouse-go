package clubhouse

import (
	"encoding/json"
	"fmt"
	"time"
)

type CreateStory struct {
	Comments        []CreateComment   `json:"comments"`
	CreatedAt       *time.Time        `json:"created_at,omitempty"`
	Deadline        *time.Time        `json:"deadline,omitempty"`
	Description     string            `json:"description"`
	EpicID          int64             `json:"epic_id,omitempty"`
	Estimate        int64             `json:"estimate"`
	ExternalID      string            `json:"external_id"`
	FileIds         []int64           `json:"file_ids"`
	FollowerIds     []string          `json:"follower_ids"`
	Labels          []CreateLabel     `json:"labels"`
	LinkedFileIds   []int64           `json:"linked_file_ids"`
	Name            string            `json:"name"`
	OwnerIds        []string          `json:"owner_ids"`
	ProjectID       int64             `json:"project_id"`
	RequestedByID   string            `json:"requested_by_id"`
	StoryLinks      []CreateStoryLink `json:"story_links"`
	StoryType       string            `json:"story_type"`
	Tasks           []CreateTask      `json:"tasks"`
	UpdatedAt       *time.Time        `json:"updated_at,omitempty"`
	WorkflowStateID int64             `json:"workflow_state_id"`
}

type Story struct {
	Archived        bool        `json:"archived"`
	Comments        []Comment   `json:"comments"`
	CreatedAt       string      `json:"created_at"`
	Deadline        string      `json:"deadline"`
	Description     string      `json:"description"`
	EpicID          int64       `json:"epic_id"`
	Estimate        int64       `json:"estimate"`
	FileIds         []int64     `json:"file_ids"`
	FollowerIds     []string    `json:"follower_ids"`
	ID              int64       `json:"id"`
	Labels          []Label     `json:"labels"`
	LinkedFileIds   []int64     `json:"linked_file_ids"`
	Name            string      `json:"name"`
	OwnerIds        []string    `json:"owner_ids"`
	Position        int64       `json:"position"`
	ProjectID       int64       `json:"project_id"`
	RequestedByID   string      `json:"requested_by_id"`
	StoryLinks      []StoryLink `json:"story_links"`
	StoryType       string      `json:"story_type"`
	Tasks           []Task      `json:"tasks"`
	UpdatedAt       string      `json:"updated_at"`
	WorkflowStateID int64       `json:"workflow_state_id"`
}

type UpdateStory struct {
	AfterID         int64         `json:"after_id"`
	Archived        bool          `json:"archived"`
	BeforeID        int64         `json:"before_id"`
	Deadline        string        `json:"deadline"`
	Description     string        `json:"description"`
	EpicID          int64         `json:"epic_id"`
	Estimate        int64         `json:"estimate"`
	FileIds         []int64       `json:"file_ids"`
	FollowerIds     []string      `json:"follower_ids"`
	Labels          []CreateLabel `json:"labels"`
	LinkedFileIds   []int64       `json:"linked_file_ids"`
	Name            string        `json:"name"`
	OwnerIds        []string      `json:"owner_ids"`
	ProjectID       int64         `json:"project_id"`
	RequestedByID   string        `json:"requested_by_id"`
	StoryType       string        `json:"story_type"`
	WorkflowStateID int64         `json:"workflow_state_id"`
}

type SearchStory struct {
	Archived           bool      `json:"archived,omitempty"`
	CreatedAtEnd       time.Time `json:"created_at_end,omitempty"`
	CreatedAtStart     time.Time `json:"created_at_start,omitempty"`
	EpicID             int64     `json:"epic_id,omitempty"`
	EpicIds            []int64   `json:"epic_ids,omitempty"`
	Estimate           int64     `json:"estimate,omitempty"`
	LabelName          string    `json:"label_name,omitempty"`
	OwnerID            string    `json:"owner_id,omitempty"`
	OwnerIds           []string  `json:"owner_ids,omitempty"`
	ProjectID          int64     `json:"project_id,omitempty"`
	ProjectIds         []int64   `json:"project_ids,omitempty"`
	RequestedByID      string    `json:"requested_by_id,omitempty"`
	StoryType          string    `json:"story_type,omitempty"`
	Text               string    `json:"text,omitempty"`
	UpdatedAtEnd       time.Time `json:"updated_at_end,omitempty"`
	UpdatedAtStart     time.Time `json:"updated_at_start,omitempty"`
	WorkflowStateID    int64     `json:"workflow_state_id,omitempty"`
	WorkflowStateTypes []string  `json:"workflow_state_types,omitempty"`
}

func (ch *Clubhouse) CreateMultipleStories(newStories []CreateStory) ([]Story, error) {
	jsonStr, _ := json.Marshal(newStories)

	body, err := ch.createObject("stories/bulk", jsonStr)
	if err != nil {
		return []Story{}, err
	}
	stories := []Story{}
	json.Unmarshal(body, &stories)
	return stories, nil
}

func (ch *Clubhouse) UpdateMultipleStories(updatedStories []UpdateStory) ([]Story, error) {
	jsonStr, _ := json.Marshal(updatedStories)

	body, err := ch.updateResource("stories/bulk", jsonStr)
	if err != nil {
		return []Story{}, err
	}
	stories := []Story{}
	json.Unmarshal(body, &stories)
	return stories, nil
}

func (ch *Clubhouse) CreateStory(newStory CreateStory) (Story, error) {
	jsonStr, _ := json.Marshal(newStory)

	body, err := ch.createObject("stories", jsonStr)
	if err != nil {
		return Story{}, err
	}
	story := Story{}
	json.Unmarshal(body, &story)
	return story, nil
}

func (ch *Clubhouse) GetStory(storyID int64) (Story, error) {
	body, err := ch.getResource(fmt.Sprintf("%s/%d", "stories", storyID))
	if err != nil {
		return Story{}, err
	}
	story := Story{}
	json.Unmarshal(body, &story)
	return story, nil
}

func (ch *Clubhouse) UpdateStory(updatedStory UpdateStory, storyID int64) (Story, error) {
	jsonStr, _ := json.Marshal(updatedStory)
	body, err := ch.updateResource(fmt.Sprintf("%s/%d", "stories", storyID), jsonStr)
	if err != nil {
		return Story{}, err
	}
	story := Story{}
	json.Unmarshal(body, &story)
	return story, nil
}

func (ch *Clubhouse) DeleteStory(storyID int64) error {
	return ch.deleteResource(fmt.Sprintf("%s/%d", "stories", storyID))
}
