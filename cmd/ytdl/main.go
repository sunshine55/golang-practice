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
	var urls []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "https://www.youtube.com/") || strings.HasPrefix(line, "https://youtu.be/") {
			urls = append(urls, line)
		}
	}

	if len(urls) == 0 {
		fmt.Println("No YouTube URLs found in the file.")
		return
	}

	downloadDir := filepath.Join("../../assets", "output")
	for _, url := range urls {
		fmt.Println("Processing:", url)
		parts := strings.SplitN(url, "@", 2)
		source := parts[0]
		title := parts[1] + ".mp3"
		// Download video and extract audio using yt-dlp
		cmd := exec.Command("yt-dlp", "-x", "--audio-format", "mp3", "-o", filepath.Join(downloadDir, title), source)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("Error processing", url, ":", err)
		}
	}

	fmt.Println("All audio files saved to", downloadDir)
}
