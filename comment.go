package clubhouse

import (
	"encoding/json"
	"fmt"
	"time"
)

//ThreadedComment are comments associated with Epic Discussions
type ThreadedComment struct {
	AuthorID  string            `json:"author_id"`
	CreatedAt time.Time         `json:"created_at"`
	Deleted   bool              `json:"deleted"`
	ID        int64             `json:"id"`
	Text      string            `json:"text"`
	UpdatedAt time.Time         `json:"updated_at"`
	Comments  []ThreadedComment `json:"comments"`
}

//CreateComment is the body used for ch.CommentCreate()
//Required fields are:
// CreateComment.Text
type CreateComment struct {
	AuthorID   string    `json:"author_id,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	ExternalID string    `json:"external_id,omitempty"`
	Text       string    `json:"text"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

//A Comment is any note added within the Comment field of a Story.
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

//UpdateComment is the body used for ch.CommentUpdate()
type UpdateComment struct {
	Text string `json:"text"`
}

//CommentCreate allows you to create a Comment on any Story.
// Calls POST https://api.clubhouse.io/api/v1/stories/{storyID}/comments
func (ch *Clubhouse) CommentCreate(newComment CreateComment, storyID int64) (Comment, error) {
	jsonStr, _ := json.Marshal(newComment)

	body, err := ch.createObject(fmt.Sprintf("%s/%d/%s", "stories", storyID, "comments"), jsonStr)
	if err != nil {
		return Comment{}, err
	}
	comment := Comment{}
	json.Unmarshal(body, &comment)
	return comment, nil
}

//CommentGet is used to get Comment information.
//Calls  GET https://api.clubhouse.io/api/v1/stories/{storyID}/comments/{commentID}
func (ch *Clubhouse) CommentGet(storyID int64, commentID int64) (Comment, error) {
	body, err := ch.getResource(fmt.Sprintf("%s/%d/%s/%d", "stories", storyID, "comments", commentID))
	if err != nil {
		return Comment{}, err
	}
	comment := Comment{}
	json.Unmarshal(body, &comment)
	return comment, nil
}

//CommentUpdate replaces the text of the existing Comment.
//Calls PUT https://api.clubhouse.io/api/v1/stories/{storyID}/comments/{commentID}
func (ch *Clubhouse) CommentUpdate(updatedComment UpdateComment, storyID int64, commentID int64) (Comment, error) {
	jsonStr, _ := json.Marshal(updatedComment)
	body, err := ch.updateResource(fmt.Sprintf("%s/%d/%s/%d", "stories", storyID, "comments", commentID), jsonStr)
	if err != nil {
		return Comment{}, err
	}
	comment := Comment{}
	json.Unmarshal(body, &comment)
	return comment, nil
}

//CommentDelete deletes a comment from any story.
//Calls DELETE https://api.clubhouse.io/api/v1/stories/{storyID}/comments/{commentID}
func (ch *Clubhouse) CommentDelete(storyID int64, commentID int64) error {
	return ch.deleteResource(fmt.Sprintf("%s/%d/%s/%d", "stories", storyID, "comments", commentID))
}
