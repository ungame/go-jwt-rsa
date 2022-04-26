build:
	- mkdir build
	go build -o build/main.exe .

ut:
	go clean --cache
	go test --cover go-jwt-rsa/...

run:
	build/main.exe

clear:
	- rmdir /q /s build


all: clear build ut run
