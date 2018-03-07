package clubhouse

import (
	"encoding/json"
	"fmt"
)

type Member struct {
	ID      string  `json:"id"`
	Profile Profile `json:"profile"`
}

type Profile struct {
	Deactivated            bool   `json:"deactivated"`
	EmailAddress           string `json:"email_address"`
	Name                   string `json:"name"`
	TwoFactorAuthActivated bool   `json:"two_factor_auth_activated"`
	Membername             string `json:"membername"`
}

func (ch *Clubhouse) GetMember(memberID int64) (Member, error) {
	body, err := ch.getResource(fmt.Sprintf("%s/%d", "members", memberID))
	if err != nil {
		return Member{}, err
	}

	member := Member{}
	json.Unmarshal(body, &member)

	return member, nil
}

func (ch *Clubhouse) ListMembers() ([]Member, error) {
	body, err := ch.listResources("members")
	if err != nil {
		return []Member{}, err
	}
	members := []Member{}
	json.Unmarshal(body, &members)
	return members, nil
}
