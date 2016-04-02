// +build windows

package runcmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type ccli struct {
	command *Command
	cli     string
}

var testdata = []ccli{
	{command: &Command{},
		cli: "",
	},
	{command: &Command{
		CommandLine: `echo "home=%HOME%"`,
		UseEnv:      true,
	},
		cli: `cmd /C echo "home=%HOME%"`,
	},
}

func TestCluiVerbosityLevelMute(t *testing.T) {
	for _, d := range testdata {
		cmd := d.command
		actual := cmd.FullCommand()
		expected := d.cli
		if actual != expected {
			t.Fatalf("%s: expected %s but got %s", cmd, expected, actual)
		}
	}
}

type commandExecution struct {
	command  *Command
	success  bool
	exitCode int
	stdout   string
	stderr   string
}

var executions = []commandExecution{
	{
		command:  &Command{},
		success:  false,
		exitCode: 1,
	},
	{
		command: &Command{
			CommandLine: "/not/found/here",
		},
		success:  false,
		exitCode: 1,
		stderr:   "The system cannot find the path specified.",
	},
	{
		command: &Command{
			Exe: "/not/found/here",
		},
		success:  false,
		exitCode: 1,
	},
	{
		command: &Command{
			CommandLine: "dir",
			//Args: []string{"/bin"},
		},
		success:  true,
		exitCode: 0,
		stdout:   "_",
	},
	{
		command: &Command{
			CommandLine: `echo BAR=%BAR%!`,
		},
		success:  true,
		exitCode: 0,
		// on windows no substitution for undeclared variable
		stdout: "BAR=%BAR%!",
	},
	{
		command: &Command{
			CommandLine: `echo "BAR=%BAR%!"`,
			Env:         Env{"BAR": "foo"},
		},
		success:  true,
		exitCode: 0,
		stdout:   "BAR=foo!",
	},
}

func TestCommandExecution(t *testing.T) {
	for _, d := range executions {
		command := d.command
		res := command.Run()
		if res.Success() != d.success {
			t.Fatalf("%s: expected success %t but got %t", command, d.success, res.Success())
		}
		expectedCode := d.exitCode
		actualCode := res.ExitStatus()
		if actualCode != expectedCode {
			t.Fatalf("%s: expected exit code %d but got %d", command, expectedCode, actualCode)
		}
		assertStringContains(t, res.Stdout().String(), d.stdout)
		assertStringContains(t, res.Stderr().String(), d.stderr)
	}
}

func TestWorkingDir(t *testing.T) {
	wd, err := ioutil.TempDir(os.Getenv("TMP"), "TestWorkdir")
	if err != nil {
		t.Errorf("Unexpected error creating temporary directory")
	}

	fmt.Println(wd)
	command := &Command{
		CommandLine: `echo %cd%`,
		WorkingDir:  wd,
	}
	res := command.Run()
	if !res.Success() {
		t.Fatalf("%s: expected success but got fail", command)
	}
	assertStringContains(t, res.Stdout().String(), wd)
}

func assertStringContains(t *testing.T, s string, substr string) {
	if substr != "" && !strings.Contains(s, substr) {
		t.Fatalf("expected output\n%s\n  does not contain\n%s\n", s, substr)
	}
}
