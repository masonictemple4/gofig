package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const GONEW_INSTALL_URL = "golang.org/x/tools/cmd/gonew"
const GONEW_DEFAULT_VERSION = "latest"

const (
	ConfigTypeTmux = "tmux"
	ConfigTypeGo   = "go"
)

func IsConfigType(configType string) bool {
	switch configType {
	case ConfigTypeTmux, ConfigTypeGo:
		return true
	}
	return false
}

// Config is the configuration object itself. The root object.
// May tweak a bit more, my initial vision was
// the config would be either local or remote,
// but there are definitely projects with multiple
// hosts, so allowing for multiple terminal tabs,
// both local and remote tmux sessions, you can really
// do whatever you wish.
//
// Example:
//
//	By default the config will open your local terminal
//	you can leave term tabs empty to stick with the single
//	tab and connect to a remote instance. Where you can
//	decide to either use tmux or not.
type Config struct {
	// Confuration type defined above.
	Type string `json:"type" yaml:"type"`
	// Full configuratiion name.
	Name string `json:"name" yaml:"name"`
	// Shorthand name or alias for the config. Kind of like
	// CLI flags (ie., --version, -v) but instead
	Alias string `json:"alias" yaml:"alias"`
	// A description for the configuration.
	Description string `json:"description" yaml:"description"`
	// Project path - this is the path to the project.
	ProjectPath string `json:"project_path" yaml:"project_path"`
	// Tmux settings for this environment.
	Tmux      TmuxConfig `json:"tmux" yaml:"tmux"`
	GoProject GoConfig   `json:"go_project" yaml:"go_project"`
}

func New(path string) *Config {

	ext := filepath.Ext(path)

	var conf *Config

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	switch ext {
	case ".json":
		json.NewDecoder(file).Decode(conf)
	case ".yaml":
		yaml.NewDecoder(file).Decode(conf)
	default:
		panic("Unsupported configuration file type. Please use .json or .yaml")
	}

	if conf == nil {
		panic("Failed to generate configuration.")
	}

	return conf

}

type TmuxConfig struct {
	// Local, Remote, or possibly even docker.
	ExecutionEnv string     `json:"execution_env" yaml:"execution_env"`
	Host         HostConfig `json:"host" yaml:"host"`
	// Sessions to open.
	Sessions []TmuxSession `json:"tmux_sessions" yaml:"tmux_sessions"`
	// If you want all sessions inside of your tmux instance
	// to use the same layout you can specify it here.
	// You can override this at session, window, or pane level.
	Layout string `json:"global_layout" yaml:"global_layout"`
}

// HostConfig represents the configuration settings
// When you're running a local environment you can
// leave everything else as it's default.
type HostConfig struct {
	// Local, Remote
	// Future support for docker. There are various
	// different methods of using docker for development
	// environments.
	Host string `json:"host" yaml:"host"`
	// Port to connect to.
	// I'm not sure we'll actually need this but just in case.
	Port int64 `json:"port" yaml:"port"`
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
	// Primarily used with SSH.
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
	// Commands to run at the session level when created.
	StartCommands []string `json:"start_commands" yaml:"start_commands"`
	// Define window layout at the session level.
	// can be overriden at the window level.
	Layout string `json:"layout" yaml:"layout"`
}

type TmuxWindow struct {
	// The session index (ie. 0, 1, 2, etc.)
	Pos int64 `json:"pos" yaml:"pos"`
	// The name or unique identifier of the session.
	Id string `json:"id" yaml:"id"`
	// The Tmux panes in the window.
	Panes []TmuxWindowPane `json:"panes" yaml:"panes"`
	// Commands to run at the session window level when created.
	StartCommands []string `json:"start_commands" yaml:"start_commands"`
	// The layout for the individual window.
	// Layout determines how we configure the panes within the window.
	Layout string `json:"layout" yaml:"layout"`
}

// Pane within a window.
type TmuxWindowPane struct {
	// Pane identifier
	Id string `json:"id" yaml:"id"`
	// Commands to run at the window pane level when created.
	StartCommands []string `json:"start_commands" yaml:"start_commands"`
}

// Configuration for go projects.
type GoConfig struct {
	// Full path to the root dir.
	Path string `json:"path" yaml:"path"`
	// Template path can be a local project or a remote project.
	TemplatePath string `json:"target_url" yaml:"target_url"`
	// Mod path is the value you specify when running go mod init <mod_path>
	ModPath string `json:"mod_path" yaml:"mod_path"`
}
