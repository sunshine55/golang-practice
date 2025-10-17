package main

import (
	"fmt"
	"log"

	"google.golang.org/api/drive/v3"

	"github.com/sunshine55/golang-practice/go-cli/ytdl/internal/app"
)

func executeUpload(srv *drive.Service) {
	r, err := srv.Files.List().PageSize(10).
		Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	fmt.Println("Files:")
	if len(r.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, i := range r.Files {
			fmt.Printf("%s (%s)\n", i.Name, i.Id)
		}
	}
}

func main() {
	input := app.AskInput()

	app.Parse(input)

	var srv *drive.Service
	switch input.AuthMethod {
	case "1":
		srv = app.InitServiceAccount()
		fmt.Println("Google Drive service initialized with Service Account")
	case "2":
		fmt.Println("Google Drive service initialized with Client ID (work in progress)")
		return
	default:
		fmt.Println("Invalid authentication method selected")
		return
	}

	executeUpload(srv)
}
