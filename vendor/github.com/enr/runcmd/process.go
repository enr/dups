package runcmd

import (
	"os"
	"path/filepath"
)

// Start start a process without waiting
func (c *Command) Start() error {

	cmd, err := c.buildCmd()
	if err != nil {
		return err
	}
	cmd.Stdout = logfile(c.GetLogfile())
	cmd.Stderr = logfile(c.GetLogfile())
	if c.WorkingDir != "" {
		cmd.Dir = c.WorkingDir
	}

	if c.UseEnv {
		flagEnv := filepath.Join(cmd.Dir, ".env")
		env, _ := ReadEnv(flagEnv)
		cmd.Env = env.asArray()
	} else if len(c.Env) > 0 {
		cmd.Env = c.Env.asArray()
	}

	pid, err := start(cmd)
	if err != nil {
		return err
	}
	c.Pid = pid
	return nil
}

func logfile(path string) *os.File {
	if path == "" {
		return nil
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		return nil
	}
	return file
}
