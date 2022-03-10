package notifications

import (

	//REPOSITORIES

	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"gopkg.in/maddevsio/fcm.v1"

	models "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"
	notification_repository "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/repositories/notifications"
)

func AddNotification_Service(notification models.Mo_Notifications) (int, bool, string, string) {

	notification.DateRegistered = time.Now()
	notification.WasView = false

	//Obtenemos las categorias
	error_update := notification_repository.Mo_Add(notification)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar agregar el Notificacion, detalles: " + error_update.Error(), ""
	}

	if notification.Priority == 1 && notification.TypeUser == 1 {

		//Obtenemos los datos del auth
		respuesta, _ := http.Get("http://a-registro-authenticacion.restoner-api.fun:5000/v1/export/" + strconv.Itoa(notification.IDUser))
		var get_respuesta Income_IDDevice
		error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
		if error_decode_respuesta != nil {
			return 500, true, "Error en el sevidor interno al intentar obtener todos los ID Device, detalles: " + error_decode_respuesta.Error(), ""
		}

		/*=============FIREBASE CLOUD MESSAGE=============*/
		data := map[string]string{
			"msg": notification.Message,
			"sum": notification.Message,
		}
		c := fcm.NewFCM("AAAAnX6Cb1g:APA91bESe-FuIIH2z_lv6JvWQX-r5hz_Ta6jRii-TwBZnqZQZBgSz9rSo5TIybr0RkznoQkY21WoA1yrdQUo0IuUWiZrrAIDLxzva5BZEoF4z5UPVIyFTv1-87_c8p_u3EDD93GiGQvf")
		_, err := c.Send(fcm.Message{
			Data:             data,
			RegistrationIDs:  get_respuesta.Data,
			ContentAvailable: true,
			Priority:         fcm.PriorityHigh,
			Notification: fcm.Notification{
				Title: notification.Title,
				Body:  notification.Message,
			},
		})
		if err != nil {
			log.Fatal("Error en la conexión con Firebase Cloud Message, detalles: " + err.Error())
		}
		/*===============================================*/

	}

	if notification.Priority == 1 && notification.TypeUser == 2 {

		//Obtenemos los datos del auth
		respuesta, _ := http.Get("http://c-registro-authenticacion.restoner-api.fun:3000/v1/export/" + strconv.Itoa(notification.IDUser))
		var get_respuesta Income_IDDevice
		error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
		if error_decode_respuesta != nil {
			return 500, true, "Error en el sevidor interno al intentar obtener todos los ID Device, detalles: " + error_decode_respuesta.Error(), ""
		}

		/*=============FIREBASE CLOUD MESSAGE=============*/
		data := map[string]string{
			"msg": notification.Message,
			"sum": notification.Message,
		}
		c := fcm.NewFCM("AAAAnX6Cb1g:APA91bESe-FuIIH2z_lv6JvWQX-r5hz_Ta6jRii-TwBZnqZQZBgSz9rSo5TIybr0RkznoQkY21WoA1yrdQUo0IuUWiZrrAIDLxzva5BZEoF4z5UPVIyFTv1-87_c8p_u3EDD93GiGQvf")
		_, err := c.Send(fcm.Message{
			Data:             data,
			RegistrationIDs:  get_respuesta.Data,
			ContentAvailable: true,
			Priority:         fcm.PriorityHigh,
			Notification: fcm.Notification{
				Title: notification.Title,
				Body:  notification.Message,
			},
		})
		if err != nil {
			log.Fatal("Error en la conexión con Firebase Cloud Message, detalles: " + err.Error())
		}
		/*===============================================*/

	}

	return 201, false, "", "Notificacion agregado correctamente"
}

func ShowNotification_Service(iduser_int int, page_int int64, idtype_int int) (int, bool, string, []models.Mo_NotificationShow) {

	var notifications_trash []models.Mo_NotificationShow

	notifications, error_show := notification_repository.Mo_Find(iduser_int, page_int, idtype_int)
	if error_show != nil {
		return 500, true, "Error en el servidor interno al intentar mostrar las notificaciones, detalles: " + error_show.Error(), notifications
	}

	error_update := notification_repository.Mo_Update(iduser_int, idtype_int)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar las notificaciones, detalles: " + error_update.Error(), notifications_trash
	}

	return 201, false, "", notifications
}
