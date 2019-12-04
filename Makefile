init:
	mkdir bin
	mkdir bin/dist

build:
	GOOS=windows GOARCH=amd64 go build -o bin/web-server.exe
	GOOS=linux   GOARCH=amd64 go build -o bin/web-server-linux
	GOOS=darwin  GOARCH=amd64 go build -o bin/web-server-mac
