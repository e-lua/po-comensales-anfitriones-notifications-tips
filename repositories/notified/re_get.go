package repositories

import (
	"strconv"

	models "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"
	"github.com/gomodule/redigo/redis"
)

func Re_Get_Notified(idbusiness int, code int) (int, error) {

	var idbusiness_int int

	reply, err := redis.String(models.RedisCN.Get().Do("GET", strconv.Itoa(idbusiness)+strconv.Itoa(code)))
	if err != nil {
		return idbusiness_int, err
	}

	idbusiness_int2, _ := strconv.Atoi(reply)

	if err != nil {
		return idbusiness_int2, err
	}

	return idbusiness_int2, nil
}
