package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Tmux is a binding to tmux
type Tmux struct {
	TmuxCmd string
}

// NewTmux createas a new tmux binding
func NewTmux() *Tmux {
	tmux := &Tmux{"tmux"}
	return tmux
}

// NewSession creates a detached tmux session with with given name if it does not exist
func (t *Tmux) NewSession(name, path, windowName string) error {
	err := exec.Command(t.TmuxCmd, "new-session", "-d", "-s", name, "-c", path, "-n", windowName).Run()
	return err
}

// CreateWindow creates a window inside a tmux session with given name
func (t *Tmux) CreateWindow(session, name, path string) error {
	err := exec.Command(t.TmuxCmd, "new-window", "-t", session, "-n", name, "-c", path).Run()
	return err
}

// SelectWindow selects window inside a tmux session
func (t *Tmux) SelectWindow(session, window string) error {
	path := fmt.Sprintf("%s:%s", session, window)
	err := exec.Command(t.TmuxCmd, "select-window", "-t", path).Run()
	return err
}

// RenameWindow renames a window inside a session
func (t *Tmux) RenameWindow(session, window, newName string) error {
	path := fmt.Sprintf("%s:%s", session, window)
	err := exec.Command(t.TmuxCmd, "rename-window", "-t", path, newName).Run()
	return err
}

// Exec executes a command inside a tmux window
func (t *Tmux) Exec(session, window, command string) error {
	pane := fmt.Sprintf("%s:%s.%d", session, window, 0)
	err := exec.Command(t.TmuxCmd, "send-keys", "-t", pane, command, "C-m").Run()
	return err
}

// Attach attaches you to the tmux session
func (t *Tmux) Attach(session string) error {
	cmd := exec.Command(t.TmuxCmd, "attach-session", "-t", session)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	return err
}
