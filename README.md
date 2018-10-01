# Tmuxor

Tmuxor was inspired by a wonderful project called [tmuxinator](https://github.com/tmuxinator/tmuxinator) and provides alternative way to describe tmux sessions in a single yaml file to run them with one click.

## Why not tmuxinator?

Tmuxor is written in Go, so it doesn’t require Ruby or any other dependency at all. That could be much more convenient if you use tmux on a server or inside a Docker container and don’t want to install any additional dependencies.

## Installation

You can download the latest release for Linux and MacOS from the [releases](https://github.com/zenwalker/tmuxor/releases) page.

Also, if you’re a Go user, you can install it with Go package manager:

```
go install github.com/zenwalker/tmuxor
```

## Quick start

1. Create a file `.tmuxor.yml` with the following content:

    ```yaml
    session:
      name: test
      detached: false
      startup_window: first

      windows:
        - name: first
          cmd: echo hello

        - name: second
          cmd: sleep 10 && echo "world"
    ```

2. Run `tmuxor` in the console.

## Does it support panes?

No. I don’t use ones so that there is no panes support. But if you want such a feature, I’d be happy to merge your PR.
