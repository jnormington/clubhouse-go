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

func (ch *Clubhouse) GetEpic(epicID int64) (Epic, error) {
	req, err := http.NewRequest("GET", getURLWithID("epics", epicID, ch.Token), nil)
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

func (ch *Clubhouse) UpdateEpic(updatedEpic UpdateEpic, epicID int64) (Epic, error) {
	jsonStr, _ := json.Marshal(updatedEpic)
	req, err := http.NewRequest("PUT", getURLWithID("epics", epicID, ch.Token), bytes.NewBuffer(jsonStr))
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

func (ch *Clubhouse) DeleteEpic(epicID int64) error {
	req, err := http.NewRequest("DELETE", getURLWithID("epics", epicID, ch.Token), nil)
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

func (ch *Clubhouse) ListLabels() ([]LabelWithCounts, error) {
	req, err := http.NewRequest("GET", getURL("labels", ch.Token), nil)
	if err != nil {
		return []LabelWithCounts{}, err
	}
	resp, err := ch.Client.Do(req)
	if err != nil {
		return []LabelWithCounts{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return []LabelWithCounts{}, fmt.Errorf("API Returned HTTP Status Code of %d", resp.StatusCode)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	labels := []LabelWithCounts{}
	json.Unmarshal(body, &labels)
	return labels, nil
}

func (ch *Clubhouse) CreateLabel(newLabel CreateLabel) (LabelWithCounts, error) {
	jsonStr, _ := json.Marshal(newLabel)
	req, err := http.NewRequest("POST", getURL("labels", ch.Token), bytes.NewBuffer(jsonStr))
	if err != nil {
		return LabelWithCounts{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := ch.Client.Do(req)
	if err != nil {
		return LabelWithCounts{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		return LabelWithCounts{}, fmt.Errorf("API Returned HTTP Status Code of %d", resp.StatusCode)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	label := LabelWithCounts{}
	json.Unmarshal(body, &label)
	return label, nil
}

func (ch *Clubhouse) UpdateLabel(updatedLabel UpdateLabel, labelID int64) (LabelWithCounts, error) {
	jsonStr, _ := json.Marshal(updatedLabel)
	req, err := http.NewRequest("PUT", getURLWithID("labels", labelID, ch.Token), bytes.NewBuffer(jsonStr))
	if err != nil {
		return LabelWithCounts{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := ch.Client.Do(req)
	if err != nil {
		return LabelWithCounts{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return LabelWithCounts{}, fmt.Errorf("API Returned HTTP Status Code of %d", resp.StatusCode)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	label := LabelWithCounts{}
	json.Unmarshal(body, &label)
	return label, nil
}

func (ch *Clubhouse) DeleteLabel(labelID int64) error {
	req, err := http.NewRequest("DELETE", getURLWithID("labels", labelID, ch.Token), nil)
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
