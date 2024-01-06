package tmux

import (
	"fmt"
	"strconv"
)

const SESSION_FORMAT = "#{session_id}|#{session_name}|#{session_path}"

type SessionField int

const (
	SessionId SessionField = iota
	SessionName
	SessionPath
)

type Session struct {
	// session_id
	Id int64 `json:"id" yaml:"id"`
	// session_name
	Name string `json:"name" yaml:"name"`
	// this will be the default path every new window opens in.
	// this can be overridden at the window and pane level.
	// session_path
	WorkDir string   `json:"work_dir" yaml:"work_dir"`
	Windows []Window `json:"windows" yaml:"windows"`
}

func GetSessionsList() *[]Session {

	args := []string{
		"list-sessions",
		"-F",
		SESSION_FORMAT,
	}

	output, err := Exec(args)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	sessions := SessionsFromString(output)

	return sessions
}

func SessionsFromString(input string) *[]Session {
	var sessions []Session

	sessionStrs := splitLines(input)

	for _, sessionStr := range sessionStrs {
		session := sessionFromString(sessionStr)
		session.Windows = session.GetWindows()
		sessions = append(sessions, session)
	}

	return &sessions
}

func sessionFromString(input string) Session {
	fields := splitFields(input)

	// Remove the first character from the session id.
	fields[SessionId] = fields[SessionId][1:]

	sid, _ := strconv.ParseInt(fields[SessionId], 10, 64)

	return Session{
		Id:      sid,
		Name:    fields[SessionName],
		WorkDir: fields[SessionPath],
	}
}

func GetCurrentSessionName() string {
	args := []string{
		"display-message",
		"-F",
		"#{sesion_name}",
	}

	name, err := Exec(args)
	if err != nil {
		panic(err)
	}

	return name
}

func GetCurrentSession() *Session {
	args := []string{
		"display-message",
		"-F",
		SESSION_FORMAT,
		"#S",
	}

	output, err := Exec(args)
	if err != nil {
		panic(err)
	}

	session := sessionFromString(output)

	return &session

}

func (s *Session) GetWindows() []Window {
	args := []string{
		"list-windows",
		"-t",
		s.Name,
		"-F",
		WINDOW_FORMAT,
	}

	output, err := Exec(args)
	if err != nil {
		panic(err)
	}

	windows := WindowsFromString(output)

	return *windows
}
