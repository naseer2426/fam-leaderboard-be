APP_NAME=fam-leaderboard-be
BIN=bin/$(APP_NAME)

.PHONY: build run clean tidy

build:
	@mkdir -p bin
	go build -o $(BIN) .

run:
	./bin/fam-leaderboard-be
clean:
	rm -rf bin

tidy:
	go mod tidy


