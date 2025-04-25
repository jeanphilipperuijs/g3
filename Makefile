BUILD_PATH=./build
SOURCE=cmd/g3/main.go
BINARY_NAME=${BUILD_PATH}/g3

test:
	go test -v ${SOURCE}

build:
	rm -rf ${BUILD_PATH}
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux_amd64 ${SOURCE}
	GOARCH=arm64 GOOS=linux go build -o ${BINARY_NAME}-linux_arm64 ${SOURCE}
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin_amd64 ${SOURCE}
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows_amd64 ${SOURCE}
