package main

import (
	"path"

	"github.com/enr/dups/lib/colorgen"
	"github.com/jwalton/gchalk"
)

type reporter struct {
	color *colorgen.Generator
}

func (e *reporter) printFirstDup(checksum string, fp string) {
	f := fp
	if fullPath {
		f, _ = normalizePath(path.Join(baseDirectory, fp))
	}
	e.p(checksum, f)
}

func (e *reporter) printDup(checksum string, fil file) {
	f := fil.id
	if fullPath {
		f = fil.path
	}
	e.p(checksum, f)
}

func (e *reporter) p(checksum string, p string) {
	colorFn := gchalk.RGB(e.color.GenerateRGB(checksum))
	if names {
		logger.Println(colorFn(p))
	} else {
		logger.Printf("%s %s", colorFn(checksum), p)
	}
}
