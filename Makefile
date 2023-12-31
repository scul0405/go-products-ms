.PHONY:

# ==============================================================================

run:
	go run cmd/main.go

client:
	go run internal/client/client.go

proto:
	rm -f proto/product/*.go
	protoc proto/product/product.proto --go_out=proto/product --go-grpc_out=proto/product

local:
	docker-compose -f docker-compose.local.yaml up

crate_topics:
	#docker exec -it kafka1 kafka-topics --zookeeper zookeeper:2181 --create --topic products --partitions 3 --replication-factor 2
	docker exec -it kafka1 kafka-topics --zookeeper zookeeper:2181 --create --topic create-product --partitions 3 --replication-factor 2
	docker exec -it kafka1 kafka-topics --zookeeper zookeeper:2181 --create --topic update-product --partitions 3 --replication-factor 2
	docker exec -it kafka1 kafka-topics --zookeeper zookeeper:2181 --create --topic dead-letter-queue --partitions 3 --replication-factor 2

cert:
	cd ./ssl && sh instructions.sh

# ==============================================================================
# Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -t -d -v ./...
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache

# ==============================================================================