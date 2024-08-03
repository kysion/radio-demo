// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IDemo interface {
		// Auth 授权
		Auth(ctx context.Context, info interface{}) error
		// UnAuth 取消授权
		UnAuth(ctx context.Context, info interface{}) error
	}
)

var (
	localDemo IDemo
)

func Demo() IDemo {
	if localDemo == nil {
		panic("implement not found for interface IDemo, forgot register?")
	}
	return localDemo
}

func RegisterDemo(i IDemo) {
	localDemo = i
}
