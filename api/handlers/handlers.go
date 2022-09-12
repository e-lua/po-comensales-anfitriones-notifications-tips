package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/cors"

	"github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"
	notification "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/services/notifications"
	planbusiness "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/services/planbusiness"
	tip "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/services/tips"
)

func Manejadores() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CONSUMER
	go Consumer_LegalIdentity()

	e.GET("/", index)
	//VERSION
	version_1 := e.Group("/v1")

	/*====================FLUJO DE INFORMACIÓN====================*/

	//V1 FROM V1 TO ...TO ENTITY TIP
	router_tip := version_1.Group("/tip")
	router_tip.POST("", tip.TipsRouter_pg.AddTip)
	router_tip.GET("", tip.TipsRouter_pg.ShowTipsAll)

	//V1 FROM V1 TO ...TO ENTITY TIP
	router_notification := version_1.Group("/notification")
	router_notification.POST("", notification.NotificationsRouter_pg.AddNotification)
	router_notification.GET("", notification.NotificationsRouter_pg.ShowNotification)

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

func Consumer_LegalIdentity() {

	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {
		log.Fatal("Error connection canal " + error_conection.Error())
	}

	msgs, err_consume := ch.Consume("anfitrion/legalidentity", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola " + err_consume.Error())
	}

	noStopLegal := make(chan bool)

	go func() {
		for d := range msgs {
			var anfitriones_planbusiness_all []models.Mqtt_LegalIdentity
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err_consume := decoder.Decode(&anfitriones_planbusiness_all)
			if err_consume != nil {
				log.Fatal("Error decoding")
			}
			planbusiness.PlanbusinessRouter_pg.AddBusiness(anfitriones_planbusiness_all)

			time.Sleep(5 * time.Second)
		}
	}()

	<-noStopLegal
}
