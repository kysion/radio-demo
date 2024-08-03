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
  return (&sRadio{}).registerHookMessage()
}

func init() {
  service.RegisterRadio(NewRadio())
  service.Radio().MakePublish()
}

func (s *sRadio) registerHookMessage() service.IRadio {

  // 注册网关Hook消息,用于接收通过网络发送来的HOOK消息
  base_hook.RegisterHookMessage(&s.DefaultHook)
  base_hook.RegisterHookMessage(&s.OtherHook)

  // 注册网关Hook消息
  //base_hook.RegisterHookMessage(base_enum.Hook.BusinessType.Default, &s.DefaultHook)
  //base_hook.RegisterHookMessage(base_enum.Hook.BusinessType.Default, &s.OtherHook)

  //base_hook.Gateway().RegisterHookMessage(func(model base_model.HookModel) {
  //  if model.BusinessType() == base_enum.Hook.BusinessType.Default {
  //    // model.Data 可以获取到消息内容,按实际使用场景按需转换数据类型
  //    // 广播消息
  //    base_hook.PublishHookMessage(context.Background(), &s.DefaultHook, base_hook.Option{Data: model.Data})
  //
  //    // 广播消息
  //    base_hook.PublishHookMessage(context.Background(), &s.OtherHook, base_hook.Option{Data: model.Data})
  //
  //    //OtherHook.Iterator(func(key int, value base_hook.DefaultHookFunc) {
  //    //  err := value(context.Background(), model.Data)
  //    //  if err != nil {
  //    //    fmt.Println(err)
  //    //  }
  //    //})
  //  }
  //})
  return s
}

func (s *sRadio) GetDefaultHook() *base_hook.BaseHook[int, base_hook.DefaultHookFunc] {
  return &s.DefaultHook
}

func (s *sRadio) GetOtherHook() *base_hook.BaseHook[int, base_hook.UserHookFunc] {
  return &s.OtherHook
}

func (s *sRadio) MakePublish() {
  //s.Publish(context.Background(), "abcd")
  //base_hook.PublishHookMessage(context.Background(), &s.DefaultHook, base_hook.Option{Data: "abcd"})
}

func bc[T any]() error {
  return nil
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
