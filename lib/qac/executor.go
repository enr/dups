package qac

import (
	"github.com/enr/runcmd"
)

type executionResult struct {
	success  bool
	exitCode int
	stdout   string
	stderr   string
}
type executor interface {
	execute(c Command) executionResult
}
type runcmdExecutor struct {
}

func (e *runcmdExecutor) execute(c Command) executionResult {
	command := e.toRuncmd(c)
	res := command.Run()
	return executionResult{
		success:  res.Success(),
		exitCode: res.ExitStatus(),
		stdout:   res.Stdout().String(),
		stderr:   res.Stderr().String(),
	}
}

func (e *runcmdExecutor) toRuncmd(command Command) *runcmd.Command {
	c := &runcmd.Command{
		Exe:         command.Exe,
		Args:        command.Args,
		CommandLine: command.Cli,
		WorkingDir:  command.WorkingDir,
	}
	return c
}
