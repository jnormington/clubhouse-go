package clubhouse

import (
	"encoding/json"
	"fmt"
	"time"
)

//CreateEpic is the object passed to Clubhouse API to create an epic.
//Required fields are:
// CreateEpic.Name
type CreateEpic struct {
	CreatedAt   time.Time `json:"created_at,omitempty"`
	Deadline    time.Time `json:"deadline,omitempty"`
	Description string    `json:"description,omitempty"`
	ExternalID  string    `json:"external_id,omitempty"`
	FollowerIds []string  `json:"follower_ids,omitempty"`
	Name        string    `json:"name"`
	OwnerIds    []string  `json:"owner_ids,omitempty"`
	State       string    `json:"state,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

//An Epic is a collection of stories that together might make up a release, a milestone, or some other large initiative that your organization is working on.
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

// UpdateEpic the body used for ch.EpicUpdate()
type UpdateEpic struct {
	AfterID     int64     `json:"after_id,omitempty"`
	Archived    bool      `json:"archived,omitempty"`
	BeforeID    int64     `json:"before_id,omitempty"`
	Deadline    time.Time `json:"deadline,omitempty"`
	Description string    `json:"description,omitempty"`
	FollowerIds []string  `json:"follower_ids,omitempty"`
	Name        string    `json:"name,omitempty"`
	OwnerIds    []string  `json:"owner_ids,omitempty"`
	State       string    `json:"state,omitempty"`
}

// EpicGet returns information about the selected Epic.
//calls GET https://api.clubhouse.io/api/v1/epics/{epicID} to retrieve the specified epicID
func (ch *Clubhouse) EpicGet(epicID int64) (Epic, error) {
	body, err := ch.getResource(fmt.Sprintf("%s/%d", "epics", epicID))
	if err != nil {
		return Epic{}, err
	}
	epic := Epic{}
	json.Unmarshal(body, &epic)
	return epic, nil
}

// EpicUpdate can be used to update numerous fields in the Epic. The only required parameter is Epic ID, which can be found in the Clubhouse UI.
//calls PUT https://api.clubhouse.io/api/v1/epics/{epicID} and updates it with the data in the UpdateEpic object.
func (ch *Clubhouse) EpicUpdate(updatedEpic UpdateEpic, epicID int64) (Epic, error) {
	jsonStr, _ := json.Marshal(updatedEpic)
	body, err := ch.updateResource(fmt.Sprintf("%s/%d", "epics", epicID), jsonStr)
	if err != nil {
		return Epic{}, err
	}
	epic := Epic{}
	json.Unmarshal(body, &epic)
	return epic, nil
}

// EpicDelete can be used to delete the Epic. The only required parameter is Epic ID.
//Calls DELETE https://api.clubhouse.io/api/v1/epics/{epicID}
func (ch *Clubhouse) EpicDelete(epicID int64) error {
	return ch.deleteResource(fmt.Sprintf("%s/%d", "epics", epicID))
}

// EpicList returns a list of all Epics and their attributes.
//Calls GET https://api.clubhouse.io/api/v1/epics/
func (ch *Clubhouse) EpicList() ([]Epic, error) {
	body, err := ch.listResources("epics")
	if err != nil {
		return []Epic{}, err
	}
	epics := []Epic{}
	json.Unmarshal(body, &epics)
	return epics, nil
}

//EpicCreate allows you to create a new Epic in Clubhouse.
//Calls POST https://api.clubhouse.io/api/v1/epics/
func (ch *Clubhouse) EpicCreate(newEpic CreateEpic) (Epic, error) {
	jsonStr, _ := json.Marshal(newEpic)
	body, err := ch.createObject("epics", jsonStr)
	if err != nil {
		return Epic{}, err
	}
	epic := Epic{}
	json.Unmarshal(body, &epic)
	return epic, nil
}
