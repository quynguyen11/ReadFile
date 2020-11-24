hello:
	echo "Hello"

build:
	go build -v --ldflags "-s -w" -modfile go.mod -o bin/application -trimpath *.go

run:
	go run main.go