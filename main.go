package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/danielgtaylor/huma/v2/humacli"
)

type Options struct {
	Port int `help:"Port to listen on" default:"8888"`
}
type Ping struct {
	Name  string `json:"name" doc:"Name to greet"`
	Count int    `json:"count" doc:"Number of times to greet" default:"1"`
}

type PingInput struct {
	Body    Ping `json:"body" doc:"Ping input"`
	RawBody []byte
}

type PingOutput struct {
	Body struct {
		Message string `json:"message" doc:"Greeting message" example:"Hello, world!"`
	}
}

func main() {
	cli := humacli.New(func(hooks humacli.Hooks, opts *Options) {
		router := http.NewServeMux()
		config := huma.DefaultConfig("My API", "1.0.0")
		api := humago.New(router, config)
		addRoute(api)

		srv := &http.Server{
			Addr:    fmt.Sprintf("%s:%d", "localhost", opts.Port),
			Handler: router,
		}

		hooks.OnStart(func() {
			log.Printf("Server is running with: host:%v port:%v\n", "localhost", opts.Port)

			err := srv.ListenAndServe()
			if err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		})
	})

	cli.Run()
}

func addRoute(api huma.API) {

	huma.Register(api, huma.Operation{
		OperationID: "Ping",
		Summary:     "Ping",
		Method:      http.MethodPost,
		Path:        "/ping",
	}, func(ctx context.Context, input *PingInput) (*PingOutput, error) {
		doWork(input.RawBody)
		resp := &PingOutput{}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Body.Name)
		return resp, nil
	})

}

func doWork(input []byte) {
	log.Printf("Doing work with input: %s\n", string(input))
	time.Sleep(4 * time.Second)
}
