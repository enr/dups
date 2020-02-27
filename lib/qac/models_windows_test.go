// +build windows

package qac

import (
	"strings"
	"testing"
)

func TestConventionalToFullSpec(t *testing.T) {
	conventional := ConventionalSpec{
		CommandExe: "testdata/executable.cmd",
	}
	spec := toFullSpec(conventional)
	if !strings.Contains(spec.Command.Exe, "testdata/executable") {
		t.Errorf("Expected command exe containing <testdata/executable> but got <%s>", spec.Command.Exe)
	}
}
