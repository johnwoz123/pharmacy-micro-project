
build:
	protoc -I. --micro_out=. --go_out=plugins=micro:. \
	  		proto/prescription/pharmacy.proto

	GOOS=linux GOARCH=amd64 go build

#	  DOCKER
	docker build -t prescription-service .

#	  RUN DOCKER
run:
	docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 prescription-service
