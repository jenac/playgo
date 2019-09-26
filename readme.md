# Environment Setup

## Install golang
```
sudo apt install golang
```

## Set Environment Variables
```
mkdir ~/go
export GOPATH=/home/lihe/go
export PATH=$PATH:$GOPATH/bin
```

## Folder Structure
```
mkdir -p ~/go
mkdir -p ~/go/src
mkdir -p ~/go/pkg
mkdir -p ~/go/bin
```

## Put your source repo
for example, https://github.com/jenac/playgo.git
```
mkdir -p ~/go/src/github.com/jenac
cd ~/go/src/github.com/jenac
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
