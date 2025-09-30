package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Print("Enter the path to the file containing YouTube URLs: ")
	var filePath string
	fmt.Scanln(&filePath)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.Contains(line, "https://www.youtube.com/") || strings.Contains(line, "https://youtu.be/") {
			lines = append(lines, line)
		}
	}

	if len(lines) == 0 {
		fmt.Println("No YouTube URLs found in the file.")
		return
	}

	downloadDir := filepath.Join(filepath.Dir(filePath), "output")
	if err := os.MkdirAll(downloadDir, 0777); err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}

	for _, line := range lines {
		fmt.Println("Processing:", line)
		parts := strings.SplitN(line, "@", 2)
		title := parts[0] + ".mp3"
		url := parts[1]
		cmd := exec.Command("yt-dlp", "-x", "--audio-format", "mp3", "-o", filepath.Join(downloadDir, title), url)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("Error processing", url, ":", err)
		}
	}

	fmt.Println("Audio files are extracted to", downloadDir)
}
