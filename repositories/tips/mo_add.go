package tips

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"
)

func Mo_Add(tip models.Mo_Tips) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	db := models.MongoCN.Database("restoner")
	col := db.Collection("tips")

	_, err := col.InsertOne(ctx, tip)
	if err != nil {
		return err
	}

	return nil
}
