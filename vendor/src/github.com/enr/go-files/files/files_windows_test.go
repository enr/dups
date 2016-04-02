// +build windows

package files

import (
	"testing"
)

var isSymlinkDataWin = []maybeln{
	{"", false},
	{"   ", false},
	{"?|!", false},
	{".notfound", false},
	{".", false},
	{"testdata", false},
	{"testdata/", false},
	{"testdata/files", false},
	{"testdata/files/", false},
	{"testdata/files/01.txt", false},
	{"testdata/files/linkto01", false},
	{"testdata/files/sub", false},
	{"testdata/files/sub/", false},
}

func TestIsSymlinkWin(t *testing.T) {
	for _, data := range isSymlinkDataWin {
		is := IsSymlink(data.path)
		if is != data.isln {
			t.Errorf(`Expected IsSymlink=%t for path "%s"`, data.isln, data.path)
		}
	}
}
