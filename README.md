# Go in Practice

Collection of Go programming exercises and use cases

## Forewords

Each feature branch is a use case or collection of exercises about leveraging Go to do some tasks. The `master` branch contains minimal setup to develop Go in remote container using VSCode.

Each feature is a distinguish use case and not related to each other.

## Usage

Create folder and fetch only one branch into the folder with the same name:
```bash
# For example, in order to enhance 'feature/wma2mp3'
mkdir go-cli && cd go-cli
git init
git remote add origin https://github.com/sunshine55/python-practice.git
git fetch feature/go-cli
git checkout feature/go-cli
```
Follow README on each feature branch to setup and run project