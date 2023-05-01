package service

import "context"

// 客户端在建立连接的时候 传入opts参数 实现添加认证头部
// 需要实现该接口PerRPCCredentials实现添加认证信息

func NewClientAuthInfo(id, key string) *ClientAuthInfo {
	return &ClientAuthInfo{
		secretId:  id,
		secretKey: key,
	}
}

type ClientAuthInfo struct {
	secretId  string
	secretKey string
}

func (c *ClientAuthInfo) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"clientId":  c.secretId,
		"clientKey": c.secretKey,
	}, nil
}

func (c *ClientAuthInfo) RequireTransportSecurity() bool {
	// 不设置证书
	return false
}
