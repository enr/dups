// +build darwin freebsd linux netbsd openbsd

package main

import (
	"github.com/enr/dups/lib/qac"
)

var specs = []qac.ConventionalSpec{
	// error no args
	{
		CommandExe:  "../../bin/dups",
		CommandArgs: []string{},
		Success:     false,
		ExitCode:    1,
		Stdout: []string{
			"error missing path",
		},
	},
	// version
	{
		CommandExe:  "../../bin/dups",
		CommandArgs: []string{"--version"},
		Success:     true,
		ExitCode:    0,
		Stdout: []string{"dups",
			"Revision", "Build date"},
	},
	{
		CommandExe:  "../../bin/dups",
		CommandArgs: []string{"../testdata/01"},
		WorkingDir:  "../../bin",
		Success:     false,
		ExitCode:    2,
		Stdout:      []string{"f1d2d2f924e986ac86fdf7b36c94bcdf32beec15 sub/010.txt"},
	},

	{
		CommandExe:  "../../bin/dups",
		CommandArgs: []string{"--names-only", "../../testdata/01"},
		Success:     false,
		ExitCode:    2,
		Stdout:      []string{"sub/010.txt"},
	},
	{
		CommandExe:  "../../bin/dups",
		CommandArgs: []string{"--full-path", "../../testdata/01"},
		Success:     false,
		ExitCode:    2,
		Stdout:      []string{"/testdata/01/sub/010.txt"},
	},
}
