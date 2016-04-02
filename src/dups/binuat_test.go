package main

import (
	"strings"
	"testing"

	"github.com/enr/runcmd"
)

type CommandExecution struct {
	Command  *runcmd.Command
	Success  bool
	ExitCode int
	Stdout   string
	Stderr   string
}

func VerifyExecution(t *testing.T, execution CommandExecution) {
	command := execution.Command
	res := command.Run()
	if res.Success() != execution.Success {
		t.Fatalf("`%s`: expected success %t but got %t", command, execution.Success, res.Success())
	}
	expectedCode := execution.ExitCode
	actualCode := res.ExitStatus()
	if actualCode != expectedCode {
		t.Fatalf("%s: expected exit code %d but got %d", command, expectedCode, actualCode)
	}
	assertStringContains(t, res.Stdout().String(), execution.Stdout)
	assertStringContains(t, res.Stderr().String(), execution.Stderr)
}

func VerifyExecutions(t *testing.T, executions []CommandExecution) {
	for _, command := range executions {
		VerifyExecution(t, command)
	}
}

func assertStringContains(t *testing.T, s string, substr string) {
	if substr != "" && !strings.Contains(s, substr) {
		t.Fatalf("expected output\n%s\n  does not contain\n%s\n", s, substr)
	}
}
