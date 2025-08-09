package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func Run(args []string) {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, args...)...)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID | syscall.CLONE_NEWUTS | syscall.CLONE_NEWNS,
	}

	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()

	if err != nil {
		fmt.Print(err)
	}
}
