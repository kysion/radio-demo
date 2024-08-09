// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/kysion/base-library/base_hook"
)

type (
	IRadio interface {
		GetDefaultHook() *base_hook.BaseHook[int, base_hook.DefaultHookFunc]
		GetOtherHook() *base_hook.BaseHook[int, base_hook.UserHookFunc]
		// Publish 发布广播
		Publish(ctx context.Context, option base_hook.Option, isRemoteMessage ...bool) error
	}
)

var (
	localRadio IRadio
)

func Radio() IRadio {
	if localRadio == nil {
		panic("implement not found for interface IRadio, forgot register?")
	}
	return localRadio
}

func RegisterRadio(i IRadio) {
	localRadio = i
}
