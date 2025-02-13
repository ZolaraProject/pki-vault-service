package pkivault

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health/grpc_health_v1"
)

var (
	DbPort     string
	DbHostname string
	DbUser     string
	DbPassword string
	DbName     string
)

type HealthChecker struct {
	health.UnimplementedHealthServer
}

func NewHealthChecker() *HealthChecker {
	return &HealthChecker{}
}

func (s *HealthChecker) Check(ctx context.Context, req *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{
		Status: health.HealthCheckResponse_SERVING,
	}, nil
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

// Run implements Run
func Run() {
	port := func(def string) string {
		port, ok := os.LookupEnv("EXPOSE_PORT")
		if !ok {
			return def
		}
		return port
	}("8080")
	log.Printf("started gRPC server on port %s", fmt.Sprintf("%v", port))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	provider := trace.NewNoopTracerProvider()
	s := grpc.NewServer(grpc.MaxRecvMsgSize(16*1024*1024), grpc.StatsHandler(otelgrpc.NewServerHandler(otelgrpc.WithTracerProvider(provider))))
	// RegisterConfigVaultServiceServer(s, &server{})

	healthService := NewHealthChecker()
	health.RegisterHealthServer(s, healthService)

	go func() {
		http.HandleFunc("/healthz", healthzHandler)
		log.Fatal(http.ListenAndServe(":8081", nil))
	}()

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGTERM) // Received after the preStop hook

	go func() {
		log.Println("Starting gRPC server")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	select {
	case c := <-termChan:
		log.Printf("Received signal %v, stopping gracefully", c)
		s.GracefulStop()
		log.Printf("Server stopped, exiting. ")
	}
}

func DbUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/zolara", DbUser, DbPassword, DbHostname, DbPort)
}
