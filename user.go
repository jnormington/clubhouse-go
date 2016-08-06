package clubhouse

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Deactivated            bool         `json:"deactivated"`
	ID                     string       `json:"id"`
	Name                   string       `json:"name"`
	Permissions            []Permission `json:"permissions"`
	TwoFactorAuthActivated bool         `json:"two_factor_auth_activated"`
	Username               string       `json:"username"`
}

type Permission struct {
	CreatedAt    string    `json:"created_at"`
	Disabled     bool      `json:"disabled"`
	EmailAddress string    `json:"email_address"`
	GravatarHash string    `json:"gravatar_hash"`
	ID           string    `json:"id"`
	Initials     string    `json:"initials"`
	Role         string    `json:"role"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (ch *Clubhouse) GetUser(userID int64) (User, error) {
	body, err := ch.getResource(fmt.Sprintf("%s/%d", "users", userID))
	if err != nil {
		return User{}, err
	}
	user := User{}
	json.Unmarshal(body, &user)
	return user, nil
}

func (ch *Clubhouse) ListUsers() ([]User, error) {
	body, err := ch.listResources("users")
	if err != nil {
		return []User{}, err
	}
	users := []User{}
	json.Unmarshal(body, &users)
	return users, nil
}
