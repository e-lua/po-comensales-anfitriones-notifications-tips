package notifications

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"
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

	return nil
}
