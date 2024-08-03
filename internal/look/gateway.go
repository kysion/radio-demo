package look

import (
  "github.com/gogf/gf/v2/net/ghttp"
  "github.com/kysion/base-library/base_model"
  "github.com/kysion/base-library/utility/kmap"
)

type sGateway struct {
  wsArr kmap.HashMap[string, *ghttp.WebSocket]
}

var localGateway *sGateway

func Gateway() *sGateway {
  if localGateway != nil {
    return localGateway
  }

  localGateway = &sGateway{}

  return localGateway
}

func (s *sGateway) SetRequest(ws *ghttp.WebSocket) {
  localGateway.wsArr.Set(ws.RemoteAddr().String(), ws)
}

func (s *sGateway) Publish(target base_model.HookModel) error {
  ////ws := s.wsArr.Get(target.Target.GetAddr())
  //if ws == nil {
  //  return nil
  //}
  //
  //err := ws.WriteJSON(target)
  //if err != nil {
  //  return err
  //}
  return nil
}
