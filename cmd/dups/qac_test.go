package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/enr/dups/lib/qac"
)

func TestCommandExecution2(t *testing.T) {
	guarantor := qac.NewGuarantor()
	for _, spec := range specs {
		result := guarantor.VerifyConventional(spec)
		if len(result.Errors()) > 0 {
			t.Errorf("QAC errors %v", result.Errors())
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
