package main

import (
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
)

// Glob is a pattern for matching on paths.
type Glob string

func init() {
	var glob Glob
	var _ pflag.Value = &glob // ensure it implements the interface
}

// String implements [pflag.Value] and [fmt.Stringer].
func (g *Glob) String() string {
	if g == nil {
		return ""
	}
	return string(*g)
}

// Set parses and updates this glob pattern, or returns error on invalid pattern.
// Implements [pflag.Value].
func (g *Glob) Set(value string) error {
	if _, err := filepath.Match(value, ""); err != nil {
		return err
	}
	*g = Glob(value)
	return nil
}

// Type returns this [pflag.Value] type name. Used in CLI output.
func (g *Glob) Type() string {
	return "glob"
}

// Match tries to compare the glob pattern with the given path and
// returns true if the pattern matches the path.
func (g *Glob) Match(path string) bool {
	if g == nil {
		return false
	}
	_, base := filepath.Split(path)
	// Ignoring error because it was already validated in [Glob.Set]
	ok, _ := filepath.Match(string(*g), base)
	return ok
}

// GlobArray is a slice of patterns for matching on paths.
// It implements [pflag.Value] and can therefore be registered as a flag.
// Setting this flag multiple times just appends more patterns to the same slice.
type GlobArray []Glob

var _ pflag.Value = &GlobArray{} // ensure it implements the interface

// String implements [pflag.Value] and [fmt.Stringer].
func (g *GlobArray) String() string {
	if g == nil {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteByte('[')
	for i, v := range *g {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(string(v))
	}
	sb.WriteByte(']')
	return sb.String()
}

// Set parses and appends to this glob pattern slice, or returns error on
// invalid pattern. Implements [pflag.Value].
func (g *GlobArray) Set(value string) error {
	var glob Glob
	if err := glob.Set(value); err != nil {
		return err
	}
	*g = append(*g, glob)
	return nil
}

// Type returns this [pflag.Value] type name. Used in CLI output.
func (g *GlobArray) Type() string {
	return "globs"
}

// Match tries to compare all of its glob patterns with the given path and
// returns true if any of the patterns matches the path.
func (g *GlobArray) Match(path string) bool {
	if g == nil {
		return false
	}
	for _, glob := range *g {
		if glob.Match(path) {
			return true
		}
	}
	return false
}
