build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t prescription-cli .

run:
	docker run prescription-cli
