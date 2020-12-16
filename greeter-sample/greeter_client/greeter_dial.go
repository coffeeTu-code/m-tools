package greeter_client

import (
	"context"
	"google.golang.org/grpc"
	"log"
	greeter "m-tools/api/greeter/go"
	"m-tools/greeter-sample/configer"
	"m-tools/m-rpc/consul_helper"
	"time"
)

func init() {
	resolver, err := consul_helper.NewConsulResolver(configer.ConsulAddress, configer.ConsulPort, configer.TargetService, configer.MyService, consul_helper.SetInterval(10))
	if err != nil {
		log.Panic(err)
	}
	__GreeterClient = &GreeterClient{resolver: resolver}
}

// client
var __GreeterClient *GreeterClient

type GreeterClient struct {
	resolver *consul_helper.ConsulResolver
}

func GreeterDiscover() string {
	node := __GreeterClient.resolver.DiscoverNode()
	if node != nil {
		return node.Address
	}
	return ""
}

// dial
type GreeterContext struct {
	Request  *greeter.GreetingRequest
	Response *greeter.GreetingResponse
	Err      error
}

func GreeterDial(connTime, readTime int, target string, greeterCtx *GreeterContext) {
	ctxConn, celConn := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(connTime))
	defer celConn()
	conn, err := grpc.DialContext(ctxConn, target, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		greeterCtx.Err = err
		return
	}
	defer func() {
		_ = conn.Close()
	}()
	c := greeter.NewGreeterClient(conn)

	ctxRead, celRead := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(readTime))
	defer celRead()

	greeterCtx.Response, greeterCtx.Err = c.Greeting(ctxRead, greeterCtx.Request)
}
