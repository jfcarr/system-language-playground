SET GOOS=windows
SET GOARCH=386

go build -ldflags "-s -w"
