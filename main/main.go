package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"awesomeProject1/calculator"
	"awesomeProject1/grpc/service/calc"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type calcServer struct {
	calc.UnimplementedCalcServer
}

func AuthorizationInter(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unknown, "invalid metadata")
	}

	tokenMD := md.Get("authorization")
	if len(tokenMD) == 0 {
		return nil, status.Error(codes.Unauthenticated, "empty token")
	}

	token := strings.TrimPrefix(tokenMD[0], "Bearer ")
	decodedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil
	})

	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Unauthenticated, "unauthorized invalid token")
	}

	if !decodedToken.Valid {
		return nil, status.Error(codes.Unauthenticated, "unauthorized invalid token")
	}

	return handler(ctx, req)
}

func (c *calcServer) Equation(context.Context, *calc.Input) (*calc.Output, error) {
	return &calc.Output{
		Result: calculator.Calc("", 5),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":8080"))
	if err != nil {
		log.Fatal(err)
	}

	interceptors := grpc.ChainUnaryInterceptor(AuthorizationInter)
	server := grpc.NewServer(interceptors)
	calcS := &calcServer{}
	calc.RegisterCalcServer(server, calcS)

	log.Printf("Serving and listening on: %s\n", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
