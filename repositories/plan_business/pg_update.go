package plan_business

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"
)

func Pg_UpdateActive(active bool, idbusiness int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	//Actualizamos la foto de la categoría
	q := "UPDATE PlanBusiness SET active=$1 WHERE idbusiness=$2"
	if _, err_update := db.Exec(ctx, q, active, idbusiness); err_update != nil {
		return err_update
	}

	return nil
}
