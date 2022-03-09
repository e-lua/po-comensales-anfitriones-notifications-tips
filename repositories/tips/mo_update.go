package tips

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
	col := db.Collection("tips")

	updtString := bson.M{
		"$push": bson.M{
			"viewbusiness": iduser,
		},
	}

	filtro := bson.M{"typeuser": typeuser}

	_, error_update := col.UpdateMany(ctx, filtro, updtString)

	if error_update != nil {
		return error_update
	}

	return nil
}
