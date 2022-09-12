package plan_business

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Mo_Update(intput_legadata models.Mqtt_LegalIdentity) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := models.MongoCN.Database("restoner")
	col := db.Collection("anfitriones_planbusiness")

	updtString := bson.M{
		"$set": bson.M{
			"legalidentity":   intput_legadata.LegalIdentity,
			"typesuscription": intput_legadata.TypeSuscription,
			"iva":             intput_legadata.IVA,
		},
	}

	filtro := bson.M{"idbusiness": intput_legadata.IdBusiness}

	_, error_update := col.UpdateOne(ctx, filtro, updtString)

	if error_update != nil {
		return error_update
	}

	return nil
}
