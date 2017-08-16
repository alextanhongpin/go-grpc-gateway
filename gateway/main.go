package main

import (
	"log"
	"net/http"
	"path"
	"strings"

	"golang.org/x/net/context"

	gw "github.com/alextanhongpin/go-grpc-gateway/hello"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
		glog.Errorf("Not found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}
	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join("hello", p)
	http.ServeFile(w, r, p)
}
func run() error {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterHelloServiceHandlerFromEndpoint(ctx, mux, "localhost:9090", opts)
	if err != nil {
		return err
	}
	r := http.NewServeMux()
	r.HandleFunc("/swagger/", serveSwagger)
	r.Handle("/", mux)
	log.Println("listening to port *:8080")
	return http.ListenAndServe(":8080", r)
}

func main() {
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
