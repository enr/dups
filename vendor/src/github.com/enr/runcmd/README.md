Runcmd
======

[![Build Status](https://travis-ci.org/enr/runcmd.png?branch=master)](https://travis-ci.org/enr/runcmd)
[![Build status](https://ci.appveyor.com/api/projects/status/cklfbhqkoi2356if?svg=true)](https://ci.appveyor.com/project/enr/runcmd)

Should be a Go library to execute external commands.

**Nothing to see here now.**

Import the library:

```Go
import (
    "github.com/enr/runcmd"
)
```
You can use this library in two ways.

Run a command, wait for it to complete and get a result:

```Go
executable := "/usr/bin/ls"
args := []string{"-al"}
command := &runcmd.Command{
    Exe:  executable,
    Args: args,
}
res := command.Run()
if res.Success() {
    fmt.Printf("standard output: %s", res.Stdout().String())
} else {
    fmt.Printf("error executing %s. Exit code %d", command, res.ExitStatus())
    fmt.Printf("error output: %s", res.Stderr().String())
    fmt.Printf("the error: %v", res.Error())
}
```

Start a command as a process. In Unix systems this process will survive to the parent.

```Go
executable := "/usr/local/bin/start-server"
command := &runcmd.Command{
    Exe:  executable,
}
logFile := cmd.GetLogfile()
// maybe you want to follow logs...
t, _ := tail.TailFile(logFile, tail.Config{Follow: true})
go func() {
    cmd.Start()
}()
```

License
-------

Mozilla Public License Version 2.0 - see LICENSE file.

Copyright 2015 runcmd contributors
