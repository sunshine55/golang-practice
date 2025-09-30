# Go CLI App

## Setup

This is checked out from master branch and adjust the name of the container name to `go-cli` in `.devcontainer.json` and `compose.yml` so that Docker Compose will isolate a new dev env and VSCode can start to connect to new container.

Implement a Dockerfile that will provision the container (i.e.: OS-level media processing libs, guest/host permissions...)

First time: `docker compose -d --build`

Next time: `docker compose start`

## Rebuild

Destroy old service: `docker compose down --remove-orphans`

Rebuild when Dockerfile changes: `docker compose up -d --build`

## Requirements

Implment a few CLI programs in Go language

### Audio Retriever

Input:
1. Ask for the file path input
2. Check if the file content contains newline delimited URLs
3. Delimit the URLs and read them

Processing for each URL:
1. Download the video
2. Separate audio from video

Output: store the audios at local folder

### Personal Cloud Uploader

Input:
1. Ask for the directory or file input
2. Ask for personal cloud credentials

Processing:
1. Check availability of the file(s) from first input
2. Programmatically log in to personal cloud
3. Upload the files to to the cloud

Output: clean up the files at local folder if success