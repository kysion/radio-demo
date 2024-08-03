package cmd

import (
  "context"
  "github.com/gogf/gf/v2/frame/g"
  "github.com/gogf/gf/v2/net/ghttp"
  "github.com/gogf/gf/v2/os/gcmd"
  "github.com/kysion/base-library/base_hook"
  "github.com/kysion/radio-demo/internal/controller/user_controller"
)

var (
  Main = gcmd.Command{
    Name:  "main",
    Usage: "main",
    Brief: "start http server",
    Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
      s := g.Server()
      s.Group("/", func(group *ghttp.RouterGroup) {
        group.Middleware(ghttp.MiddlewareHandlerResponse)
        //group.Bind(
        //	hello.NewV1(),
        //)
      })

      s.Group("/user", func(group *ghttp.RouterGroup) {
        group.Bind(user_controller.User)
      })

      // UDP 服务
      //go gudp.NewServer("127.0.0.1:8001", func(conn *gudp.Conn) {
      //  defer conn.Close()
      //  for {
      //    data, err := conn.Recv(-1)
      //    if len(data) > 0 {
      //      if err := conn.Send(append([]byte("> "), data...)); err != nil {
      //        g.Log().Error(context.Background(), err)
      //      }
      //    }
      //    if err != nil {
      //      g.Log().Error(context.Background(), err)
      //    }
      //  }
      //}).Run()

      //hook.Gateway()
      s.BindHandler("/ws", base_hook.HookDistribution) // 接收消息
      // websocket 服务
      //s.BindHandler("/ws", func(r *ghttp.Request) {
      //  ws, err := r.WebSocket()
      //
      //  if err != nil {
      //    glog.Error(ctx, err)
      //    r.Exit()
      //  }
      //
      //  look.Gateway().SetRequest(ws)
      //  //
      //  //for {
      //  //  msgType, msg, err := ws.ReadMessage()
      //  //  if err != nil {
      //  //    return
      //  //  }
      //  //  //fmt.Println(string(msg))
      //  //  //
      //  //  //data := base_model.HookModel{}
      //  //  //err = gjson.DecodeTo(msg, &data)
      //  //  //if err != nil {
      //  //  //  fmt.Println(err)
      //  //  //  continue
      //  //  //}
      //  //
      //  //  if err = ws.WriteMessage(msgType, msg); err != nil {
      //  //    return
      //  //  }
      //  //}
      //})

      s.Run()
      return nil
    },
  }
)
