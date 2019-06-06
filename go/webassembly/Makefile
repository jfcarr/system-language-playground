GoCmd=go
BUILDARGS=-v

default:
	@echo 'Targets:'
	@echo ' build-wasm        - Build webassembly'
	@echo ' analyze'
	@echo ' format'
	@echo ' clean'

build-wasm:
	env GOOS=js GOARCH=wasm $(GoCmd) build $(BUILDARGS) -o main.wasm

analyze:
	$(GoCmd) vet

format:
	$(GoCmd) fmt hello.go

clean:
	-rm -f main.wasm
