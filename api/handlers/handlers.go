package api

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/cors"

	notification "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/services/notifications"
	tip "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/services/tips"
)

func Manejadores() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", index)
	//VERSION
	version_1 := e.Group("/v1")

	/*====================FLUJO DE INFORMACIÓN====================*/

	//V1 FROM V1 TO ...TO ENTITY TIP
	router_tip := version_1.Group("/tip")
	router_tip.POST("", tip.TipsRouter_pg.AddTip)
	router_tip.GET("", tip.TipsRouter_pg.ShowTips)

	//V1 FROM V1 TO ...TO ENTITY TIP
	router_notification := version_1.Group("/notification")
	router_notification.POST("", notification.NotificationsRouter_pg.AddNotification)
	router_notification.GET("", notification.NotificationsRouter_pg.ShowNotification)
	router_notification.PUT("", notification.NotificationsRouter_pg.UpdateNotification)

	//Abrimos el puerto
	PORT := os.Getenv("PORT")
	//Si dice que existe PORT
	if PORT == "" {
		PORT = "5000"
	}

	//cors son los permisos que se le da a la API
	//para que sea accesibl esde cualquier lugar
	handler := cors.AllowAll().Handler(e)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

func index(c echo.Context) error {
	return c.JSON(401, "Acceso no autorizado")
}
