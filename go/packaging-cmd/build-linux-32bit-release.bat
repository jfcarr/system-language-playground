SET GOOS=linux
SET GOARCH=386

go build -ldflags "-s -w"
