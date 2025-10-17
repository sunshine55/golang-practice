package app

import (
	"fmt"
)

type Input struct {
	PlayList   string
	AuthMethod string
}

func AskInput() Input {
	fmt.Print("Enter the path to the file containing YouTube URLs\n> ")
	var playlist string
	fmt.Scanln(&playlist)

	fmt.Print("Select authentication method\n1. service_account\n2. client_id\n> ")
	var authMethod string
	fmt.Scanln(&authMethod)

	return Input{
		PlayList:   playlist,
		AuthMethod: authMethod,
	}
}
