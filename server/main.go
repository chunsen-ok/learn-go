//Package main
//
// ## http 会话管理的几种方式。
//
// ### 基于服务器的会话管理
//
// 在服务端生成并管理 session id 。客户端访问时，需上传该 session id 。
// 服务端根据客户端的 session id ，去查找用户登录状态等，判断该用户是否登录了。
// 如果已经登录则可进行数据访问等其他操作。
//
// ### 基于 Cookie 的会话管理
//
// 服务器利用加密算法，将登录成功后的用户信息加密以及生成摘要等处理后，返回给客户端，
// 由客户端保存在 Cookie 中，而服务器不需要保存。每次客户端的访问，都需要带上这些信息，
// 服务端尽心解密，验证。
//
// ### 基于 token 的会话管理
//
//
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	srv := gin.Default()
	srv.StaticFile("/", "web/index.html")

	// srv.POST("/login", login)

	srv.Run(":8544")
}
