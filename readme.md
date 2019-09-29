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
go test algorithms/qsort
go test algorithms/bubblesort
go run sorter.go
```

## Run molayer.go
```
cd src/app
go test library
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

