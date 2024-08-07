server:
	go build -o bin/rdirekt-server cmd/rdirekt-server/main.go

app:
	go build -o bin/rdirekt cmd/rdirekt/main.go

clean:
	rm bin/*
