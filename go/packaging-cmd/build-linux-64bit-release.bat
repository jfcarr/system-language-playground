SET GOOS=linux
SET GOARCH=amd64

go build -ldflags "-s -w"