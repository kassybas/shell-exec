# Shell Exec

Small Go library for easy execution of shell scripts.

Based on article: https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html

See exec/exec.go for details

## Example

``` go
  opts := exec.Options{
    Silent:       false, // if false: streams to stdout and stderr
    IgnoreResult: false, // if false: returns stdout, stderr and status code
    ShieldEnv:    false, // if false: exposes the current process' environment variables for the shell process

    ShellPath:    "/bin/bash" // shell to be executed (default: sh)
    ShellExtraFlags: []string{"--posix"}, // extra options passed to the shell
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
  fmt.Println("OUT:", so)  // hello world
  fmt.Println("ERR:", se)  // hi
  fmt.Println("RC:", src)  // 11
```
