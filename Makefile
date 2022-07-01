BUILD_FLAGS := -ldflags="-w -s"
BUILD_CMD := go build ${BUILD_FLAGS}

all: linux macos windows

linux:
	GOOS=linux GOARCH=amd64 ${BUILD_CMD} -o bin/linux-x86_64 main.go
	GOOS=linux GOARCH=arm64 ${BUILD_CMD} -o bin/linux-arm64 main.go

macos:
	GOOS=darwin GOARCH=amd64 ${BUILD_CMD} -o bin/darwin-x86_64 main.go
	GOOS=darwin GOARCH=arm64 ${BUILD_CMD} -o bin/darwin-arm64 main.go

windows:
	GOOS=windows GOARCH=amd64 ${BUILD_CMD} -o bin/windows-x86_64.exe main.go
	GOOS=windows GOARCH=arm64 ${BUILD_CMD} -o bin/windows-arm64.exe main.go

run:
	go run main.go
