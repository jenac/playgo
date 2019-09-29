# Environment Setup

## Install golang
```
sudo apt install golang
```
## Clone repo and go to work foler

## Set Environment Variables
```
export GOPATH=$(pwd)
echo $GOPATH
```

## Run sorter.go
```
cd src/app
go run sorter.go
```

## Install All Project Dependencies
```
go get ./...
```

## Misc comamnd
```
go get <package name>
go build
go install
go run [somefile].go
```

