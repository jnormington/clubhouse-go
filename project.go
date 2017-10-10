package clubhouse

import (
	"encoding/json"
	"fmt"
	"time"
)

type CreateProject struct {
	Abbreviation string    `json:"abbreviation"`
	Color        string    `json:"color"`
	CreatedAt    time.Time `json:"created_at"`
	Description  string    `json:"description"`
	ExternalID   string    `json:"external_id"`
	FollowerIds  []string  `json:"follower_ids"`
	Name         string    `json:"name"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Project struct {
	Abbreviation string    `json:"abbreviation"`
	Archived     bool      `json:"archived"`
	Color        string    `json:"color"`
	CreatedAt    time.Time `json:"created_at"`
	Description  string    `json:"description"`
	FollowerIds  []string  `json:"follower_ids"`
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	NumPoints    int64     `json:"num_points"`
	NumStories   int64     `json:"num_stories"`
	UpdatedAt    time.Time `json:"updated_at"`
	TeamID       int64     `json:"team_id"`
}

type UpdateProject struct {
	Abbreviation string   `json:"abbreviation"`
	Archived     bool     `json:"archived"`
	Color        string   `json:"color"`
	Description  string   `json:"description"`
	FollowerIds  []string `json:"follower_ids"`
	Name         string   `json:"name"`
}

func (ch *Clubhouse) ListProjects() ([]Project, error) {
	body, err := ch.listResources("projects")
	if err != nil {
		return []Project{}, err
	}
	projects := []Project{}
	json.Unmarshal(body, &projects)
	return projects, nil
}

func (ch *Clubhouse) CreateProject(newProject CreateProject) (Project, error) {
	jsonStr, _ := json.Marshal(newProject)

	body, err := ch.createObject("projects", jsonStr)
	if err != nil {
		return Project{}, err
	}
	project := Project{}
	json.Unmarshal(body, &project)
	return project, nil
}

func (ch *Clubhouse) ListStories(projectID int64) ([]Story, error) {
	body, err := ch.listResources(fmt.Sprintf("%s/%d/%s", "projects", projectID, "stories"))
	if err != nil {
		return []Story{}, err
	}
	stories := []Story{}
	json.Unmarshal(body, &stories)
	return stories, nil
}

func (ch *Clubhouse) GetProject(projectID int64) (Project, error) {
	body, err := ch.getResource(fmt.Sprintf("%s/%d", "projects", projectID))
	if err != nil {
		return Project{}, err
	}
	project := Project{}
	json.Unmarshal(body, &project)
	return project, nil
}

func (ch *Clubhouse) UpdateProject(updatedProject UpdateProject, projectID int64) (Project, error) {
	jsonStr, _ := json.Marshal(updatedProject)
	body, err := ch.updateResource(fmt.Sprintf("%s/%d", "projects", projectID), jsonStr)
	if err != nil {
		return Project{}, err
	}
	project := Project{}
	json.Unmarshal(body, &project)
	return project, nil
}

func (ch *Clubhouse) DeleteProject(projectID int64) error {
	return ch.deleteResource(fmt.Sprintf("%s/%d", "projects", projectID))
}
