MAIN_FILE = ./cmd/pngb/main.go
OUTPUT_FILE = ./bin/pngb.exe

run:
	go run ${MAIN_FILE}

build:
	go mod tidy
	go build -o ${OUTPUT_FILE} ${MAIN_FILE}

deps:
	go mod tidy
	go get github.com/joho/godotenv
	go get github.com/gorilla/mux
	go get go.mongodb.org/mongo-driver