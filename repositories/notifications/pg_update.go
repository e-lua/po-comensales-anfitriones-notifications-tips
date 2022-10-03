package notifications

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"
)

func Pg_Update(iduser int, typeuser int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	//Actualizamos la foto de la categoría
	q := "UPDATE Notification SET wasview=true WHERE iduser=$1 AND typeuser=$2"
	if _, err_update := db.Exec(ctx, q, iduser, typeuser); err_update != nil {
		return err_update
	}

	return nil
}
