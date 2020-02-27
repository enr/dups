package qac

import (
	"fmt"
	"strings"
	"testing"
)

type fixedValueExecutor struct {
	success  bool
	exitCode int
	stdout   string
	stderr   string
}

func (e *fixedValueExecutor) execute(c Command) executionResult {
	return executionResult{
		success:  e.success,
		exitCode: e.exitCode,
		stdout:   e.stdout,
		stderr:   e.stderr,
	}
}

func TestSpecificationError(t *testing.T) {
	e := &fixedValueExecutor{
		success:  true,
		exitCode: 0,
		stdout:   "stdout",
		stderr:   "stderr",
	}
	sut := newGuarantor(e)

	/*
	 *  A Spec with wrong values.
	 */
	spec := Spec{
		Command: Command{
			Exe:  "test",
			Args: []string{},
		},
		Expectation: Expectation{
			Status: StatusExpectation{
				Success: false,
				Code:    6,
			},
			Output: OutputExpectations{
				Stderr: OutputExpectation{
					Tokens: []string{
						"qwerty",
					},
					Comparison: Exact,
				},
			},
		},
	}

	res := sut.Verify(spec)
	// for i, err := range res.Errors() {
	// 	fmt.Printf("%d e %v\n", i, err)
	// }
	if !atLeastOneErrorContaining(res.Errors(), "qwerty") {
		t.Errorf("Expected at least one error containing <%s>", "qwerty")
	}
	if !atLeastOneErrorContaining(res.Errors(), "expected success") {
		t.Errorf("Expected at least one error containing <%s>", "expected success")
	}
	if !atLeastOneErrorContaining(res.Errors(), "expected exit code") {
		t.Errorf("Expected at least one error containing <%s>", "expected exit code")
	}
	if !atLeastOneErrorContaining(res.Errors(), "actual output") {
		t.Errorf("Expected at least one error containing <%s>", "actual output")
	}
	if len(res.Errors()) != 3 {
		t.Errorf(`Expected 3 errors, got %d`, len(res.Errors()))
	}
}

func TestSpecificationOk(t *testing.T) {
	e := &fixedValueExecutor{
		success:  true,
		exitCode: 0,
		stdout:   "stdout",
		stderr:   "stderr",
	}
	sut := newGuarantor(e)

	/*
	 *  A Spec with wrong values.
	 */
	var spec = Spec{
		Command: Command{
			Exe:  "test",
			Args: []string{},
		},
		Expectation: Expectation{
			Status: StatusExpectation{
				Success: true,
				Code:    0,
			},
			Output: OutputExpectations{
				Stdout: OutputExpectation{
					File:       "testdata/stdout.txt",
					Comparison: Exact,
				},
				Stderr: OutputExpectation{
					Tokens: []string{
						"stderr",
					},
					Comparison: Exact,
				},
			},
		},
	}

	res := sut.Verify(spec)
	if len(res.Errors()) != 0 {
		for i, err := range res.Errors() {
			fmt.Printf("%d e %v\n", i, err)
		}
		t.Errorf(`Expected 0 errors, got %d`, len(res.Errors()))
	}
}

func atLeastOneErrorContaining(errors []error, expected string) bool {
	for _, err := range errors {
		if strings.Contains(err.Error(), expected) {
			return true
		}
	}
	return false
}
