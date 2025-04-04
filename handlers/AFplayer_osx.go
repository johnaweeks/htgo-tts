//go:build darwin

package handlers

import (
	"fmt"
	"os/exec"
)

func (x *AFplayer) Play(fileName string) error {
	player := exec.Command("afplay", fileName)
	if x.Rate != 0 {
		player = exec.Command("afplay", "--rate", fmt.Sprint(x.Rate), fileName)
	}
	err := player.Run()
	x.Pid = player.Process.Pid
	return err
}

func (x *AFplayer) Stop() error {
	if x.Pid != 0 {
		player := exec.Command("kill", fmt.Sprint(x.Pid))
		err := player.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
