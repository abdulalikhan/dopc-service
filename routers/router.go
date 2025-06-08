package routers

import (
	"dopc-service/handlers"
	"github.com/beego/beego/v2/server/web"
)

func InitializeRoutes() {
	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/delivery-order-price",
			web.NSInclude(
				&handlers.DeliveryPriceHandler{},
			),
			web.NSRouter("/", &handlers.DeliveryPriceHandler{}, "get:Get"),
		),
	)
	web.AddNamespace(ns)
}
