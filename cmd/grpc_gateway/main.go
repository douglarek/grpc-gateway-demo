package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	gw "github.com/douglarek/grpc-gateway-demo/proto/gen/go/echo/service/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

// logging for grpc-gateway log
func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		grpclog.Infof("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func httpEcho(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var b struct{ Value string }
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, `{"error":"Hello  %s"}`, err)
		return
	}

	_, _ = fmt.Fprintf(w, `{"value":"Hello  %s"}`, b.Value)
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterEchoServiceHandlerFromEndpoint(ctx, gwmux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwmux) // proxy calls to gRPC server endpoint
	mux.Handle("/swaggerui/echo_service.swagger.json", http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./proto/gen/openapiv2/echo/service/v1"))))
	mux.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./swaggerui"))))
	mux.HandleFunc("/v1/http/echo", httpEcho) // same api using http

	go func() {
		grpclog.Infoln(http.ListenAndServe(":6060", nil)) // for pprof
	}()

	// Start HTTP server
	return http.ListenAndServe(":8081", logging(mux))
}

func main() {
	flag.Parse()

	grpclog.Fatal(run())
}
