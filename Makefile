default: run

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...

run: fmt vet test
	cd cmd/complete-study-guide && go run .