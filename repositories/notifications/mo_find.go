package notifications

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Traeremos las notificaciones del Anfitrion
func Mo_Find(iduser int, pagina int64, typeuser int) ([]*models.Mo_NotificationShow, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := models.MongoCN.Database("restoner")
	col := db.Collection("notifications")

	/*Aca pude haber hecho un make, es decir, resultado:=make([]...)*/
	var resultado []*models.Mo_NotificationShow

	/*El ID del usuario que llegar por parámetro, tiene que ser igual el userid
	de la BD en mongo de Tweet*/

	condicion := bson.M{
		"typeuser": typeuser,
		"$and": []interface{}{
			bson.M{"iduser": iduser},
			bson.M{"iduser": 0},
		},
	}

	opciones := options.Find()
	/*Cuandos documentos como limite quiero*/
	opciones.SetLimit(20)
	/*Indicar como ira ordenado*/
	opciones.SetSort(bson.D{{Key: "dateregistered", Value: -1}})
	opciones.SetSkip((pagina - 1) * 20)

	/*Cursor es como una tabla de base de datos donde se van a grabar los resultados
	y podre ir recorriendo 1 a la vez*/
	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		return resultado, err
	}

	//contexto, en este caso, me crea un contexto vacio
	for cursor.Next(context.TODO()) {
		/*Aca trabajare con cada Tweet. El resultado lo grabará en registro*/
		var registro models.Mo_NotificationShow
		err := cursor.Decode(&registro)
		if err != nil {
			return resultado, err
		}
		/*Recordar que Append sirve para añadir un elemento a un slice*/
		resultado = append(resultado, &registro)
	}

	return resultado, nil

}
