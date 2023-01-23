// Code generated by Kitex v0.4.4. DO NOT EDIT.
package hello

import (
	server "github.com/cloudwego/kitex/server"
	api "github.com/go-demo/kitex-demo/kitex_gen/api"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler api.Hello, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
