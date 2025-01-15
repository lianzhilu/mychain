build-my-chain:
	go mod tidy
	go build -o output/mychain main.go