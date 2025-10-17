package main

import (
	"context"
	"log"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func InitServiceAccount() *drive.Service {
	ctx := context.Background()
	srv, err := drive.NewService(ctx, option.WithCredentialsFile("credentials/service_account.json"))
	if err != nil {
		log.Fatalf("Unable to create Drive client: %v", err)
	}
	return srv
}
