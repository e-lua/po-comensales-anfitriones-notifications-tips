package notifications

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Mo_Update(iduser int, typeuser int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN.Database("restoner")
	col := db.Collection("notifications")

	updtString := bson.M{
		"$set": bson.M{
			"wasview": true,
		},
	}

	filtro := bson.M{"iduser": iduser, "typeuser": typeuser}

	_, error_update := col.UpdateOne(ctx, filtro, updtString)

	if error_update != nil {
		return error_update
	}

	return nil
}
