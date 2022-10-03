package notifications

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"
)

//Traeremos las notificaciones del Anfitrion
func Pg_Find(iduser int, limit int, offset int, typeuser int) ([]models.Pg_NotificationShow, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()
	q := "SELECT message,dateregistered,iduser,wasview FROM Notification WHERE iduser=$1 ORDER BY dateregistered DESC LIMIT $2 OFFSET $3"
	rows, error_shown := db.Query(ctx, q, iduser, limit, offset)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListNotification []models.Pg_NotificationShow

	if error_shown != nil {

		return oListNotification, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var oNotificaion models.Pg_NotificationShow
		rows.Scan(&oNotificaion.Message, &oNotificaion.DateRegistered, &oNotificaion.IDUser, &oNotificaion.WasView)
		oListNotification = append(oListNotification, oNotificaion)
	}

	return oListNotification, nil

}
