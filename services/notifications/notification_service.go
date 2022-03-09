package notifications

import (

	//REPOSITORIES
	"fmt"
	"log"
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

	/*=============FIREBASE CLOUD MESSAGE=============*/
	data := map[string]string{
		"msg": "Hello World1",
		"sum": "Happy Day",
	}
	c := fcm.NewFCM("AAAAnX6Cb1g:APA91bESe-FuIIH2z_lv6JvWQX-r5hz_Ta6jRii-TwBZnqZQZBgSz9rSo5TIybr0RkznoQkY21WoA1yrdQUo0IuUWiZrrAIDLxzva5BZEoF4z5UPVIyFTv1-87_c8p_u3EDD93GiGQvf")
	response, err := c.Send(fcm.Message{
		Data:             data,
		RegistrationIDs:  []string{"cxHlE6xFNYZpJJPD5Cz8JA:APA91bELN9h25_QCCCxa3RqQz49dASYlXHBGexgup7kFQyD8aDRiDcqDhQciFeCusSFUEX0UgTv3XWDxbKe2TZUtRgRU7nPfjLl8uWQTAovia3fJVhODnc_9NG9b0Bv3iepipAJUOouc"},
		ContentAvailable: true,
		Priority:         fcm.PriorityHigh,
		Notification: fcm.Notification{
			Title: "Hello",
			Body:  "World",
		},
	})
	if err != nil {
		log.Fatal("Error en la conexión con Firebase Cloud Message, detalles: " + err.Error())
	}
	/*======================================*/

	fmt.Println("Status Code   :", response.StatusCode)
	fmt.Println("Success       :", response.Success)
	fmt.Println("Fail          :", response.Fail)
	fmt.Println("Canonical_ids :", response.CanonicalIDs)
	fmt.Println("Topic MsgId   :", response.MsgID)

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
