package core

func initServerLinux(opts ...config.Option) *server2.Hertz {
	h := server2.New(opts...)
	h.Use(recovery.Recovery())
	return h
}
