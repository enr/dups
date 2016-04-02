// +build darwin freebsd linux netbsd openbsd

package files

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestExistsPerms(t *testing.T) {
	startDir, err := ioutil.TempDir("/tmp", "filestest_existsperms")
	check(t, err)
	base := startDir + "/TestExistsPerms"
	err = os.MkdirAll(base, 0777)
	check(t, err)

	fileName := base + "/bar"
	f, err := os.OpenFile(fileName, os.O_CREATE, 0660)
	check(t, err)
	defer f.Close()

	execCommand(t, "/bin/chmod", "777", fileName)
	execCommand(t, "/bin/chmod", "777", base)

	e := Exists(fileName)
	if !e {
		t.Errorf(`%s : expected exists but got false`, fileName)
	}

	execCommand(t, "/bin/chmod", "000", fileName)
	execCommand(t, "/bin/chmod", "000", base)

	e2 := Exists(fileName)
	if !e2 {
		t.Errorf(`%s : expected exists but got false`, fileName)
	}
}

func TestIsAccessible(t *testing.T) {
	startDir, err := ioutil.TempDir("/tmp", "filestest_isaccessible")
	check(t, err)
	base := startDir + "/TestIsAccessible"
	err = os.MkdirAll(base, 0777)
	check(t, err)

	fileName := base + "/bar"
	f, err := os.OpenFile(fileName, os.O_CREATE, 0660)
	check(t, err)
	defer f.Close()

	execCommand(t, "/bin/chmod", "777", fileName)
	execCommand(t, "/bin/chmod", "777", base)
	a1 := IsAccessible(fileName)
	if !a1 {
		t.Errorf(`%s : expected accessible but got false`, fileName)
	}

	execCommand(t, "/bin/chmod", "000", fileName)
	execCommand(t, "/bin/chmod", "000", base)

	a2 := IsAccessible(fileName)
	if a2 {
		t.Errorf(`%s : expected not accessible but got true`, fileName)
	}
}

var isSymlinkData = []maybeln{
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
	{"testdata/files/linkto01", true},
	{"testdata/files/sub", false},
	{"testdata/files/sub/", false},
}

func TestIsSymlink(t *testing.T) {
	for _, data := range isSymlinkData {
		is := IsSymlink(data.path)
		if is != data.isln {
			t.Errorf(`Expected IsSymlink=%t for path "%s"`, data.isln, data.path)
		}
	}
}
