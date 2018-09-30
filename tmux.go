package main

import (
	"fmt"
	"log"
	"os/exec"
)

// StartTmuxSession created detached tmux session with
// with given name if it does not exist
func StartTmuxSession(name string) {
	err := exec.Command("tmux", "new-session", "-d", "-t", name).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTmuxWindow(session string, name string, path string) {
	err := exec.Command("tmux", "new-window", "-t", session, "-n", name, "-c", path).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func RunTmuxCommand(session, window, command string) {
	pane := fmt.Sprintf("%s:%s.%d", session, window, 0)
	err := exec.Command("tmux", "send-keys", "-t", pane, command, "C-m").Run()
	if err != nil {
		log.Fatal(err)
	}
}
