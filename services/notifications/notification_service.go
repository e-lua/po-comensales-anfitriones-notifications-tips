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

	var notification_var []models.Mo_NotificationShow

	notifications, error_show := notification_repository.Mo_Find(iduser_int, page_int, idtype_int)
	if error_show != nil {
		return 500, true, "Error en el servidor interno al intentar mostrar las notificaciones, detalles: " + error_show.Error(), notifications
	}

	error_update := notification_repository.Mo_Update(iduser_int, idtype_int)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar el estado de las notificacion, detalles: " + error_update.Error(), notification_var
	}

	return 201, false, "", notifications
}
