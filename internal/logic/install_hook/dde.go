package install_hook

import (
  "context"
  "fmt"
  "github.com/kysion/base-library/base_hook"
  "github.com/kysion/radio-demo/service"
)

type sDemo struct {
}

func NewDemo() service.IDemo {
  result := &sDemo{}

  // 订阅 用户授权Hook 和 取消授权Hook
  //service.Gateway().UserHook().InstallHook(1, result.Auth)
  //service.Gateway().UserHook().InstallHook(2, result.UnAuth)

  service.Radio().GetCommonHook().InstallHook(1, result.Auth)

  service.Radio().GetOtherHook().InstallHook(1, func(ctx context.Context, info *base_hook.User) error {
    fmt.Println("Auth-OtherHook")
    fmt.Println(info)

    return nil
  })
  //service.Radio().GetCommonHook().InstallHook(2, result.UnAuth)

  return result
}

func init() {
  service.RegisterDemo(NewDemo())
}

// Auth 授权
func (s *sDemo) Auth(ctx context.Context, info interface{}) error {

  fmt.Println("Auth")

  return nil
}

// UnAuth 取消授权
func (s *sDemo) UnAuth(ctx context.Context, info interface{}) error {
  fmt.Println("UnAuth")
  return nil
}
