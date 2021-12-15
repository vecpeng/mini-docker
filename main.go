package main

import (
	"os/exec"
	"syscall"
	"os"
	"log"
)

func main() {
	// 指定fork出的新进程内的初始命令，默认使用sh
	cmd := exec.Command("sh")

	// 设置系统调用参数
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}