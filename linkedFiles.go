package clubhouse

import (
	"encoding/json"
	"fmt"
	"time"
)

type LinkedFile struct {
	ContentType  string    `json:"content_type"`
	CreatedAt    time.Time `json:"created_at"`
	Description  string    `json:"description"`
	ID           int64     `json:"id"`
	MentionIds   []string  `json:"mention_ids"`
	Name         string    `json:"name"`
	Size         int64     `json:"size"`
	StoryIds     []int64   `json:"story_ids"`
	ThumbnailURL string    `json:"thumbnail_url"`
	Type         string    `json:"type"`
	UpdatedAt    time.Time `json:"updated_at"`
	UploaderID   string    `json:"uploader_id"`
	URL          string    `json:"url"`
}

type CreateLinkedFile struct {
	ContentType  string `json:"content_type"`
	Description  string `json:"description"`
	Name         string `json:"name"`
	Size         int64  `json:"size"`
	StoryID      int64  `json:"story_id,omitempty"`
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	Type         string `json:"type"`
	UploaderID   string `json:"uploader_id"`
	URL          string `json:"url"`
}

type UpdateLinkedFile struct {
	Description  string `json:"description"`
	Name         string `json:"name"`
	Size         string `json:"size"`
	ThumbnailURL string `json:"thumbnail_url"`
	Type         string `json:"type"`
	UploaderID   string `json:"uploader_id"`
	URL          string `json:"url"`
}

func (ch *Clubhouse) GetLinkedFile(fileID int64) (LinkedFile, error) {
	body, err := ch.getResource(fmt.Sprintf("%s/%d", "linked-files", fileID))
	if err != nil {
		return LinkedFile{}, err
	}
	file := LinkedFile{}
	json.Unmarshal(body, &file)
	return file, nil
}

func (ch *Clubhouse) UpdateLinkedFile(updatedFile UpdateLinkedFile, fileID int64) (LinkedFile, error) {
	jsonStr, _ := json.Marshal(updatedFile)
	body, err := ch.updateResource(fmt.Sprintf("%s/%d", "linked-files", fileID), jsonStr)
	if err != nil {
		return LinkedFile{}, err
	}
	file := LinkedFile{}
	json.Unmarshal(body, &file)
	return file, nil
}

func (ch *Clubhouse) DeleteLinkedFile(fileID int64) error {
	return ch.deleteResource(fmt.Sprintf("%s/%d", "linked-files", fileID))
}

func (ch *Clubhouse) ListLinkedFiles() ([]LinkedFile, error) {
	body, err := ch.listResources("linked-files")
	if err != nil {
		return []LinkedFile{}, err
	}
	files := []LinkedFile{}
	json.Unmarshal(body, &files)
	return files, nil
}

func (ch *Clubhouse) CreateLinkedFiles(newLinkedFile CreateLinkedFile) (LinkedFile, error) {
	jsonStr, _ := json.Marshal(newLinkedFile)
	body, err := ch.createObject("linked-files", jsonStr)
	if err != nil {
		return LinkedFile{}, err
	}

	file := LinkedFile{}
	err = json.Unmarshal(body, &file)
	if err != nil {
		return LinkedFile{}, err
	}

	return file, nil
}
