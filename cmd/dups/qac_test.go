package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/enr/dups/lib/qac"
)

func TestCommandExecution2(t *testing.T) {
	guarantor := qac.NewGuarantor()
	for idx, spec := range specs {
		s := fmt.Sprintf("QAC errors in spec %d:\n", idx)
		result := guarantor.VerifyConventional(spec)
		if len(result.Errors()) > 0 {
			for _, e := range result.Errors() {
				s = s + fmt.Sprintf("- %s\n", e.Error())
			}
			t.Error(s)
		}
	}
}

func cwd(f string) string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.ToSlash(path.Join(dir, f))
}
