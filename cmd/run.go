package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func Run() {
	cmd := exec.Command("/bin/ash")

	rootfs := "<rootfs-path>"

	if err := syscall.Chroot(rootfs); err != nil {
		fmt.Print("Syscall Error", err)
		return
	}

	if err := os.Chdir("/"); err != nil {

		fmt.Print("Chdir Error", err)
		return
	}
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID | syscall.CLONE_NEWUTS,
	}

	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()

	if err != nil {
		fmt.Print(err)
	}
}
