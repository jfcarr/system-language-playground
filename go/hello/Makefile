GoCmd=go
EXE=hello
EXE32=hello32
EXEWIN=hellowin.exe
EXEWIN32=hellowin32.exe
BUILDARGS_D=-v
BUILDARGS_R=-ldflags "-s" -v
BUILDARGS=-v

default:
	@echo 'Targets:'
	@echo ' build       - Linux 64 bit'
	@echo ' build32     - Linux 32 bit'
	@echo ' buildwin    - Windows 64 bit'
	@echo ' buildwin32  - Windows 32 bit'
	@echo ' analyze'
	@echo ' format'
	@echo ' clean'

run: build
	@./$(EXE)

buildall: build build32 buildwin buildwin32

build: $(EXE)

build32: $(EXE32)

buildwin: $(EXEWIN)

buildwin32: $(EXEWIN32)

$(EXE): hello.go
	env GOOS=linux GOARCH=amd64 $(GoCmd) build $(BUILDARGS) -o $(EXE)

$(EXE32): hello.go
	env GOOS=linux GOARCH=386 $(GoCmd) build $(BUILDARGS) -o $(EXE32)

$(EXEWIN): hello.go
	env GOOS=windows GOARCH=amd64 $(GoCmd) build $(BUILDARGS) -o $(EXEWIN)

$(EXEWIN32): hello.go
	env GOOS=windows GOARCH=386 $(GoCmd) build $(BUILDARGS) -o $(EXEWIN32)

analyze:
	$(GoCmd) vet

format:
	$(GoCmd) fmt

clean:
	-rm -f $(EXE)
	-rm -f $(EXE32)
	-rm -f $(EXEWIN)
	-rm -f $(EXEWIN32)
	-rm -f debug
