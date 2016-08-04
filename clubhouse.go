package clubhouse

import (
	"bytes"
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

func getURL(resource string, token string) string {
	return fmt.Sprintf("%s%s?token=%s", "https://api.clubhouse.io/api/v1/", resource, token)
}

func getURLWithID(resource string, resourceID int64, token string) string {
	return fmt.Sprintf("%s%s/%d?token=%s", "https://api.clubhouse.io/api/v1/", resource, resourceID, token)
}

func (ch *Clubhouse) getResource(resource string, resourceID int64) ([]byte, error) {
	req, err := http.NewRequest("GET", getURLWithID(resource, resourceID, ch.Token), nil)
	if err != nil {
		return []byte{}, err
	}
	resp, err := ch.Client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return []byte{}, fmt.Errorf("API Returned HTTP Status Code of %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

func (ch *Clubhouse) updateResource(resource string, resourceID int64, jsonStr []byte) ([]byte, error) {
	req, err := http.NewRequest("PUT", getURLWithID(resource, resourceID, ch.Token), bytes.NewBuffer(jsonStr))
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := ch.Client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return []byte{}, fmt.Errorf("API Returned HTTP Status Code of %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

func (ch *Clubhouse) deleteResource(resource string, resourceID int64) error {
	req, err := http.NewRequest("DELETE", getURLWithID(resource, resourceID, ch.Token), nil)
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

func (ch *Clubhouse) listResources(resource string) ([]byte, error) {
	req, err := http.NewRequest("GET", getURL(resource, ch.Token), nil)
	if err != nil {
		return []byte{}, err
	}
	resp, err := ch.Client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return []byte{}, fmt.Errorf("API Returned HTTP Status Code of %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

func (ch *Clubhouse) createObject(resource string, jsonStr []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", getURL(resource, ch.Token), bytes.NewBuffer(jsonStr))
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := ch.Client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		return []byte{}, fmt.Errorf("API Returned HTTP Status Code of %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}
