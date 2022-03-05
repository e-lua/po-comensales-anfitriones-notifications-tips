package notifications

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
)

func Mo_Add(notification models.Mo_Notifications) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	db := models.MongoCN.Database("restoner")
	col := db.Collection("notifications")

	_, err := col.InsertOne(ctx, notification)
	if err != nil {
		return err
	}

	/*-------------------------------*/

	var serialize_notification models.Mo_NotificationShow

	serialize_notification.Message = notification.Message
	serialize_notification.IDUser = notification.IDUser

	go func() {
		//Comienza el proceso de MQTT
		ch, error_conection := models.MqttCN.Channel()
		if error_conection != nil {
			log.Error(error_conection)
		}

		bytes, error_serializar := serialize(serialize_notification)
		if error_serializar != nil {
			log.Error(error_serializar)
		}

		if notification.TypeUser == 1 {
			error_publish := ch.Publish("", "notification/anfitrion", false, false,
				amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "text/plain",
					Body:         bytes,
				})
			if error_publish != nil {
				log.Error(error_publish)
			}
		}

		if notification.TypeUser == 2 {
			error_publish := ch.Publish("", "notification/comensal", false, false,
				amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "text/plain",
					Body:         bytes,
				})
			if error_publish != nil {
				log.Error(error_publish)
			}
		}

		if notification.TypeUser == 3 {
			error_publish := ch.Publish("", "notification/colaborador", false, false,
				amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "text/plain",
					Body:         bytes,
				})
			if error_publish != nil {
				log.Error(error_publish)
			}
		}

	}()

	return nil
}

//SERIALIZADORA
func serialize(serialize_notification_input models.Mo_NotificationShow) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(serialize_notification_input)
	if err != nil {
		return b.Bytes(), err
	}
	return b.Bytes(), nil
}
