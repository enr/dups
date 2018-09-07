// +build darwin freebsd linux netbsd openbsd

package runcmd

import (
	"fmt"
	"os"
	"path/filepath"
)

func (c *Command) useShell() {
	command := c.CommandLine
	var shell string
	if c.ForceShell != "" {
		shell = c.ForceShell
	} else {
		defaultShell := os.Getenv("SHELL")
		if defaultShell == "" {
			shell = "/bin/sh"
		} else {
			shell = defaultShell
		}
	}
	shellArgument := "-c"
	var shellCommand string
	if c.UseProfile {
		profile := filepath.Join(c.WorkingDir, ".profile")
		// "." is portable, "source" is bash only
		shellCommand = fmt.Sprintf(". \"%s\" 2>/dev/null; %s", profile, command)
	} else {
		shellCommand = command
	}
	c.Exe = shell
	c.Args = []string{shellArgument, shellCommand}
}

// shellAndArgs is a helper function that returns an OS specific
// shell and arguments for that particular shell
func shellAndArgs() (string, []string) {
	var com []string
	com = []string{
		"/bin/bash",
		"-s",
	}
	return com[0], com[1:]
}
