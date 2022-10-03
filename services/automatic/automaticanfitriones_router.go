package automatic

import (
	"encoding/json"
	"log"
	"net/http"
)

var AutomaticanfitrionesRouter_pg *automaticanfitrionesRouter_pg

type automaticanfitrionesRouter_pg struct {
}

func (ar *automaticanfitrionesRouter_pg) AddNotificationInsumo() {

	//OBTENEMOS LOS DATOS DE INSUMOS POR TERMINAR
	respuesta_toend, _ := http.Get("https://produccion-anfitrion.restoner-api.fun/v1/notify/insumo/toended")
	var get_respuesta_toend Response_Notify_Insumo
	error_decode_respuesta_toend := json.NewDecoder(respuesta_toend.Body).Decode(&get_respuesta_toend)
	if error_decode_respuesta_toend != nil {
		results := Response{Error: true, DataError: "Error al obtener la notificacion de los insumos por finalizar, detalles: " + error_decode_respuesta_toend.Error(), Data: ""}
		log.Println(500, results)
	}

	//OBTENEMOS LOS DATOS DE INSUMOS TERMINADOS
	respuesta_ended, _ := http.Get("https://produccion-anfitrion.restoner-api.fun/v1/notify/insumo/ended")
	var get_respuesta_ended Response_Notify_Insumo
	error_decode_respuesta_ended := json.NewDecoder(respuesta_ended.Body).Decode(&get_respuesta_ended)
	if error_decode_respuesta_ended != nil {
		results := Response{Error: true, DataError: "Error al obtener la notificacion de los insumos finalizados, detalles: " + error_decode_respuesta_toend.Error(), Data: ""}
		log.Println(500, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddNotificationInsumo_Service(get_respuesta_toend.Data, get_respuesta_ended.Data)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	log.Println(status, results)
}
