package tmux

import (
	"encoding/json"
	"testing"
)

func TestListSessions(t *testing.T) {
	sessionList := GetSessionsList()

	data, _ := json.MarshalIndent(sessionList, "", "  ")
	t.Logf("Session list:\n%s\n", string(data))

}
