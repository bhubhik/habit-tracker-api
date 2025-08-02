.PHONY: gql localstack proto run-grpc run-gql

# TODO: Once the docker-compose is updated to include all the services,
# 			create a target that will run all services in docker.

gql:
	@echo "Generating graphql code"
	gqlgen generate

localstack:
	@echo "Running localstack in docker"
	./localstack-init.sh

proto:
	@echo "Generating protos"
		protoc \
		--go_out=. \
		--go-grpc_out=. \
		proto/habit/v1/habit.proto

run-grpc:
	go run cmd/server/main.go

run-gql:
	go run main.go

