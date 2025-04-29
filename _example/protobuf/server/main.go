package main

import (
	"context"
	"fmt"
	"github.com/ja7ad/captcha"
	"github.com/ja7ad/captcha/_example/protobuf/echo"
	"github.com/ja7ad/captcha/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type EchoService struct {
	echo.UnsafeEchoServiceServer
}

func main() {
	c, err := proto.NewCaptcha(captcha.CloudflareProvider, os.Getenv("CAPTCHA_SECRET"), "")
	if err != nil {
		log.Fatal(err)
	}

	sv := grpc.NewServer(grpc.ChainUnaryInterceptor(c.UnaryServerInterceptor()))
	echo.RegisterEchoServiceServer(sv, &EchoService{})

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := sv.Serve(ln); err != nil {
		log.Fatal(err)
	}
}

func (e *EchoService) Echo(ctx context.Context, in *echo.EchoRequest) (*echo.EchoResponse, error) {
	return &echo.EchoResponse{
		Message: fmt.Sprintf("echo %s", in.Name),
	}, nil
}
