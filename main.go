package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	link := os.Args[1]
	var name string
	fmt.Print("Filename: ")
	fmt.Scanf("%s", &name)
	exec.Command("yt-dlp", "-f", "bestaudio", "--embed-thumbnail", "--add-metadata", "--audio-format", "mp3", "-o", "C:/Users/Manzo/Downloads/"+name+".%(ext)s", link).Run()
	exec.Command("adb", "shell", "mkdir", "-p", "/sdcard/audiobooks/"+name).Run()
	files, _ := filepath.Glob("C:/Users/Manzo/Downloads/" + name + ".*")
	for _, file := range files {
		exec.Command("adb", "push", file, "/sdcard/audiobooks/"+name+"/").Run()
		os.Remove(file)
	}
}
