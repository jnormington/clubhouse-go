package clubhouse

import (
	"encoding/json"
	"time"
)

type CHFile struct {
	ContentType  string    `json:"content_type"`
	CreatedAt    time.Time `json:"created_at"`
	Description  string    `json:"description"`
	Filename     string    `json:"filename"`
	ID           int64     `json:"id"`
	MentionIds   []string  `json:"mention_ids"`
	Name         string    `json:"name"`
	Size         int64     `json:"size"`
	StoryIds     []int64   `json:"story_ids"`
	ThumbnailURL string    `json:"thumbnail_url"`
	UpdatedAt    time.Time `json:"updated_at"`
	UploaderID   string    `json:"uploader_id"`
	URL          string    `json:"url"`
}

type CHUpdateFile struct {
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	ExternalID  string    `json:"external_id"`
	Name        string    `json:"name"`
	UpdatedAt   time.Time `json:"updated_at"`
	UploaderID  string    `json:"uploader_id"`
}

func (ch *Clubhouse) GetFile(fileID int64) (CHFile, error) {
	body, err := ch.getResource("files", fileID)
	if err != nil {
		return CHFile{}, err
	}
	file := CHFile{}
	json.Unmarshal(body, &file)
	return file, nil
}

func (ch *Clubhouse) UpdateFile(updatedFile CHUpdateFile, fileID int64) (CHFile, error) {
	jsonStr, _ := json.Marshal(updatedFile)
	body, err := ch.updateResource("files", fileID, jsonStr)
	if err != nil {
		return CHFile{}, err
	}
	file := CHFile{}
	json.Unmarshal(body, &file)
	return file, nil
}

func (ch *Clubhouse) DeleteFile(fileID int64) error {
	return ch.deleteResource("files", fileID)
}

func (ch *Clubhouse) ListFiles() ([]CHFile, error) {
	body, err := ch.listResources("files")
	if err != nil {
		return []CHFile{}, err
	}
	files := []CHFile{}
	json.Unmarshal(body, &files)
	return files, nil
}
