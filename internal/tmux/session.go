package tmux

// TODO: Move tmux session related functions and models here.
type Session struct {
	Id      int64
	Name    string
	WorkDir string
	Windows []Window
}

func NewSession() *Session {
	return &Session{}
}

func (s *Session) Attach() error {
	args := []string{}

	return nil
}

func CurrentSession() *Session {
	args := []string{
		"display-message",
		"-p",
		"#{session_name}",
	}
	return nil
}
