package command

import (
	"fmt"
	"github.com/oklog/oklog/pkg/group"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	greeter "m-tools/api/greeter/go"
	"m-tools/greeter-sample/configer"
	"m-tools/greeter-sample/greeter_server"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func NewServerCmd() *cobra.Command {
	var serverCmd = &cobra.Command{
		Use:   "serve",
		Short: "server start",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			timeBegin := time.Now()
			log.Println("retarget server init @" + timeBegin.String())
		},
		Run: func(cmd *cobra.Command, args []string) {
			timeBegin := time.Now()
			log.Println("retarget server start @" + timeBegin.String())

			var service greeter_server.Service

			var (
				endpoints   = greeter_server.MakeServerEndpoints(service)
				httpHandler = greeter_server.NewHTTPHandler(endpoints)
				registar    = greeter_server.ConsulRegister(configer.GRPCPort)
				grpcServer  = greeter_server.NewGRPCServer(endpoints)
			)

			var g group.Group
			{
				// This function just sits and waits for ctrl-C.
				cancelInterrupt := make(chan struct{})
				g.Add(func() error {
					c := make(chan os.Signal, 1)
					signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
					select {
					case sig := <-c:
						return fmt.Errorf("received signal %s", sig)
					case <-cancelInterrupt:
						return nil
					}
				}, func(err error) {
					close(cancelInterrupt)
				})
			}
			{
				// The gRPC listener mounts the Go kit gRPC server we created.
				grpcListener, err := net.Listen("tcp", ":"+configer.GRPCPort)
				if err != nil {
					log.Panic("transport", "gRPC", "during", "Listen", "err", err)
				}
				g.Add(func() error {
					registar.Register()
					baseServer := grpc.NewServer()
					greeter.RegisterGreeterServer(baseServer, grpcServer)
					grpc_health_v1.RegisterHealthServer(baseServer, &greeter_server.Health{})
					return baseServer.Serve(grpcListener)
				}, func(err error) {
					registar.Deregister()
					grpcListener.Close()
				})
			}
			{
				// The service discovery registration.
				httpServer := &http.Server{
					Addr:           ":" + configer.HTTPPort,
					ReadTimeout:    10 * time.Second,
					WriteTimeout:   10 * time.Minute,
					MaxHeaderBytes: 1 << 20,
					Handler:        httpHandler,
				}
				g.Add(func() error {
					return httpServer.ListenAndServe()
				}, func(err error) {
					_ = httpServer.Close()
				})
			}

		},
	}
	return serverCmd
}
