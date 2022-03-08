package notifications

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"
	"github.com/labstack/echo/v4"
)

var NotificationsRouter_pg *notificationsRouter_pg

type notificationsRouter_pg struct {
}

/*-------------------------------*/
func GetJWT(jwt string, service int, module int, epic int, endpoint int) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://a-registro-authenticacion.restoner-api.fun:5000/v1/trylogin?jwt=" + jwt + "&service=" + strconv.Itoa(service) + "&module=" + strconv.Itoa(module) + "&epic=" + strconv.Itoa(epic) + "&endpoint=" + strconv.Itoa(endpoint))
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0
	}
	return 200, false, "", get_respuesta.Data.IdBusiness
}

/*-------------------------------*/

func (nr *notificationsRouter_pg) AddNotification(c echo.Context) error {

	//Instanciamos una variable del modelo B_Name
	var notification models.Mo_Notifications

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&notification)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el nombre del negocio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if len(notification.Message) < 5 || notification.IDUser < 0 || notification.TypeUser < 0 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: ""}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddNotification_Service(notification)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (nr *notificationsRouter_pg) ShowNotification(c echo.Context) error {

	user_string := c.Request().URL.Query().Get("user")
	user_int, _ := strconv.Atoi(user_string)

	page_string := c.Request().URL.Query().Get("page")
	page_int, _ := strconv.ParseInt(page_string, 10, 64)

	type_string := c.Request().URL.Query().Get("typeuser")
	type_int, _ := strconv.Atoi(type_string)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := ShowNotification_Service(user_int, page_int, type_int)
	results := Response_Notifications{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (nr *notificationsRouter_pg) UpdateNotification(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"), 2, 2, 1, 10)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	type_string := c.Request().URL.Query().Get("typeuser")
	type_int, _ := strconv.Atoi(type_string)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateNotification_Service(data_idbusiness, type_int)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
