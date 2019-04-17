package clubhouse

import (
	"encoding/json"
	"fmt"
	"time"
)

type Team struct {
	CreatedAt   *time.Time `json:"created_at"`
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
	body, err := ch.listResources("teams")
	if err != nil {
		return []Team{}, err
	}
	teams := []Team{}
	json.Unmarshal(body, &teams)
	return teams, nil
}

func (ch *Clubhouse) GetTeam(teamID int64) (Team, error) {
	body, err := ch.getResource(fmt.Sprintf("%s/%d", "teams", teamID))
	if err != nil {
		return Team{}, err
	}
	team := Team{}
	json.Unmarshal(body, &team)
	return team, nil
}
