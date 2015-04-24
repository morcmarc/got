clean:
	@rm -rf ./gottest*

test: clean
	go test ./... -cover