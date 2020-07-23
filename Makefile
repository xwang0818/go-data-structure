export GOPATH=$(PWD)

test:
	@echo GOPATH = $(value GOPATH)
	go test ./...

clean:
