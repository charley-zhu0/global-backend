.PHONY: all build run gotool clean help

BINARY="gloabl-backend"

all: gotool build

build:
	go build -o ${BINARY} -v ./src/app.go 

gotool:
	gofmt -w ./

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

run:
	go run ./src/main.go

help:
	@echo "make build"
	@echo "       build the source code"
	@echo "make clean"
	@echo "       remove binary file"
	@echo "make run"
	@echo "       run the source code"
	@echo "make gotool"
	@echo "       run go fmt against code"
	@echo "make help"
	@echo "       show this message"
