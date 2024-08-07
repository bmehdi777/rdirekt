server:
	go build -o bin/rdirekt-server cmd/rdirekt-server/main.go
	./bin/gordirekt-server

app:
	go build -o bin/rdirekt cmd/rdirekt/main.go
	./bin/gordirekt

clean:
	rm bin/*
