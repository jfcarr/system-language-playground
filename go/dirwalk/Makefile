default:
	@echo 'Targets:'
	@echo '  run'
	@echo '  build'
	@echo '  format'
	@echo '  clean'

run: build
	./dirwalk

build: dirwalk

dirwalk: dirwalk.go
	go build

format:
	go fmt

clean:
	go clean
