// +build windows

package runcmd

func (c *Command) useShell() {
	command := c.CommandLine
	c.Exe = "cmd"
	c.Args = []string{"/C", command}
}

// shellAndArgs is a helper function that returns an OS specific
// shell and arguments for that particular shell
func shellAndArgs() (string, []string) {
	var com []string
	com = []string{
		"powershell.exe",
		"-noprofile",
		"-noninteractive",
		"-command",
		"try{$input|iex; exit $LastExitCode}catch{Write-Error -Message $Error[0]; exit 1}",
	}
	return com[0], com[1:]
}
