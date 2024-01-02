package tmux

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/masonictemple4/gofig/internal/config"
)

func ProcessConfig(confName string) {
	opts := config.New(confName)

	if opts == nil {
		panic("Failed to generate configuration.")
	}

	switch opts.Type {
	case config.ConfigTypeTerm:
		processTerm(opts)
	case config.ConfigTypeTmux:
		processTmux(opts)
	case config.ConfigTypeGo:
		processGo(opts)
	}

	// Figure out if it's term go or tmux

}

func processTmux(config *config.Config) {
}

func processTerm(config *config.Config) {
}

func processGo(conf *config.Config) {
	gonew, err := exec.LookPath("gonew")
	if err != nil {
		panic(err)
	}

	args := append([]string{gonew}, conf.GoProject.TemplatePath, conf.GoProject.ModPath)
	err = syscall.Exec(gonew, args, os.Environ())
	if err != nil {
		panic(err)
	}

	// Navigate to new dir
	err = os.Chdir(conf.GoProject.Path)
	if err != nil {
		panic(err)
	}

}
