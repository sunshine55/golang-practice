FROM golang:1.25-bookworm

RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y --no-install-recommends ffmpeg pipx && \
    rm -rf /var/lib/apt/lists/*
RUN pipx ensurepath && \
    pipx install yt-dlp