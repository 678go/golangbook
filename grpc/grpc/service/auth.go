package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type grpcAuth struct{}

func (g *grpcAuth) Auth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 从上下文中获取信息
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("ctx is not an grpc incoming context")
	}
	id, key := g.getClientCredentialsFromMeta(md)
	if err := g.validateServiceCredential(id, key); err != nil {
		return nil, err
	}
	// 向下传递
	return handler(ctx, req)
}

// 获取id、key
func (g *grpcAuth) getClientCredentialsFromMeta(md metadata.MD) (id, key string) {
	idList := md.Get("clientId")
	keyList := md.Get("clientKey")
	if len(idList) > 0 {
		id = idList[0]
	}
	if len(keyList) > 0 {
		key = keyList[0]
	}
	return id, key
}

// 校验数据
func (g *grpcAuth) validateServiceCredential(id, key string) error {
	if id == "" && key == "" {
		return status.Errorf(codes.Unauthenticated, "client_id or client_secret is \"\"")
	}
	if !(id == "admin" && key == "admin") {
		return status.Errorf(codes.Unauthenticated, "client_id or client_secret invalidate")
	}
	return nil
}

func NewGrpcAuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return (&grpcAuth{}).Auth
}
