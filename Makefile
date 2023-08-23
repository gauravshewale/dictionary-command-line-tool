BINARY_NAME=command-line-dictionary

build:
  go build -o ${BINARY_NAME}.exe .\main.go

run: build
 .\${BINARY_NAME}