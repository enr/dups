//go:build darwin || freebsd || linux || netbsd || openbsd
// +build darwin freebsd linux netbsd openbsd

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
		Stdout:   "f1d2d2f924e986ac86fdf7b36c94bcdf32beec15 sub/010.txt",
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
			Args: []string{"--dups-exit", abs("../../testdata/02")},
		},
		Success:  false,
		ExitCode: 6,
		Stdout:   "5619cedcdc1d07a16eb2cb8f132ecdd08e1d699a exc/b2",
	},
	{
		Command: &runcmd.Command{
			Exe:  "../../bin/dups",
			Args: []string{"--dups-exit", "--include", "f*", abs("../../testdata/02")},
		},
		Success:  false,
		ExitCode: 4,
		Stdout:   "e630d1f8dd29477ad933ee8355f9b9712bcb8fe4 inc/f2",
	},
	{
		Command: &runcmd.Command{
			Exe:  "../../bin/dups",
			Args: []string{"--dups-exit", "--exclude", "exc", abs("../../testdata/02")},
		},
		Success:  false,
		ExitCode: 3,
		Stdout:   "e630d1f8dd29477ad933ee8355f9b9712bcb8fe4 inc/f3",
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
