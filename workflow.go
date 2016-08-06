package clubhouse

import "encoding/json"

type Workflow struct {
	CreatedAt      string  `json:"created_at"`
	DefaultStateID int64   `json:"default_state_id"`
	ID             int64   `json:"id"`
	States         []State `json:"states"`
	UpdatedAt      string  `json:"updated_at"`
}

type State struct {
	Color       string `json:"color"`
	CreatedAt   string `json:"created_at"`
	Description string `json:"description"`
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	NumStories  int64  `json:"num_stories"`
	Position    int64  `json:"position"`
	Type        string `json:"type"`
	UpdatedAt   string `json:"updated_at"`
	Verb        string `json:"verb"`
}

func (ch *Clubhouse) ListWorkflow() ([]Workflow, error) {
	body, err := ch.listResources("users")
	if err != nil {
		return []Workflow{}, err
	}
	workflows := []Workflow{}
	json.Unmarshal(body, &workflows)
	return workflows, nil
}
