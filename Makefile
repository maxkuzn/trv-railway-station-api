.PHONY: build
build:
	go build cmd/trv-railway-station-api/main.go

.PHONY: test
test:
	go test -v ./...
