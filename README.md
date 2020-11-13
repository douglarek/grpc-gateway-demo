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

You need install [wrk](https://github.com/wg/wrk) and [ghz](https://github.com/bojand/ghz).

```
MacBook Pro, 2.2 GHz Quad-Core Intel Core i7, 16 GB 1600 MHz DDR3

$ make bench

wrk -c 100 -t 10 -d 60s -s script/post.lua http://localhost:8081/v1/echo
Running 1m test @ http://localhost:8081/v1/echo
  10 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     6.72ms    6.55ms 125.51ms   94.79%
    Req/Sec     1.66k   312.74     2.92k    80.13%
  994330 requests in 1.00m, 167.84MB read
Requests/sec:  16547.56
Transfer/sec:      2.79MB
wrk -c 100 -t 10 -d 60s -s script/post.lua http://localhost:8081/v1/http/echo
Running 1m test @ http://localhost:8081/v1/http/echo
  10 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.66ms    5.83ms 143.79ms   98.39%
    Req/Sec     4.86k   753.06     6.28k    83.08%
  2899982 requests in 1.00m, 365.06MB read
Requests/sec:  48315.93
Transfer/sec:      6.08MB
ghz --insecure -c 100 --connections 10 -x 60s --call echo.service.v1.EchoService.Echo -d '{"value":"world"}' localhost:9090

Summary:
  Count:	1471168
  Total:	60.00 s
  Slowest:	115.95 ms
  Fastest:	0.15 ms
  Average:	3.86 ms
  Requests/sec:	24519.07

Response time histogram:
  0.148 [1]	|
  11.729 [967417]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  23.309 [28073]	|∎
  34.889 [3171]	|
  46.469 [862]	|
  58.050 [266]	|
  69.630 [131]	|
  81.210 [37]	|
  92.791 [32]	|
  104.371 [2]	|
  115.951 [8]	|

Latency distribution:
  10 % in 1.10 ms
  25 % in 1.81 ms
  50 % in 2.99 ms
  75 % in 4.79 ms
  90 % in 7.44 ms
  95 % in 9.89 ms
  99 % in 18.18 ms

Status code distribution:
  [OK]            1471157 responses
  [Unavailable]   10 responses
  [Canceled]      1 responses

Error distribution:
  [10]   rpc error: code = Unavailable desc = transport is closing
  [1]    rpc error: code = Canceled desc = grpc: the client connection is closing
```
