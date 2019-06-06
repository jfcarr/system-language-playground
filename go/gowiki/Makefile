default:
	@echo 'Targets:'
	@echo ' wiki'
	@echo ' analyze'
	@echo ' format'
	@echo ' clean'

wiki: wiki.go
	go build -o wiki

analyze:
	go vet

format:
	go fmt

clean:
	-rm -f wiki
	-rm -f TestPage.txt
	-rm -f ANewPage.txt
