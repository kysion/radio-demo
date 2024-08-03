package user_v1

import "github.com/gogf/gf/v2/frame/g"

type AuthReq struct { // 第三方应用的相关消息
  g.Meta `path:"/auth" method:"get" summary:"用户登陆" tags:"用户"`

  Username string `json:"username" form:"username" example:"admin" comment:"用户名"`
}

type AuthRes bool
