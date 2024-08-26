package main

import (
    "context"
    "net/http"
    "github.com/twitchtv/twirp"
    "helloworld/helloworld"
)

type helloWorldServer struct{}

func (h *helloWorldServer) Hello(ctx context.Context, req *helloworld.HelloReq) (*helloworld.HelloResp, error) {
    if req.Name == "" {
        return nil, twirp.InvalidArgumentError("name", "cannot be empty")
    }
    return &helloworld.HelloResp{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
    server := &helloWorldServer{}
    twirpHandler := helloworld.NewHelloWorldServer(server, nil)

    http.ListenAndServe(":8080", twirpHandler)
}