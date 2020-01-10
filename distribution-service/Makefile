
build:
	protoc -I. --micro_out=. --go_out=plugins=micro:. \
	  		proto/distribution/distribution.proto

#	  DOCKER
	docker build -t distribution-service .

#	  RUN DOCKER
run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 distribution-service
