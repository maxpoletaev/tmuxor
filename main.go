package main

import (
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v1"
)

var (
	Shell      = "/bin/bash"
	ConfigFile = ".tmuxor.yml"
)

type Window struct {
	Name    string `yaml:"name"`
	WorkDir string `yaml:"workdir"`
	Cmd     string `yaml:"cmd"`
}

type Session struct {
	Name    string   `yaml:"name"`
	Windows []Window `yaml:"windows"`
}

func main() {
	b, err := ioutil.ReadFile(".tmuxor.yml")
	if err != nil {
		err = errors.Wrap(err, "Unable to open .tmuxor.yml")
		log.Fatal(err)
	}

	session := Session{}
	err = yaml.Unmarshal(b, &session)
	if err != nil {
		err = errors.Wrap(err, "Unable to read .tmuxor.yml")
		log.Fatal(err)
	}

	StartTmuxSession(session.Name)
	for _, window := range session.Windows {
		CreateTmuxWindow(session.Name, window.Name, window.WorkDir)
		RunTmuxCommand(session.Name, window.Name, window.Cmd)
	}
}
