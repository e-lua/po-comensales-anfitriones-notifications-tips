package planbusiness

import (
	"log"

	"github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"
)

var PlanbusinessRouter_pg *planbusinessRouter_pg

type planbusinessRouter_pg struct {
}

func (pr *planbusinessRouter_pg) AddBusiness(anfitriones_planbusiness_all []models.Mqtt_LegalIdentity) {

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddBusiness_Service(anfitriones_planbusiness_all)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	log.Println(status, results)
}
