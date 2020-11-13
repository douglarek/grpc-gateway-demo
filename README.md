# grpc-gateway-demo

An example modified from the [official](https://github.com/grpc-ecosystem/grpc-gateway/) grpc-gateway that looks clearer.

## Generate proto

```
$ make deps // run once only
$ make proto
```

## Run grpc and grpc-gateway

```
$ export GRPC_GO_LOG_SEVERITY_LEVEL=info GRPC_GO_LOG_VERBOSITY_LEVEL=2 // for info log
$ go run cmd/grpc/main.go // in a terminal
$ go run cmd/grpc_gateway/main.go // in another terminal
```

## Request

```
$ curl -XPOST localhost:8081/v1/echo  -d '{"value":" world"}'

{"value":"Hello  world"}

```

## Access the swaggerui

Open [http://localhost:8081/swaggerui/](http://localhost:8081/swaggerui/) in the browser.
