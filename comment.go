package clubhouse

import (
	"encoding/json"
	"fmt"
	"time"
)

type ThreadedComment struct {
	AuthorID  string            `json:"author_id"`
	CreatedAt time.Time         `json:"created_at"`
	Deleted   bool              `json:"deleted"`
	ID        int64             `json:"id"`
	Text      string            `json:"text"`
	UpdatedAt time.Time         `json:"updated_at"`
	Comments  []ThreadedComment `json:"comments"`
}

type CreateComment struct {
	AuthorID   string    `json:"author_id"`
	CreatedAt  time.Time `json:"created_at"`
	ExternalID string    `json:"external_id"`
	Text       string    `json:"text"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Comment struct {
	AuthorID   string    `json:"author_id"`
	CreatedAt  time.Time `json:"created_at"`
	ID         int64     `json:"id"`
	MentionIds []string  `json:"mention_ids"`
	Position   int64     `json:"position"`
	StoryID    int64     `json:"story_id"`
	Text       string    `json:"text"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UpdateComment struct {
	Text string `json:"text"`
}

func (ch *Clubhouse) CreateComment(newComment CreateComment, storyID int64) (Comment, error) {
	jsonStr, _ := json.Marshal(newComment)

	body, err := ch.createObject(fmt.Sprintf("%s/%d/%s", "stories", storyID, "comments"), jsonStr)
	if err != nil {
		return Comment{}, err
	}
	comment := Comment{}
	json.Unmarshal(body, &comment)
	return comment, nil
}

func (ch *Clubhouse) GetComment(storyID int64, commentID int64) (Comment, error) {
	body, err := ch.getResource(fmt.Sprintf("%s/%d/%s/%d", "stories", storyID, "comments", commentID))
	if err != nil {
		return Comment{}, err
	}
	comment := Comment{}
	json.Unmarshal(body, &comment)
	return comment, nil
}

func (ch *Clubhouse) UpdateComment(updatedComment UpdateComment, storyID int64, commentID int64) (Comment, error) {
	jsonStr, _ := json.Marshal(updatedComment)
	body, err := ch.updateResource(fmt.Sprintf("%s/%d/%s/%d", "stories", storyID, "comments", commentID), jsonStr)
	if err != nil {
		return Comment{}, err
	}
	comment := Comment{}
	json.Unmarshal(body, &comment)
	return comment, nil
}

func (ch *Clubhouse) DeleteComment(storyID int64, commentID int64) error {
	return ch.deleteResource(fmt.Sprintf("%s/%d/%s/%d", "stories", storyID, "comments", commentID))
}
