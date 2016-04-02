// +build darwin freebsd linux netbsd openbsd

package runcmd

var testdata2 = []cnl{
	{command: &Command{
		Logfile: "out.log",
	},
		name:    "",
		logfile: "out.log",
	},
	{command: &Command{
		CommandLine: `echo "home=$HOME"`,
		UseEnv:      true,
	},
		name:    "echo-home-home",
		logfile: "runcmd-echo-home-home.log",
	},
	{command: &Command{
		Exe:    `/usr/local/bin/myapp`,
		Args:   []string{"-a", `"say hello!"`},
		UseEnv: true,
	},
		name:    "usr-local-bin-myapp-a-say-hello",
		logfile: "runcmd-usr-local-bin-myapp-a-say-hello.log",
	},
}
