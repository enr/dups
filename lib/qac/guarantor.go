package qac

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// usefult for tests
func newGuarantor(e executor) *Guarantor {
	return &Guarantor{
		executor: e,
	}
}

// NewGuarantor creates a default implementation for Guarantor.
func NewGuarantor() *Guarantor {
	return &Guarantor{
		executor: &runcmdExecutor{},
	}
}

// Guarantor checks the results respect expectations.
type Guarantor struct {
	executor executor
}

// Verify the given specification.
func (g *Guarantor) Verify(spec Spec) *SpecExecutionResult {
	command := spec.Command
	ser := &SpecExecutionResult{}
	res := g.executor.execute(command)
	verifyStatusExpectation(spec.Expectation.Status, ser, res, command.String())
	verifyOutputExpectation(spec.Expectation.Output.Stdout, ser, res.stdout, command.String())
	verifyOutputExpectation(spec.Expectation.Output.Stderr, ser, res.stderr, command.String())
	return ser
}

// VerifyConventional the given specification.
func (g *Guarantor) VerifyConventional(conventional ConventionalSpec) *SpecExecutionResult {
	return g.Verify(toFullSpec(conventional))
}

func verifyStatusExpectation(expectation StatusExpectation, ser *SpecExecutionResult, res executionResult, desc string) {
	if res.success != expectation.Success {
		ser.addError(fmt.Errorf("%s: expected success %t but got %t", desc, expectation.Success, res.success))
	}
	expectedCode := expectation.Code
	actualCode := res.exitCode
	if actualCode != expectedCode {
		ser.addError(fmt.Errorf("%s: expected exit code %d but got %d", desc, expectedCode, actualCode))
	}
}

func verifyOutputExpectation(expectation OutputExpectation, ser *SpecExecutionResult, out string, desc string) {
	if expectation.File != "" {
		content, err := ioutil.ReadFile(expectation.File)
		if err != nil {
			ser.addError(err)
			return
		}
		// Convert []byte to string and print to screen
		t := strings.TrimSpace(string(content))
		if out != t {
			ser.addError(fmt.Errorf("%s: actual output\n_%s_\ndoes not contain:\n_%s_", desc, out, t))
		}
	}
	if len(expectation.Tokens) > 0 {
		for _, t := range expectation.Tokens {
			if !strings.Contains(out, t) {
				ser.addError(fmt.Errorf("%s: actual output\n%s\ndoes not contain:\n%s", desc, out, t))
			}
		}
	}
}
