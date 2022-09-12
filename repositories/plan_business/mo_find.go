package plan_business

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"
	"go.mongodb.org/mongo-driver/bson"
)

//Traeremos las notificaciones del Anfitrion
func Mo_Find(idbusiness int) (models.Mqtt_LegalIdentity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := models.MongoCN.Database("restoner")
	col := db.Collection("anfitriones_planbusiness")

	var resultado models.Mqtt_LegalIdentity

	condicion := bson.M{
		"idbusiness": idbusiness,
	}

	/*Cursor es como una tabla de base de datos donde se van a grabar los resultados
	y podre ir recorriendo 1 a la vez*/
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		return resultado, err
	}

	return resultado, nil
}
