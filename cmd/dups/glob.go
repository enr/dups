package main

import (
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
)

type Glob string

func init() {
	var glob Glob
	var _ pflag.Value = &glob // ensure it implements the interface
}

func (g *Glob) String() string {
	if g == nil {
		return ""
	}
	return string(*g)
}

func (g *Glob) Set(value string) error {
	if _, err := filepath.Match(value, ""); err != nil {
		return err
	}
	*g = Glob(value)
	return nil
}

func (g *Glob) Type() string {
	return "glob"
}

func (g *Glob) Match(path string) bool {
	if g == nil {
		return false
	}
	_, base := filepath.Split(path)
	// Ignoring error because it was already validated in [Glob.Set]
	ok, _ := filepath.Match(string(*g), base)
	return ok
}

type GlobArray []Glob

var _ pflag.Value = &GlobArray{} // ensure it implements the interface

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

func (g *GlobArray) Set(value string) error {
	var glob Glob
	if err := glob.Set(value); err != nil {
		return err
	}
	*g = append(*g, glob)
	return nil
}

func (g *GlobArray) Type() string {
	return "globs"
}

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
