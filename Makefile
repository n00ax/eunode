all: build

build/eunode:
	GOPATH=${CURDIR} GOOS=linux go build -o build/eunode src/main/entrypoint.go

clean:
	rm build/eunode