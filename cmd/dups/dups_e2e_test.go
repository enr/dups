package main

import (
	"testing"

	"github.com/enr/runcmd"
)

var executions = []CommandExecution{
	{
		Command: &runcmd.Command{
			Exe:  "../../bin/dups",
			Args: []string{"../../testdata/01"},
		},
		Success:  false,
		ExitCode: 1,
		Stdout:   "f1d2d2f924e986ac86fdf7b36c94bcdf32beec15 sub/010.txt",
	},
	{
		Command: &runcmd.Command{
			Exe:  "../../bin/dups",
			Args: []string{"--names-only", "../../testdata/01"},
		},
		Success:  false,
		ExitCode: 1,
		Stdout:   "sub/010.txt",
	},
	{
		Command: &runcmd.Command{
			Exe:  "../../bin/dups",
			Args: []string{},
		},
		Success:  false,
		ExitCode: 1,
	},
}

func TestCommandExecution(t *testing.T) {
	VerifyExecutions(t, executions)
}
