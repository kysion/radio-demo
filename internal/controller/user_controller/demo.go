package user_controller

import (
  "context"
  "github.com/kysion/base-library/base_hook"
  "github.com/kysion/radio-demo/api/user_v1"
  "github.com/kysion/radio-demo/service"
)

var User = cUser{}

type cUser struct{}

func (c *cUser) Auth(ctx context.Context, req *user_v1.AuthReq) (user_v1.AuthRes, error) {
  //err := service.Radio().Publish(ctx, base_hook.Option{Data: req.Username})
  err := service.Radio().Publish(ctx, base_hook.Option{Data: &base_hook.User{Username: req.Username}, NetMessage: true})
  //base_hook.PublishHookMessage(context.Background(), &s.OtherHook, option)

  return err == nil, err
}
