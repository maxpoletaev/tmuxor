package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v1"
)

// Window describes a window inside a session
type Window struct {
	Name    string `yaml:"name"`
	WorkDir string `yaml:"workdir"`
	Cmd     string `yaml:"cmd"`
}

// Session describes a tmux session
type Session struct {
	Name          string   `yaml:"name"`
	WorkDir       string   `yaml:"workdir"`
	Windows       []Window `yaml:"windows"`
	Detached      bool     `yaml:"detached"`
	DefaultWindow string   `yaml:"default_window"`
}

// Config is a root element of YAML file
type Config struct {
	Session Session `yaml:"session"`
}

// GetWorkingDir returns current working directory
func GetWorkingDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

// ReadConfig reads configuration from the YAML file
func ReadConfig(configName string) (config Config, err error) {
	b, err := ioutil.ReadFile(configName)
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("Unable to open %s", configName))
		return config, err
	}

	err = yaml.Unmarshal(b, &config)
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("Unable to read %s", configName))
		return config, err
	}

	return config, nil
}

// CleanConfig checks that all the properties filled properly and provides default values
func CleanConfig(config *Config) error {
	session := config.Session
	if session.Name == "" {
		return errors.New("Session should contain name (key: session.name)")
	}
	if len(session.Windows) == 0 {
		return errors.New("Session should contain at leat one window (key: session.windows)")
	}
	if session.WorkDir == "" {
		session.WorkDir = GetWorkingDir()
	}

	for i, window := range session.Windows {
		if window.Name == "" {
			return errors.Errorf("Window should contain name (key: session.windows.%d.name)", i)
		}
		if window.Cmd == "" {
			return errors.Errorf("Window should contain command (key: session.windows.%d.cmd)", i)
		}
	}

	return nil
}

// getConfigName returns config name from args or default config name
func getConfigName() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return ".tmuxor.yml"
}

func main() {
	configName := getConfigName()
	config, err := ReadConfig(configName)
	if err != nil {
		log.Fatal(err)
	}

	err = CleanConfig(&config)
	if err != nil {
		log.Fatal(err)
	}

	tmux := NewTmux()
	session := config.Session
	err = tmux.StartSession(session.Name)
	if err != nil {
		log.Fatal(err)
	}

	for _, window := range session.Windows {
		workDir := session.WorkDir
		if window.WorkDir != "" {
			workDir = window.WorkDir
		}

		err = tmux.CreateWindow(session.Name, window.Name, workDir)
		if err != nil {
			log.Fatal(err)
		}

		err = tmux.Exec(session.Name, window.Name, window.Cmd)
		if err != nil {
			log.Fatal(err)
		}
	}

	if session.DefaultWindow != "" {
		tmux.SelectWindow(session.Name, session.DefaultWindow)
	}

	if !session.Detached {
		tmux.Attach(session.Name)
	}
}
