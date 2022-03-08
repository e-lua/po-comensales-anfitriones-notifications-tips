package notifications

import (

	//REPOSITORIES
	"time"

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

	return 201, false, "", "Notificacion agregado correctamente"
}

func ShowNotification_Service(iduser_int int, page_int int64, idtype_int int) (int, bool, string, []models.Mo_NotificationShow) {

	notifications, error_show := notification_repository.Mo_Find(iduser_int, page_int, idtype_int)
	if error_show != nil {
		return 500, true, "Error en el servidor interno al intentar mostrar las notificaciones, detalles: " + error_show.Error(), notifications
	}

	return 201, false, "", notifications
}

func UpdateNotification_Service(data_idbusiness int, type_int int) (int, bool, string, string) {

	error_show := notification_repository.Mo_Update(data_idbusiness, type_int)
	if error_show != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar las notificaciones, detalles: " + error_show.Error(), ""
	}

	return 201, false, "", "Actualizacion realizada correctamente"
}
