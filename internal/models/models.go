package models

// Config is the configuration object itself. The root object.
type Config struct {
	Type         string        `json:"type" yaml:"type"`
	Name         string        `json:"name" yaml:"name"`
	Description  string        `json:"description" yaml:"description"`
	TmuxSessions []TmuxSession `json:"tmux_sessions" yaml:"tmux_sessions"`
	TermTabs     []TermTab     `json:"term_tabs" yaml:"term_tabs"`
}

// SSHConfig represents the configuration settings
type SSHConfig struct {
	Host string `json:"host" yaml:"host"`
	// Keyfile is the full path name to the ssh key to use
	// to connect to the host.
	SSHKeyfile string `json:"ssh_keyfile" yaml:"ssh_keyfile"`
	// The user to connect as.
	User string `json:"user" yaml:"user"`
	// Password is the remote user's password.
	// (Please Don't commit this to source control)
	// or host publically.
	Password string `json:"password" yaml:"password"`
	// Enable to manually respond to prompts when connecting.
	// Otherwise we will use the default settings
	// This will essentially just focus the tab and wait for the user input.
	// Don't know that it needs to get any more complicated than that.
	InteractiveMode bool `json:"interactive_mode" yaml:"interactive_mode"`
	// If prompted to accept the host key, automatically accept it without prompting.
	AcceptHost bool `json:"accept_host" yaml:"accept_host"`
}

// TmuxSession represents the configuration settings
// for a tmux session.
type TmuxSession struct {
	// The session index (ie. 0, 1, 2, etc.)
	Pos int64 `json:"pos" yaml:"pos"`
	// The name or unique identifier of the session.
	Id string `json:"id" yaml:"id"`
	// The default windows to open for this session when it is created.
	Windows []TmuxWindow `json:"windows" yaml:"windows"`
}

type TmuxWindow struct {
	// The session index (ie. 0, 1, 2, etc.)
	Pos int64 `json:"pos" yaml:"pos"`
	// The name or unique identifier of the session.
	Id string `json:"id" yaml:"id"`
	// The Tmux panes in the window.
	Panes []TmuxPane `json:"panes" yaml:"panes"`
}

// The pane,
type TmuxPane struct {
}

// TermTab represents the configuration settings for each tab.
type TermTab struct {
	// Pos is the tabs position (ie. 0, 1, 2, etc.)
	// This is in case they are ever out of order.
	Pos int64 `json:"pos" yaml:"pos"`
	// This can be a name or unique identifier to identify the tab.
	Id string `json:"id" yaml:"id"`
	// StartCommands are a list of commands to execute in order
	// when the tab is opened.
	StartCommands []string `json:"start_commands" yaml:"start_commands"`
	// If SSHEnabled is true, the tab will open an SSH session.
	// This corresponds to whether or not SSHConfig is required.
	// You could also just pass the ssh command via StartCommand.
	SSHEnabled bool      `json:"ssh_enabled" yaml:"ssh_enabled"`
	SSHConfig  SSHConfig `json:"ssh_config" yaml:"ssh_config"`
}
