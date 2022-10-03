package notifications

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"
	"github.com/google/uuid"
)

func Pg_Add(notification models.Pg_Notifications) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	//Nos conectamos con la BD
	db := models.Conectar_Pg_DB()

	query_order := `INSERT INTO Notification(id,title,message,dateregistered,iduser,multipleuser,priority,wasview,typeuser) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`
	if _, err := db.Exec(ctx, query_order, uuid.New().String(), notification.Title, notification.Message, time.Now(), notification.IDUser, notification.MultipleUser, notification.Priority, notification.WasView, notification.TypeUser); err != nil {
		return err
	}

	return nil
}
