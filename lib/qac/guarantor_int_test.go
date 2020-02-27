// +build darwin freebsd linux netbsd openbsd

package qac

import (
	"fmt"
	"testing"
)

func TestSpecificationResult(t *testing.T) {
	// default guarantor run the command
	sut := NewGuarantor()

	/*
	 *  A Spec with wrong values.
	 */
	spec := Spec{
		Command: Command{
			Exe:  "/bin/bash",
			Args: []string{"-c", "testdata/exit-3"},
		},
		Expectation: Expectation{
			Status: StatusExpectation{
				Success: false,
				Code:    3,
			},
			Output: OutputExpectations{
				Stdout: OutputExpectation{
					Tokens: []string{
						"stdout exit-3",
					},
					Comparison: Exact,
				},
				Stderr: OutputExpectation{
					Tokens: []string{
						"stderr exit-3",
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
