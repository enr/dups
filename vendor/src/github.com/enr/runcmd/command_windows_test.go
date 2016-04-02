// +build windows

package runcmd

var testdata2 = []cnl{
	{command: &Command{
		Logfile: "out.log",
	},
		name:    "",
		logfile: "out.log",
	},
	{command: &Command{
		CommandLine: `echo "home=%HOME%"`,
		UseEnv:      true,
	},
		name:    "echo-home-home",
		logfile: "runcmd-echo-home-home.log",
	},
	{command: &Command{
		Exe:    `c:\bin\myapp.exe`,
		Args:   []string{"-a", `"say hello!"`},
		UseEnv: true,
	},
		name:    "c-bin-myapp-exe-a-say-hello",
		logfile: "runcmd-c-bin-myapp-exe-a-say-hello.log",
	},
}
