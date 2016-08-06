package clubhouse

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiURL string = "https://api.clubhouse.io/api/v1/"

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

func (ch *Clubhouse) getURL(resource string) string {
	return fmt.Sprintf("%s%s?token=%s", apiURL, resource, ch.Token)
}

func (ch *Clubhouse) getResource(resource string) ([]byte, error) {
	req, err := http.NewRequest("GET", ch.getURL(resource), nil)
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

func (ch *Clubhouse) updateResource(resource string, jsonStr []byte) ([]byte, error) {
	req, err := http.NewRequest("PUT", ch.getURL(resource), bytes.NewBuffer(jsonStr))
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

func (ch *Clubhouse) deleteResource(resource string) error {
	req, err := http.NewRequest("DELETE", ch.getURL(resource), nil)
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
	req, err := http.NewRequest("GET", ch.getURL(resource), nil)
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
	req, err := http.NewRequest("POST", ch.getURL(resource), bytes.NewBuffer(jsonStr))
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
