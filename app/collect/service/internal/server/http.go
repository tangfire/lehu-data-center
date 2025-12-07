package server

//import (
//	v1 "lehu-data-center/api/collect/service/v1"
//	"lehu-data-center/app/collect/service/internal/conf"
//	"lehu-data-center/app/collect/service/internal/service"
//
//	"github.com/go-kratos/kratos/v2/log"
//	"github.com/go-kratos/kratos/v2/middleware/recovery"
//	"github.com/go-kratos/kratos/v2/transport/http"
//)
//
//// NewHTTPServer new an HTTP server.
//funcs NewHTTPServer(c *conf.Server,  collectService *service.CollectService, logger log.Logger) *http.Server {
//	var opts = []http.ServerOption{
//		http.Middleware(
//			recovery.Recovery(),
//		),
//	}
//	if c.Http.Network != "" {
//		opts = append(opts, http.Network(c.Http.Network))
//	}
//	if c.Http.Addr != "" {
//		opts = append(opts, http.Address(c.Http.Addr))
//	}
//	if c.Http.Timeout != nil {
//		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
//	}
//	srv := http.NewServer(opts...)
//	v1.RegisterCollectServer(srv, collectService)
//	return srv
//}
