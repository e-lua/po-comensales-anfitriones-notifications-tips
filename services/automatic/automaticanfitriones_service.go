package automatic

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/maddevsio/fcm.v1"

	models "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"
	notified_repository "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/repositories/notified"
)

func AddNotificationInsumo_Service(insumo_toend []models.Pg_Notifications, insumo_ended []models.Pg_Notifications) (int, bool, string, string) {

	/*--------------------------------INSUMOS POR TERMINAR--------------------------------*/
	for _, notify_toend := range insumo_toend {

		quantity, error_get := notified_repository.Re_Get_Notified(notify_toend.IDUser, notify_toend.CodeNotify)
		if error_get != nil {
			error_re := notified_repository.Re_Set_Notified(notify_toend.IDUser, notify_toend.CodeNotify, 0)
			if error_re != nil {
				return 500, true, error_re.Error(), ""
			}
		}

		if quantity < 1 {
			//Obtenemos los datos del auth
			send_to_get_idbusiness := map[string]interface{}{
				"idbusiness": notify_toend.IDUser,
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
				"msg": notify_toend.Message,
				"sum": notify_toend.Message,
			}
			c := fcm.NewFCM("AAAAnX6Cb1g:APA91bESe-FuIIH2z_lv6JvWQX-r5hz_Ta6jRii-TwBZnqZQZBgSz9rSo5TIybr0RkznoQkY21WoA1yrdQUo0IuUWiZrrAIDLxzva5BZEoF4z5UPVIyFTv1-87_c8p_u3EDD93GiGQvf")
			_, err := c.Send(fcm.Message{
				Data:             data,
				RegistrationIDs:  get_respuesta.Data,
				ContentAvailable: true,
				Priority:         fcm.PriorityHigh,
				Notification: fcm.Notification{
					Title: notify_toend.Title,
					Body:  notify_toend.Message,
				},
			})
			if err != nil {
				log.Fatal("Error en la conexión con Firebase Cloud Message, detalles: " + err.Error())
			}
		}

		error_re := notified_repository.Re_Set_Notified(notify_toend.IDUser, notify_toend.CodeNotify, quantity+1)
		if error_re != nil {
			return 500, true, error_re.Error(), ""
		}

	}
	/*-----------------------------------------------------------------------------------*/

	/*--------------------------------INSUMOS TERMINADO--------------------------------*/
	for _, notify_ended := range insumo_ended {

		quantity, error_get := notified_repository.Re_Get_Notified(notify_ended.IDUser, notify_ended.CodeNotify)
		if error_get != nil {
			error_re := notified_repository.Re_Set_Notified(notify_ended.IDUser, notify_ended.CodeNotify, 0)
			if error_re != nil {
				return 500, true, error_re.Error(), ""
			}
		}

		if quantity < 1 {
			//Obtenemos los datos del auth
			send_to_get_idbusiness := map[string]interface{}{
				"idbusiness": notify_ended.IDUser,
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
				"msg": notify_ended.Message,
				"sum": notify_ended.Message,
			}
			c := fcm.NewFCM("AAAAnX6Cb1g:APA91bESe-FuIIH2z_lv6JvWQX-r5hz_Ta6jRii-TwBZnqZQZBgSz9rSo5TIybr0RkznoQkY21WoA1yrdQUo0IuUWiZrrAIDLxzva5BZEoF4z5UPVIyFTv1-87_c8p_u3EDD93GiGQvf")
			_, err := c.Send(fcm.Message{
				Data:             data,
				RegistrationIDs:  get_respuesta.Data,
				ContentAvailable: true,
				Priority:         fcm.PriorityHigh,
				Notification: fcm.Notification{
					Title: notify_ended.Title,
					Body:  notify_ended.Message,
				},
			})
			if err != nil {
				log.Fatal("Error en la conexión con Firebase Cloud Message, detalles: " + err.Error())
			}
		}

		error_re := notified_repository.Re_Set_Notified(notify_ended.IDUser, notify_ended.CodeNotify, quantity+1)
		if error_re != nil {
			return 500, true, error_re.Error(), ""
		}

	}
	/*-----------------------------------------------------------------------------------*/

	return 201, false, "", "Insumo Notificado exitosamente"
}
