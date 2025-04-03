//go:build darwin

package handlers

import (
	"os/exec"
)

type AFplayer struct{}

func (AFplayer *AFplayer) Play(fileName string) error {
	player := exec.Command("afplay", fileName)
	return player.Run()
}
