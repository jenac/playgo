# Environment Setup

## Install golang
```
sudo apt install golang
```

## Set Environment Variables
```
mkdir ~/Projects/go
export GOPATH=/home/lihe/Projects/go
export PATH=$PATH:$GOPATH/bin
```

## Folder Structure
```
mkdir -p ~/Projects/go
mkdir -p ~/Projects/go/src
mkdir -p ~/Projects/go/pkg
mkdir -p ~/Projects/go/bin
```

## Put your source repo
for example, https://github.com/jenac/playgo.git
```
mkdir -p ~/Projects/go/src/github.com/jenac
cd ~/Projects/go/src/github.com/jenac
git clone https://github.com/jenac/playgo.git
cd playgo
```
work in `playgo` folder

## Install All Project Dependencies
```
go get ./...
```

## Misc comamnd
```
go get <package name>
go build
go install
```
