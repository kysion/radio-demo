package main

//
//func main() {
//  //info := base_model.HookModel{
//  //  Source: base_model.HookHostInfo{
//  //    Host: "127.0.0.1",
//  //    Port: 8000,
//  //  },
//  //  Target: base_model.HookHostInfo{
//  //    Host: "127.0.0.1",
//  //    Port: 8001,
//  //  },
//  //  BusinessType: base_enum.Hook.BusinessType.Default,
//  //  Data:         []byte("hello world"),
//  //}
//
//  bytes := gbinary.Encode("hello world")
//  fmt.Println(string(bytes))
//  var ff string
//  gbinary.Decode(bytes, &ff)
//  fmt.Println(ff)
//
//  // 解析消息
//
//  //bytes := gconv.Bytes(info)
//  //bytes := gbinary.Encode(info)
//  //fmt.Println(string(bytes))
//  //
//  //info2 := base_model.HookModel{}
//  //// 将字节转换为结构体
//  ////gbinary.BeDecode(bytes, &info2)
//  //gbinary.Decode(bytes, &info2)
//
//  cmd.Main.Run(gctx.GetInitCtx())
//}

import (
  "fmt"
  "github.com/gogf/gf/v2/encoding/gbinary"
  "github.com/gogf/gf/v2/os/gctx"
  "github.com/kysion/radio-demo/internal/cmd"
  _ "github.com/kysion/radio-demo/internal/logic"
)

type MyStruct struct {
  Field1 int
  Field2 string
}

func main() {
  //ctx := gctx.New()
  // 假设这是字节数据
  //str := MyStruct{
  //  Field1: 133,
  //  Field2: "sss",
  //}
  data := gbinary.BeEncodeString("abcccccc")
  myStruct := MyStruct{}
  toString := gbinary.DecodeToString(data)
  fmt.Println(toString)
  // 输出转换后的结构体
  fmt.Println(myStruct)

  cmd.Main.Run(gctx.GetInitCtx())

}
