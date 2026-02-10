package grpckit

import (
	"context"

	"gitlab.com/wit-id/service-hub-test/toolkit/log"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type HealthCheckServer struct {
	Serving         bool
	healthCheckFunc HealthCheckFunc
}

type HealthCheckFunc func(context.Context) error

func NewHealthcheckServer(hcFunc HealthCheckFunc) *HealthCheckServer {
	return &HealthCheckServer{Serving: true, healthCheckFunc: hcFunc}
}

func (s *HealthCheckServer) Check(ctx context.Context, _ *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	resp := &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}
	if !s.Serving {
		resp.Status = grpc_health_v1.HealthCheckResponse_NOT_SERVING
		return resp, nil
	}
	if s.healthCheckFunc != nil {
		if err := s.healthCheckFunc(ctx); err != nil {
			log.FromCtx(ctx).Error(err, "grpc health check")
			resp.Status = grpc_health_v1.HealthCheckResponse_NOT_SERVING
			return resp, nil
		}
	}
	return resp, nil
}

func (s *HealthCheckServer) List(ctx context.Context, _ *grpc_health_v1.HealthListRequest) (*grpc_health_v1.HealthListResponse, error) {
	st := grpc_health_v1.HealthCheckResponse_SERVING
	if !s.Serving {
		st = grpc_health_v1.HealthCheckResponse_NOT_SERVING
	} else if s.healthCheckFunc != nil {
		if err := s.healthCheckFunc(ctx); err != nil {
			st = grpc_health_v1.HealthCheckResponse_NOT_SERVING
		}
	}
	return &grpc_health_v1.HealthListResponse{
		Statuses: map[string]*grpc_health_v1.HealthCheckResponse{"": {Status: st}},
	}, nil
}

func (s *HealthCheckServer) Watch(_ *grpc_health_v1.HealthCheckRequest, stream grpc_health_v1.Health_WatchServer) error {
	return stream.Send(&grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	})
}