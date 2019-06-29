package main

import (
	"fmt"

	"github.com/kassybas/shell-exec/exec"
)

func main() {
	opts := exec.Options{
		Silent:          false,       // output to stdout and stderr
		IgnoreResult:    false,       // return stdout, stderr and status code
		ShieldEnv:       false,       // expose the current process' environment variables for the shell process
		ShellPath:       "/bin/bash", // shell to be executed (default: sh)
		ShellExtraFlags: []string{"--posix"},
	}

	script := `
		echo hello ${FOO}
		echo hi >&2
		exit 11
	`
	envVars := []string{
		"FOO=world",
	}

	so, se, src, _ := exec.ShellExec(script, envVars, opts)
	fmt.Println("--RESULT:")
	fmt.Println("OUT:", so)
	fmt.Println("ERR:", se)
	fmt.Println("RC:", src)

}
