package repositories

import (
	"strconv"

	models "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"
)

func Re_Set_Notified(idbusiness int, code int, quantity int) error {

	_, err_do := models.RedisCN.Get().Do("SET", strconv.Itoa(idbusiness)+strconv.Itoa(code), strconv.Itoa(quantity), "EX", 3800)
	if err_do != nil {
		return err_do
	}

	return nil
}
