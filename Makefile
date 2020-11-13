deps:
	go install \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc
proto:
	protoc -I proto -I proto/third_party/googleapis \
		--go_out proto/gen/go/ --go_opt paths=source_relative \
		--go-grpc_out proto/gen/go/ --go-grpc_opt paths=source_relative \
		echo/service/v1/echo_service.proto
	protoc -I proto -I proto/third_party/googleapis/ --grpc-gateway_out proto/gen/go \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
		echo/service/v1/echo_service.proto
	protoc -I proto/ -I proto/third_party/googleapis/ \
		--openapiv2_out proto/gen/openapiv2 \
		--openapiv2_opt logtostderr=true \
		echo/service/v1/echo_service.proto

wrk:
	wrk -c 100 -t 10 -d 60s -s script/post.lua http://localhost:8081/v1/echo
	wrk -c 100 -t 10 -d 60s -s script/post.lua http://localhost:8081/v1/http/echo

.PHONY: deps proto wrk


