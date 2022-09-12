package notifications

import (

	//REPOSITORIES

	"bytes"
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

	/*==============ENVIO DE NOTIFICACIÓN A UN ANFITRION===================*/
	if notification.Priority == 1 && notification.TypeUser == 1 {

		//Obtenemos los datos del auth
		send_to_get_idbusiness := map[string]interface{}{
			"idbusiness": notification.IDUser,
			"type":       1,
		}
		json_data, _ := json.Marshal(send_to_get_idbusiness)
		respuesta, _ := http.Post("http://a-registro-authenticacion.restoner-api.fun:80/v1/export", "application/json", bytes.NewBuffer(json_data))
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

	}
	/*============================================================================*/

	/*=============ENVIO DE NOTIFICACIÓN A ALGUNOS ANFITRIONES==================*/
	if notification.Priority == 1 && notification.TypeUser == 6 {

		//Obtenemos los datos del auth
		send_to_get_idbusiness := map[string]interface{}{
			"type":           2,
			"manybusinesses": notification.MultipleUser,
		}
		json_data, _ := json.Marshal(send_to_get_idbusiness)
		respuesta, _ := http.Post("http://a-registro-authenticacion.restoner-api.fun:80/v1/export", "application/json", bytes.NewBuffer(json_data))
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

	}
	/*============================================================================*/

	/*=============ENVIO DE NOTIFICACIÓN A TODOS LOS ANFITRIONES==================*/
	if notification.Priority == 1 && notification.TypeUser == 4 {

		//Obtenemos los datos del auth
		send_to_get_idbusiness := map[string]interface{}{
			"type": 3,
		}
		json_data, _ := json.Marshal(send_to_get_idbusiness)
		respuesta, _ := http.Post("http://a-registro-authenticacion.restoner-api.fun:80/v1/export", "application/json", bytes.NewBuffer(json_data))
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
			Priority:         fcm.PriorityNormal,
			Notification: fcm.Notification{
				Title: notification.Title,
				Body:  notification.Message,
			},
		})
		if err != nil {
			log.Fatal("Error en la conexión con Firebase Cloud Message, detalles: " + err.Error())
		}

	}
	/*============================================================================*/

	/*==================ENVIO DE NOTIFICACIÓN A UN COMENSAL=================*/
	if notification.Priority == 1 && notification.TypeUser == 2 {

		//Obtenemos los datos del auth
		respuesta, _ := http.Get("http://c-registro-authenticacion.restoner-api.fun:3000/v1/export?idbusiness=" + strconv.Itoa(notification.IDUser) + "&type=1")
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

	}
	/*============================================================================*/

	/*=================ENVIO DE NOTIFICACIÓN A TODOS LOS COMENSALES===============*/
	if notification.Priority == 1 && notification.TypeUser == 5 {

		//Obtenemos los datos del auth
		respuesta, _ := http.Get("http://c-registro-authenticacion.restoner-api.fun:3000/v1/export?type=2")
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

	}
	/*============================================================================*/

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
