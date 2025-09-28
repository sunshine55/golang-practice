# Go in Practice

Collection of Go programming exercises and use cases

## Forewords

Each feature branch is a use case or collection of exercises about leveraging Go to do some tasks. They're not related to each other.

The `master` branch contains minimal setup to start developing Go app in remote container using VSCode.

## Usage

Create folder and fetch only one branch into the folder with the same name:
```bash
# For example, in order to enhance 'feature/go-cli'
mkdir go-cli && cd go-cli
git init
git remote add origin https://github.com/sunshine55/golang-practice.git
git fetch feature/go-cli
git checkout feature/go-cli
```
Follow README on each feature branch to setup and run project