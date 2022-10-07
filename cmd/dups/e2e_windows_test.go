//go:build windows
// +build windows

package main

import (
	"testing"

	"github.com/enr/runcmd"
)

var executions = []commandExecution{
	{
		Command: &runcmd.Command{
			Exe:  "../../bin/dups",
			Args: []string{"--dups-exit", "../../testdata/01"},
		},
		Success:  false,
		ExitCode: 2,
		Stdout:   "855426068ee8939df6bce2c2c4b1e7346532a133 sub/010.txt",
	},
	{
		Command: &runcmd.Command{
			Exe:  "../../bin/dups",
			Args: []string{"--dups-exit", "--names-only", "../../testdata/01"},
		},
		Success:  false,
		ExitCode: 2,
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
	verifyExecutions(t, executions)
}
