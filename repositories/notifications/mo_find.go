package notifications

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Traeremos las notificaciones del Anfitrion
func Mo_Find(iduser int, pagina int64, typeuser int) ([]models.Mo_NotificationShow, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := models.MongoCN.Database("restoner")
	col := db.Collection("notifications")

	var resultado []models.Mo_NotificationShow

	condicion := bson.M{
		"typeuser": typeuser,
		"iduser":   iduser,
	}

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "dateregistered", Value: -1}})
	opciones.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		return resultado, err
	}

	for cursor.Next(context.TODO()) {

		var registro models.Mo_NotificationShow
		err := cursor.Decode(&registro)
		if err != nil {
			return resultado, err
		}

		resultado = append(resultado, registro)
	}

	return resultado, nil

}
