package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	link := os.Args[1]
	var name string
	fmt.Print("Filename (without extension): ")
	fmt.Scanln(&name)
	exec.Command("yt-dlp", "-f", "bestaudio", "--embed-thumbnail", "--add-metadata", "--audio-format", "mp3", "-o", "C:/Users/Manzo/Downloads/"+name+".%(ext)s", link).Run()
	exec.Command("adb", "shell", "mkdir", "-p", "/sdcard/audiobooks/"+name).Run()
	exec.Command("adb", "push", "C:/Users/Manzo/Downloads/"+name+".mp3", "/sdcard/audiobooks/"+name+"/").Run()
	os.Remove("C:/Users/Manzo/Downloads/" + name + ".mp3")
}
