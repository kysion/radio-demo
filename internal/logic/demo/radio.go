package demo

import (
  "context"
  "github.com/kysion/base-library/base_hook"
  //"github.com/kysion/radio-demo/model/hook"
  "github.com/kysion/radio-demo/service"
)

type sRadio struct {
  // 用户授权Hook
  CommonHook base_hook.BaseHook[int, base_hook.CommonHookFunc]

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
  base_hook.RegisterHookMessage(&s.CommonHook)
  base_hook.RegisterHookMessage(&s.OtherHook)

  // 注册网关Hook消息
  //base_hook.RegisterHookMessage(base_enum.Hook.BusinessType.Default, &s.CommonHook)
  //base_hook.RegisterHookMessage(base_enum.Hook.BusinessType.Default, &s.OtherHook)

  //base_hook.Gateway().RegisterHookMessage(func(model base_model.HookModel) {
  //  if model.BusinessType() == base_enum.Hook.BusinessType.Default {
  //    // model.Data 可以获取到消息内容,按实际使用场景按需转换数据类型
  //    // 广播消息
  //    base_hook.PublishHookMessage(context.Background(), &s.CommonHook, base_hook.Option{Data: model.Data})
  //
  //    // 广播消息
  //    base_hook.PublishHookMessage(context.Background(), &s.OtherHook, base_hook.Option{Data: model.Data})
  //
  //    //OtherHook.Iterator(func(key int, value base_hook.CommonHookFunc) {
  //    //  err := value(context.Background(), model.Data)
  //    //  if err != nil {
  //    //    fmt.Println(err)
  //    //  }
  //    //})
  //  }
  //})
  return s
}

func (s *sRadio) GetCommonHook() *base_hook.BaseHook[int, base_hook.CommonHookFunc] {
  return &s.CommonHook
}

func (s *sRadio) GetOtherHook() *base_hook.BaseHook[int, base_hook.UserHookFunc] {
  return &s.OtherHook
}

func (s *sRadio) MakePublish() {
  //s.Publish(context.Background(), "abcd")
  //base_hook.PublishHookMessage(context.Background(), &s.CommonHook, base_hook.Option{Data: "abcd"})
}

func bc[T any]() error {
  return nil
}

// Publish 发布广播
func (s *sRadio) Publish(ctx context.Context, option base_hook.Option, isRemoteMessage ...bool) error {
  //base_hook.PublishHookMessage(context.Background(), &s.CommonHook, option)
  //s.CommonHook.Iterator(func(key int, value base_hook.CommonHookFunc) {
  //  err := value(ctx, option.Data)
  //  if err != nil {
  //    fmt.Println(err)
  //  }
  //}, option)
  base_hook.PublishHookMessage(context.Background(), &s.OtherHook, option)

  //base_hook.PublishHookMessage(context.Background(), &s.CommonHook, option)
  //base_hook.PublishHookMessage(context.Background(), &s.OtherHook, option, func(ctx context.Context, data any, value any) error {
  //  f := value.(base_hook.UserHookFunc)
  //  newData := data.(*base_hook.User)
  //  return f(ctx, newData)
  //})

  return nil
}
