package clubhouse

import (
	"encoding/json"
	"fmt"
	"time"
)

type CreateTask struct {
	Complete    bool      `json:"complete"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	ExternalID  string    `json:"external_id"`
	OwnerIds    []string  `json:"owner_ids"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Task struct {
	Complete    bool      `json:"complete"`
	CompletedAt time.Time `json:"completed_at"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	ID          int64     `json:"id"`
	MentionIds  []string  `json:"mention_ids"`
	OwnerIds    []string  `json:"owner_ids"`
	Position    int64     `json:"position"`
	StoryID     int64     `json:"story_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateTask struct {
	AfterID     int64    `json:"after_id"`
	BeforeID    int64    `json:"before_id"`
	Complete    bool     `json:"complete"`
	Description string   `json:"description"`
	OwnerIds    []string `json:"owner_ids"`
}

func (ch *Clubhouse) CreateTask(newTask CreateTask, storyID int64) (Task, error) {
	jsonStr, _ := json.Marshal(newTask)

	body, err := ch.createObject(fmt.Sprintf("%s/%d/%s", "stories", storyID, "tasks"), jsonStr)
	if err != nil {
		return Task{}, err
	}
	task := Task{}
	json.Unmarshal(body, &task)
	return task, nil
}

func (ch *Clubhouse) GetTask(storyID int64, taskID int64) (Task, error) {
	body, err := ch.getResource(fmt.Sprintf("%s/%d/%s/%d", "stories", storyID, "tasks", taskID))
	if err != nil {
		return Task{}, err
	}
	task := Task{}
	json.Unmarshal(body, &task)
	return task, nil
}

func (ch *Clubhouse) UpdateTask(updatedTask UpdateTask, storyID int64, taskID int64) (Task, error) {
	jsonStr, _ := json.Marshal(updatedTask)
	body, err := ch.updateResource(fmt.Sprintf("%s/%d/%s/%d", "stories", storyID, "tasks", taskID), jsonStr)
	if err != nil {
		return Task{}, err
	}
	task := Task{}
	json.Unmarshal(body, &task)
	return task, nil
}

func (ch *Clubhouse) DeleteTask(storyID int64, taskID int64) error {
	return ch.deleteResource(fmt.Sprintf("%s/%d/%s/%d", "stories", storyID, "tasks", taskID))
}
