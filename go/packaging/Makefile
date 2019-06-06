GoCmd=go
EXEAND32=helloand32
EXELIN=hellolin
EXELIN32=hellolin32
EXEWIN=hellowin.exe
EXEWIN32=hellowin32.exe
EXEPI=hellopi
EXEPI32=hellopi32
EXEWASM=hellowasm
BUILDARGS_D=-v
BUILDARGS_R=-ldflags "-s" -v
BUILDARGS=-v

default:
	@echo 'Targets:'
	@echo ' build-all         - Build all targets (except Android)'
	@echo ' build-android-32  - Android 32 bit'
	@echo ' build-linux       - Linux 64 bit'
	@echo ' build-linux-32    - Linux 32 bit'
	@echo ' build-windows     - Windows 64 bit'
	@echo ' build-windows-32  - Windows 32 bit'
	@echo ' build-pi          - Raspberry Pi 64 bit'
	@echo ' build-pi-32       - Raspberry Pi 32 bit'
	@echo ' build-wasm        - WebAssembly'
	@echo ' analyze'
	@echo ' format'
	@echo ' clean'

run: greetinglib build
	@./$(EXE)

build-all: build-linux build-linux-32 build-windows build-windows-32 build-pi build-pi-32 build-wasm

build-android-32: $(EXEAND32)

build-linux: $(EXELIN)

build-linux-32: $(EXELIN32)

build-windows: $(EXEWIN)

build-windows-32: $(EXEWIN32)

build-pi: $(EXEPI)

build-pi-32: $(EXEPI32)

build-wasm: $(EXEWASM)

$(EXEAND32): hello.go greetinglib/greetinglib.go
	env GOOS=android GOARCH=386 $(GoCmd) build $(BUILDARGS) -o $(EXEAND32)

$(EXELIN): hello.go greetinglib/greetinglib.go
	env GOOS=linux GOARCH=amd64 $(GoCmd) build $(BUILDARGS) -o $(EXELIN)

$(EXELIN32): hello.go
	env GOOS=linux GOARCH=386 $(GoCmd) build $(BUILDARGS) -o $(EXELIN32)

$(EXEWIN): hello.go
	env GOOS=windows GOARCH=amd64 $(GoCmd) build $(BUILDARGS) -o $(EXEWIN)

$(EXEWIN32): hello.go
	env GOOS=windows GOARCH=386 $(GoCmd) build $(BUILDARGS) -o $(EXEWIN32)

$(EXEPI): hello.go
	env GOOS=linux GOARCH=arm64 $(GoCmd) build $(BUILDARGS) -o $(EXEPI)

$(EXEPI32): hello.go
	env GOOS=linux GOARCH=arm $(GoCmd) build $(BUILDARGS) -o $(EXEPI32)

$(EXEWASM): hello.go
	env GOOS=js GOARCH=wasm $(GoCmd) build $(BUILDARGS) -o $(EXEWASM)

analyze:
	$(GoCmd) vet

format:
	$(GoCmd) fmt hello.go
	$(GoCmd) fmt greetinglib/greetinglib.go

clean:
	-rm -f $(EXEAND32)
	-rm -f $(EXELIN)
	-rm -f $(EXELIN32)
	-rm -f $(EXEWIN)
	-rm -f $(EXEWIN32)
	-rm -f $(EXEPI)
	-rm -f $(EXEPI32)
	-rm -f $(EXEWASM)
	-rm -f debug
