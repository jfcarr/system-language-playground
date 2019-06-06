default:
	@echo 'Targets:'
	@echo ' run'
	@echo ' build'
	@echo ' analyze'
	@echo ' format'
	@echo ' clean'

run: build
	@./xmlparse

build: xmlparse

xmlparse: xmlparse.go
	go build -o xmlparse

analyze:
	go vet

format:
	go fmt

clean:
	-rm -f xmlparse
