package plan_business

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"
)

func Pg_Add(planbusiness models.Mqtt_LegalIdentity) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	//Nos conectamos con la BD
	db := models.Conectar_Pg_DB()

	query_order := `INSERT INTO PlanBusiness(idbusiness,typesuscription) VALUES ($1,$2)`
	if _, err := db.Exec(ctx, query_order, planbusiness.IdBusiness, planbusiness.TypeSuscription); err != nil {
		return err
	}

	return nil
}
