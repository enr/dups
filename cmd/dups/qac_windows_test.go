//go:build windows
// +build windows

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
		Stderr: []string{
			"Error: empty search path",
		},
	},
	// error no dir
	{
		CommandExe:  "../../bin/dups",
		CommandArgs: []string{"F:/this/directory/does/not/exist"},
		Success:     false,
		ExitCode:    1,
		Stderr: []string{
			"Error reading F:/this/directory/does/not/exist",
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
		CommandArgs: []string{abs("../../testdata/01")},
		WorkingDir:  "../../bin",
		Success:     true,
		ExitCode:    0,
		Stdout:      []string{"855426068ee8939df6bce2c2c4b1e7346532a133 sub/010.txt"},
	},
	{
		CommandExe:  "../../bin/dups",
		CommandArgs: []string{"--names-only", abs("../../testdata/01")},
		Success:     true,
		ExitCode:    0,
		Stdout:      []string{"sub/010.txt"},
	},
	{
		CommandExe:  "../../bin/dups",
		CommandArgs: []string{"--dups-exit", abs("../../testdata/01")},
		WorkingDir:  "../../bin",
		Success:     false,
		ExitCode:    2,
		Stdout:      []string{"855426068ee8939df6bce2c2c4b1e7346532a133 sub/010.txt"},
	},
	{
		CommandExe:  "../../bin/dups",
		CommandArgs: []string{"--names-only", "--dups-exit", abs("../../testdata/01")},
		Success:     false,
		ExitCode:    2,
		Stdout:      []string{"sub/010.txt"},
	},
}
