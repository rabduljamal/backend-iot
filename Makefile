install:
	go mod download

run:
	CGO_ENABLED=1 GOOS=linux go run main.go

build:
	CGO_ENABLED=1 GOOS=linux go build -a -o main main.go

test:
	mkdir -p ./coverage && CGO_ENABLED=1 GOOS=linux go test -v ./... -coverprofile=./coverage/coverage.out

cleandep:
	go mod tidy