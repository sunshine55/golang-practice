// To run this package (include all files in the directory):
//   cd /workspace/cmd/ytdl
//   go run .
// or
//   go run *.go

package main

import (
	"fmt"
	"log"

	"google.golang.org/api/drive/v3"
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
	input := AskInput()

	Parse(input)

	var srv *drive.Service
	switch input.authMethod {
	case "1":
		srv = InitServiceAccount()
		fmt.Println("Google Drive service initialized with Service Account:")
	case "2":
		srv = InitClientID()
		fmt.Println("Google Drive service initialized with Client ID:")
	default:
		fmt.Println("Invalid authentication method selected.")
		return
	}

	executeUpload(srv)
}
