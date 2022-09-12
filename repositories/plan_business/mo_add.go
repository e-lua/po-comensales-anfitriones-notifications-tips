package plan_business

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"
)

func Mo_Add(anfitriones_planbusiness models.Mqtt_LegalIdentity) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	db := models.MongoCN.Database("restoner")
	col := db.Collection("anfitriones_planbusiness")

	_, err := col.InsertOne(ctx, anfitriones_planbusiness)
	if err != nil {
		return err
	}

	return nil
}
