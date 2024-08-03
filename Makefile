server:
	go build -o bin/gordirekt-server cmd/gordirekt-server/main.go
	./bin/gordirekt-server

app:
	go build -o bin/gordirekt cmd/gordirekt/main.go
	./bin/gordirekt
