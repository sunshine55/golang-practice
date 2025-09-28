# Go CLI App

## Setup

This is checked out from master branch and adjust the name of the container name to `go-cli` in `.devcontainer.json` and `compose.yml` so that Docker Compose will isolate a new dev env and VSCode can start to connect to new container.

Implement a Dockerfile that will provision the container (i.e.: OS-level media processing libs, guest/host permissions...)

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

Output:
1. Store the audios at local folder
2. Upload the audios to personal cloud

### Image Generator

***Work In Progress...***