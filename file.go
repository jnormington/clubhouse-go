package clubhouse

import (
	"encoding/json"
	"fmt"
	"time"
)

// A CHFile is any document uploaded to your Clubhouse. Files attached from a third-party service can be accessed using the Linked Files endpoint.
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

//CHUpdateFile is the body used for ch.FileUpdate()
type CHUpdateFile struct {
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	ExternalID  string    `json:"external_id"`
	Name        string    `json:"name"`
	UpdatedAt   time.Time `json:"updated_at"`
	UploaderID  string    `json:"uploader_id"`
}

// FileGet returns information about the selected File.
// Calls GET https://api.clubhouse.io/api/v1/files/{fileID}
func (ch *Clubhouse) FileGet(fileID int64) (CHFile, error) {
	body, err := ch.getResource(fmt.Sprintf("%s/%d", "files", fileID))
	if err != nil {
		return CHFile{}, err
	}
	file := CHFile{}
	json.Unmarshal(body, &file)
	return file, nil
}

// FileUpdate can used to update the properties of a file uploaded to Clubhouse.
// Calls PUT https://api.clubhouse.io/api/v1/files/{fileID}
func (ch *Clubhouse) FileUpdate(updatedFile CHUpdateFile, fileID int64) (CHFile, error) {
	jsonStr, _ := json.Marshal(updatedFile)
	body, err := ch.updateResource(fmt.Sprintf("%s/%d", "files", fileID), jsonStr)
	if err != nil {
		return CHFile{}, err
	}
	file := CHFile{}
	json.Unmarshal(body, &file)
	return file, nil
}

// FileDelete can be used to delete any previously attached File.
// Calls DELETE https://api.clubhouse.io/api/v1/files/{fileID}
func (ch *Clubhouse) FileDelete(fileID int64) error {
	return ch.deleteResource(fmt.Sprintf("%s/%d", "files", fileID))
}

// FileList returns a list of all Files and related attributes in your Clubhouse.
// Calls GET https://api.clubhouse.io/api/v1/files
func (ch *Clubhouse) FileList() ([]CHFile, error) {
	body, err := ch.listResources("files")
	if err != nil {
		return []CHFile{}, err
	}
	files := []CHFile{}
	json.Unmarshal(body, &files)
	return files, nil
}
