package main

import (
	"fmt"

	"github.com/kassybas/shell-exec/exec"
)

func main() {
	opts := exec.Options{
		Silent:       false,       // output to stdout and stderr
		IgnoreResult: false,       // return stdout, stderr and status code
		ShieldEnv:    false,       // expose the current process' environment variables for the shell process
		ShellPath:    "/bin/bash", // path of the shell to be executed
	}

	script := `
		echo hello ${FOO}
		echo hi >&2
		exit 11
	`
	so, se, src, _ := exec.ShellExec(script, []string{"FOO=world"}, opts)
	fmt.Println("--RESULT:")
	fmt.Println("OUT:", so)
	fmt.Println("ERR:", se)
	fmt.Println("RC:", src)

}
