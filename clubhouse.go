package clubhouse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Clubhouse struct {
	Token  string
	Client *http.Client
}

func New(token string) *Clubhouse {
	return &Clubhouse{
		Token:  token,
		Client: &http.Client{},
	}
}

func getURL(kind string, token string) string {
	return fmt.Sprintf("%s%s?token=%s", "https://api.clubhouse.io/api/v1/", kind, token)
}

func getURLWithID(kind string, id int64, token string) string {
	return fmt.Sprintf("%s%s/%d?token=%s", "https://api.clubhouse.io/api/v1/", kind, id, token)
}

func (ch *Clubhouse) GetEpic(epicId int64) (Epic, error) {
	req, err := http.NewRequest("GET", getURLWithID("epics", epicId, ch.Token), nil)
	if err != nil {
		return Epic{}, err
	}
	resp, err := ch.Client.Do(req)
	if err != nil {
		return Epic{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return Epic{}, fmt.Errorf("API Returned HTTP Status Code of %d", resp.StatusCode)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	epic := Epic{}
	json.Unmarshal(body, &epic)
	return epic, nil
}

func (ch *Clubhouse) UpdateEpic(updatedEpic UpdateEpic) (Epic, error) {
	jsonStr, _ := json.Marshal(updatedEpic)
	req, err := http.NewRequest("PUT", getURL("epics", ch.Token), bytes.NewBuffer(jsonStr))
	if err != nil {
		return Epic{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := ch.Client.Do(req)
	if err != nil {
		return Epic{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return Epic{}, fmt.Errorf("API Returned HTTP Status Code of %d", resp.StatusCode)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	epic := Epic{}
	json.Unmarshal(body, &epic)
	return epic, nil
}

func (ch *Clubhouse) DeleteEpic(epicId int64) error {
	req, err := http.NewRequest("DELETE", getURLWithID("epics", epicId, ch.Token), nil)
	if err != nil {
		return err
	}

	resp, err := ch.Client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf("API Returned HTTP Status Code of %d", resp.StatusCode)
	}

	return nil
}

func (ch *Clubhouse) ListEpics() ([]Epic, error) {
	req, err := http.NewRequest("GET", getURL("epics", ch.Token), nil)
	if err != nil {
		return []Epic{}, err
	}
	resp, err := ch.Client.Do(req)
	if err != nil {
		return []Epic{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return []Epic{}, fmt.Errorf("API Returned HTTP Status Code of %d", resp.StatusCode)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	epics := []Epic{}
	json.Unmarshal(body, &epics)
	return epics, nil
}

func (ch *Clubhouse) CreateEpic(newEpic CreateEpic) (Epic, error) {
	jsonStr, _ := json.Marshal(newEpic)
	req, err := http.NewRequest("POST", getURL("epics", ch.Token), bytes.NewBuffer(jsonStr))
	if err != nil {
		return Epic{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := ch.Client.Do(req)
	if err != nil {
		return Epic{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		return Epic{}, fmt.Errorf("API Returned HTTP Status Code of %d", resp.StatusCode)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	epic := Epic{}
	json.Unmarshal(body, &epic)
	return epic, nil
}
