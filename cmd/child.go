package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func Child() {

	cmd := exec.Command("/bin/ash")

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	rootfs := "/home/darshan/Projects/GO/rootfs/"

	if err := syscall.Chroot(rootfs); err != nil {
		fmt.Print("Syscall Error", err)
		return
	}

	if err := os.Chdir("/"); err != nil {
		fmt.Print("Chdir Error", err)
		return
	}

	err := syscall.Mount("proc", "/proc", "proc", 0, "")

	if err != nil {
		fmt.Print("Proc Syscall Error ", err)
		return
	}

	syscall.Sethostname([]byte("Container"))

	err = syscall.Mount("tmpfs", "/tmp", "tmpfs", 0, "")
	if err != nil {
		fmt.Print("tempfs Syscall Error ", err)
		return
	}
	err = cmd.Run()

	if err != nil {
		fmt.Print("Child run error ", err)
		return
	}

	syscall.Unmount("proc", 0)
	syscall.Unmount("tmp", 0)
}
