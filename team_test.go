package clubhouse_test

import (
	"io/ioutil"
	"os/user"
	"path/filepath"
	"strings"
	"testing"

	// Frameworks
	ch "github.com/djthorpe/clubhouse-go"
)

// In order to run tests, you will need to put your
// API token in a file called .clubhouse in your
// home directory

func Test_000(t *testing.T) {
	if client := ch.New(""); client == nil {
		t.Error("Invalid client == nil")
	} else {
		t.Log(client)
	}
}

func Test_001(t *testing.T) {
	if token, err := getToken(); err != nil {
		t.Error(err)
	} else if client := ch.New(token).SetDebug(true); client == nil {
		t.Error("Invalid client == nil")
	} else if teams, err := client.ListTeams(); err != nil {
		t.Error(err)
	} else {
		t.Log(teams)
	}
}

/////////////////////////////////////////////////////

func getToken() (string, error) {
	if user, err := user.Current(); err != nil {
		return "", err
	} else if bytes, err := ioutil.ReadFile(filepath.Join(user.HomeDir, ".clubhouse")); err != nil {
		return "", err
	} else {
		return strings.TrimSpace(string(bytes)), nil
	}
}
