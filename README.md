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


## Benchmark

```
MacBook Pro, 2.2 GHz Quad-Core Intel Core i7, 16 GB 1600 MHz DDR3

$ make wrk

wrk -c 100 -t 10 -d 60s -s script/post.lua http://localhost:8081/v1/echo
Running 1m test @ http://localhost:8081/v1/echo
  10 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     6.11ms    5.57ms 123.36ms   94.20%
    Req/Sec     1.78k   284.24     2.51k    81.53%
  1065545 requests in 1.00m, 179.86MB read
Requests/sec:  17740.54
Transfer/sec:      2.99MB

wrk -c 100 -t 10 -d 60s -s script/post.lua http://localhost:8081/v1/http/echo
Running 1m test @ http://localhost:8081/v1/http/echo
  10 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.10ms    1.75ms  62.28ms   85.35%
    Req/Sec     5.04k   381.28     6.86k    79.48%
  3006556 requests in 1.00m, 378.48MB read
Requests/sec:  50099.69
Transfer/sec:      6.31MB
```
