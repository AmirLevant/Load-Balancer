build:
	go build -o ./bin/lb ./cmd/lb/*
	go build -o ./bin/test-client ./cmd/test-client/*
	go build -o ./bin/test-server ./cmd/test-server/*

clean:
	rm -rf ./bin

