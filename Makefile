protobuf:
	protoc	--go_out=./models/  --go_opt=paths=source_relative \
		--go-grpc_out=./models/ --go-grpc_opt=paths=source_relative \
		--proto_path=./proto \
		$$(find ./proto -type f -iname "*.proto")

docker-protobuf:
	docker run --rm \
		-v $$(pwd)/proto:/proto \
		-v $$(pwd)/models:/models \
		generate-go-code \
			protoc	--go_out=/models/  --go_opt=paths=source_relative \
				--go-grpc_out=/models/ --go-grpc_opt=paths=source_relative \
				--proto_path=proto \
				$$(find proto -type f -iname "*.proto")
