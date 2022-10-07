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
			Args: []string{"--dups-exit", abs("../../testdata/01")},
		},
		Success:  false,
		ExitCode: 2,
		Stdout:   "855426068ee8939df6bce2c2c4b1e7346532a133 sub/010.txt",
	},
	{
		Command: &runcmd.Command{
			Exe:  "../../bin/dups",
			Args: []string{"--dups-exit", "--names-only", abs("../../testdata/01")},
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
		Success:  true,
		ExitCode: 0,
	},
}

func TestCommandExecution(t *testing.T) {
	verifyExecutions(t, executions)
}
