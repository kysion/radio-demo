package demo

import (
  "context"
  "github.com/kysion/base-library/base_hook"
  //"github.com/kysion/radio-demo/model/hook"
  "github.com/kysion/radio-demo/service"
)

type sRadio struct {
  // 用户授权Hook
  DefaultHook base_hook.BaseHook[int, base_hook.DefaultHookFunc]

  // 其他Hook
  OtherHook base_hook.BaseHook[int, base_hook.UserHookFunc]
}

func NewRadio() service.IRadio {
  result := &sRadio{}
  // 注册网关Hook消息,用于接收通过网络发送来的HOOK消息
  base_hook.RegisterHookMessage(&result.DefaultHook)
  base_hook.RegisterHookMessage(&result.OtherHook)

  return result
}

func init() {
  service.RegisterRadio(NewRadio())
}

func (s *sRadio) GetDefaultHook() *base_hook.BaseHook[int, base_hook.DefaultHookFunc] {
  return &s.DefaultHook
}

func (s *sRadio) GetOtherHook() *base_hook.BaseHook[int, base_hook.UserHookFunc] {
  return &s.OtherHook
}

// Publish 发布广播
func (s *sRadio) Publish(ctx context.Context, option base_hook.Option, isRemoteMessage ...bool) error {
  //base_hook.PublishHookMessage(context.Background(), &s.DefaultHook, option)
  //s.DefaultHook.Iterator(func(key int, value base_hook.DefaultHookFunc) {
  //  err := value(ctx, option.Data)
  //  if err != nil {
  //    fmt.Println(err)
  //  }
  //}, option)

  // 使用方式1：Latest方式
  base_hook.PublishHookMessage(context.Background(), &s.OtherHook, option)

  // 使用方式2：Iterator原方式
  //s.OtherHook.Iterator(func(key int, value base_hook.UserHookFunc) {
  //  res := kconv.Struct(option.Data, &base_hook.User{})
  //  err := value(ctx, res)
  //  if err != nil {
  //    fmt.Println(err)
  //  }
  //}, option)

  //base_hook.PublishHookMessage(context.Background(), &s.DefaultHook, option)
  //base_hook.PublishHookMessage(context.Background(), &s.OtherHook, option, func(ctx context.Context, data any, value any) error {
  //  f := value.(base_hook.UserHookFunc)
  //  newData := data.(*base_hook.User)
  //  return f(ctx, newData)
  //})

  return nil
}
