package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/SakuraBurst/spotify_reloader/process"
)

func main() {
	processes := process.FindProcessesByName("Spotify.exe")
	procPath := processes[0].PPath
	for ind := range processes {
		processes[ind].DeleteProcess()
	}
	fmt.Println(procPath)
	cmd := exec.Command(procPath)
	// cmd.Run()
	cmd.Start()
	log.Fatal()
}
