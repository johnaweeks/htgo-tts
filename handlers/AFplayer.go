//go:build darwin

package handlers

import (
	"fmt"
	"os/exec"
)

type AFplayer struct {
	Rate float32
}

func (x *AFplayer) Play(fileName string) error {
	player := exec.Command("afplay", fileName)
	if x.Rate != 0 {
		player = exec.Command("afplay", "--rate", fmt.Sprint(x.Rate), fileName)
	}
	return player.Run()
}
