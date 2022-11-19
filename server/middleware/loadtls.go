package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

// 用https把这个中间件在router里面use一下就好

func LoadTls() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		//middleware := secure.New(secure.Options{
		//	SSLRedirect: true,
		//	SSLHost:     "localhost:443",
		//})
		//
		//writer2 := http.ResponseWriter(c.Response.Header, []byte{}, 500)
		//
		//err := middleware.Process(writer2, c.Request))
		//if err != nil {
		//	// 如果出现错误，请不要继续
		//	fmt.Println(err)
		//	return
		//}
		// 继续往下处理
		c.Next(ctx)
	}
}

//type ResponseWriter interface {
//	Header() Header
//	Write([]byte) (int, error)
//	WriteHeader(statusCode int)
//}
