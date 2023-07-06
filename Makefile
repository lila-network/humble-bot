BINARY_NAME=humble-bot
default: build
 
build:
	go build -o ${BINARY_NAME} main.go
 
clean:
	go clean
	rm ${BINARY_NAME}
