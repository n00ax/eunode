all: build

build/eunode:
	GOPATH=${CURDIR} GOOS=linux go build -o build/eunode src/main/entrypoint.go

clean:
	rm build/eunode

getdeps:
	GOPATH=${CURDIR} GOOS=linux go get go.uber.org/zap
	GOPATH=${CURDIR} GOOS=linux go get golang.org/x/sys/unix
