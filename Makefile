test:
	go test -v ./...

lint:
	pre-commit run -a

