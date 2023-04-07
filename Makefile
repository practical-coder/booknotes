.PHONY:all clean docker_build fmt booknotes_mac booknotes scheck vet
all: booknotes_mac booknotes docker_build
clean:
	rm -f booknotes_mac booknotes
docker_build:
	docker buildx build --compress -t booknotes_testfield:latest --progress=plain .
fmt:
	go fmt ./...
booknotes_mac:
	GOOS=darwin GOARCH=amd64 go build -o booknotes_mac
booknotes: scheck
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o booknotes -ldflags '-s -w'
scheck: vet
	staticcheck ./...
vet: fmt
	go vet ./...