package main

import (
	"testing"

	"github.com/enr/runcmd"
)

var executions = []CommandExecution{
	{
		Command: &runcmd.Command{
			Exe:  "../../bin/dups",
			Args: []string{"."},
		},
		Success:  true,
		ExitCode: 0,
		Stdout:   "Produced 5",
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
