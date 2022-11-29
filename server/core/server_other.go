package core

<<<<<<< HEAD
import (
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	server2 "github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
)

=======
>>>>>>> 26e9bd732187484998dd76c73570fda2f588a100
func initServerLinux(opts ...config.Option) *server2.Hertz {
	h := server2.New(opts...)
	h.Use(recovery.Recovery())
	return h
}
