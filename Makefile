build:
	go build -o server ./cmd/api-cache-store/main.go

run: build
	./server

# air -c .air.toml
watch:
	reflex -s -r '\.go$$' make run

test:
	go test -v ./...

	

	