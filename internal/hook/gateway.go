package hook

import (
  "context"
  "fmt"
  "github.com/kysion/base-library/base_hook"
  "github.com/kysion/base-library/base_model"
  "github.com/kysion/base-library/base_model/base_enum"
)

type sGateway struct {
  userHook *base_hook.BaseHook[int, base_hook.CommonHookFunc]
}

var localGateway *sGateway

func init() {
  Gateway()
}

func Gateway() *sGateway {
  if localGateway != nil {
    return localGateway
  }

  localGateway = &sGateway{
    userHook: &base_hook.BaseHook[int, base_hook.CommonHookFunc]{},
  }

  return localGateway
}

//
//func init() {
//  //service.RegisterGateway(NewGateway())
//}

func (s *sGateway) UserHook() *base_hook.BaseHook[int, base_hook.CommonHookFunc] {
  return s.userHook
}

func (s *sGateway) BroadcastMessage(model base_model.HookModel) {
  if model.BusinessType() == base_enum.Hook.BusinessType.Default {
    // model.Data 可以获取到消息内容,按实际使用场景按需转换数据类型
    // 广播消息
    s.userHook.Iterator(func(key int, value base_hook.CommonHookFunc) {
      err := value(context.Background(), model.Data)
      if err != nil {
        fmt.Println(err)
      }
    })
  }
}
