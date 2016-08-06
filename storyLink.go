package clubhouse

import (
	"encoding/json"
	"fmt"
	"time"
)

type CreateStoryLink struct {
	ObjectID  int64  `json:"object_id"`
	SubjectID int64  `json:"subject_id"`
	Verb      string `json:"verb"`
}
type StoryLink struct {
	CreatedAt time.Time `json:"created_at"`
	ID        int64     `json:"id"`
	ObjectID  int64     `json:"object_id"`
	SubjectID int64     `json:"subject_id"`
	UpdatedAt time.Time `json:"updated_at"`
	Verb      string    `json:"verb"`
}

func (ch *Clubhouse) CreateStoryLink(newCreateStoryLink CreateStoryLink) (StoryLink, error) {
	jsonStr, _ := json.Marshal(newCreateStoryLink)

	body, err := ch.createObject("story-links", jsonStr)
	if err != nil {
		return StoryLink{}, err
	}
	storyLink := StoryLink{}
	json.Unmarshal(body, &storyLink)
	return storyLink, nil
}

func (ch *Clubhouse) GetStoryLink(storyLinkID int64) (StoryLink, error) {
	body, err := ch.getResource(fmt.Sprintf("%s/%d", "story-links", storyLinkID))
	if err != nil {
		return StoryLink{}, err
	}
	storyLink := StoryLink{}
	json.Unmarshal(body, &storyLink)
	return storyLink, nil
}

func (ch *Clubhouse) DeleteStoryLink(storyLinkID int64) error {
	return ch.deleteResource(fmt.Sprintf("%s/%d", "story-links", storyLinkID))
}
