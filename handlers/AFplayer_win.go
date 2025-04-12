//go:build windows
// +build windows

package handlers

import (
	"fmt"
	"os/exec"
	"syscall"
)

// -af "atempo=2"
func (x *AFplayer) Play(fileName string) error {
	player := exec.Command("ffplay.exe", fileName)
	player.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if x.Rate != 0 {
		player = exec.Command("ffplay.exe", "-af", fmt.Sprint("atempo=", x.Rate), fileName, "-nodisp", "-autoexit")
	}
	err := player.Run()
	x.Pid = player.Process.Pid
	return err
}

func (x *AFplayer) Stop() error {
	if x.Pid != 0 {
		player := exec.Command("taskkill.exe", "/PID", fmt.Sprint(x.Pid), "/F")
		err := player.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
