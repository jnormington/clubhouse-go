package clubhouse

import (
	"encoding/json"
	"time"
)

type Team struct {
	CreatedAt   *time.Time `json:"created_at"`
	Updated     *time.Time `json:"updated_at"`
	Description string     `json:"description"`
	EntityType  string     `json:"entity_type"`
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Position    int64      `json:"position"`
	ProjectIDs  []int64    `json:"project_ids"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Workflow    Workflow
	TeamID      int64 `json:"team_id"`
}

func (ch *Clubhouse) ListTeams() ([]Team, error) {
	body, err := ch.listResources("team")
	if err != nil {
		return []Team{}, err
	}
	teams := []Team{}
	json.Unmarshal(body, &teams)
	return teams, nil
}
