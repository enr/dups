package qac

import (
	"fmt"
	"strings"
)

// Comparison describe how to compare values.
type Comparison int

const (
	unknown Comparison = iota
	// Exact comparison
	Exact
	// ContainsAll used for output tokens
	ContainsAll
	// ContainsAny used for output tokens
	ContainsAny
	sentinel
)

// UnmarshalJSON manages the unmarshalling from string to actual Comparison type.
func (o *Comparison) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)
	*o = toComparison(str)
	return nil
}

func toComparison(s string) Comparison {
	if "EXACT" == s {
		return Exact
	}
	if "CONTAINS_ALL" == s {
		return ContainsAll

	}
	if "CONTAINS_ANY" == s {
		return ContainsAny
	}
	return unknown
}

// Command is the command to execute
type Command struct {
	WorkingDir string `json:"working_dir"`
	// the command is run in a shell, that is prepend `/bin/sh -c` or `cmd /C` to the command line
	Cli string `json:"cli"`
	// path to executable, relative form the working dir
	Exe  string   `json:"exe"`
	Args []string `json:"args"`
}

func (c Command) String() string {
	fullCommand := c.Cli
	if fullCommand == "" {
		fullCommand = strings.TrimSpace(c.Exe + " " + strings.Join(c.Args, " "))
	}
	return fmt.Sprintf("%s# %s", c.WorkingDir, fullCommand)
}

// StatusExpectation is an expectation about the status of the executed command.
type StatusExpectation struct {
	Success    bool       `json:"success"`
	Code       int        `json:"code"`
	Comparison Comparison `json:"comparison"`
}

// OutputExpectation is an expectation about the output of the executed command.
type OutputExpectation struct {
	Tokens     []string   `json:"tokens"`
	File       string     `json:"file"`
	Comparison Comparison `json:"comparison"`
}

// OutputExpectations is the container for all output expectations.
type OutputExpectations struct {
	Stdout OutputExpectation `json:"stdout"`
	Stderr OutputExpectation `json:"stderr"`
	//Combined OutputExpectation `json:"combined"`
}

// Expectation is the main container for expectations about status and output.
type Expectation struct {
	Status StatusExpectation  `json:"status"`
	Output OutputExpectations `json:"output"`
}

// Spec contains the command to run and expectations about the status and output.
type Spec struct {
	Command     Command     `json:"command"`
	Expectation Expectation `json:"expectation"`
}

// SpecExecutionResult is the container of verification results.
type SpecExecutionResult struct {
	errors []error
}

func (r *SpecExecutionResult) addError(err error) {
	r.errors = append(r.errors, err)
}

// Errors returns the errors list.
func (r *SpecExecutionResult) Errors() []error {
	return r.errors
}

// ConventionalSpec is a shortcut for creating Spec.
type ConventionalSpec struct {
	// path to the exe: it will be normalized and added of the extension
	CommandExe string
	// command args
	CommandArgs []string
	// the woking directory
	WorkingDir string
	// expected succes
	Success bool
	// expected exit code
	ExitCode int
	// tokens expected to be present in the stdout
	Stdout []string
	// tokens expected to be present in the stderr
	Stderr []string
}

func toFullSpec(conventional ConventionalSpec) Spec {
	return Spec{

		Command: Command{
			Exe:        FullExePathOrFail(conventional.CommandExe),
			Args:       conventional.CommandArgs,
			WorkingDir: conventional.WorkingDir,
		},
		Expectation: Expectation{
			Status: StatusExpectation{
				Success: conventional.Success,
				Code:    conventional.ExitCode,
			},
			Output: OutputExpectations{
				Stdout: OutputExpectation{
					Tokens:     conventional.Stdout,
					Comparison: Exact,
				},
				Stderr: OutputExpectation{
					Tokens:     conventional.Stderr,
					Comparison: Exact,
				},
			},
		},
	}

}
