package main

import (
	"context"
	"fmt"
	"log"
	"mega_api/configs"
	"mega_api/handlers"
	
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	pb "mega_api/pb"
	httpSwagger "github.com/swaggo/http-swagger" 

    _ "mega_api/docs"

)

// @title Mega_api_HTTP_GRPC API Docs
// @version 1.0.0
// @contact.name Eduardo araujo
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9000
// @BasePath /

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
 
	//configuracao de rotas TODO: extrair para outro arquivo
	r.Get("/swagger/*",httpSwagger.WrapHandler)
	r.Post("/costumer", handlers.Create)
	r.Put("/costumer/{id}", handlers.Update)
	r.Delete("/costumer/{id}", handlers.Delete)
	r.Get("/costumer", handlers.List)
	r.Get("/costumer/{id}", handlers.Get)

	// configurando server http
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", configs.GetHttpServerPort()),
		Handler: r,
	}

	//configurando o servidor grpc
	grpcSrv := grpc.NewServer()
	orderQueueServer := &configs.ServerGRPC{}
	pb.RegisterOrderQueueServer(grpcSrv, orderQueueServer)
	lis, err := net.Listen("tcp", ":"+configs.GetGrpcServerPort())
	if err != nil {
		log.Fatalf("Falha ao iniciar server: %v", err)
	}

	// goroutine Servidor Http
	go func() {
		log.Println("HTTP server is running!")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()

	// goroutine servidor Grpc
	go func() {
		log.Println("gRPC server is running!")
		if err := grpcSrv.Serve(lis); err != nil {
			log.Fatalf("gRPC server failed: %v", err)
		}
	}()


	// Canal que verifica o desligamento dos servers
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	// Desliga os servidores HTTP e gRPC com um timeout de 5 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP server shutdown failed: %v", err)
	}

	// parando servidor grpc
	grpcSrv.GracefulStop()
	log.Println("Servers stopped gracefully")

}
