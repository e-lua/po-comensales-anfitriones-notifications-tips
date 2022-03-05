package tips

import (
	"strconv"

	"github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"
	"github.com/labstack/echo/v4"
)

var TipsRouter_pg *tipsRouter_pg

type tipsRouter_pg struct {
}

func (tr *tipsRouter_pg) AddTip(c echo.Context) error {

	//Instanciamos una variable del modelo B_Name
	var tip models.Mo_Tips

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&tip)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el nombre del negocio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if len(tip.URLimg) < 5 || tip.Typetip < 0 || tip.TypeUser != 1 && tip.TypeUser != 2 && tip.TypeUser != 3 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: ""}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddTip_Service(tip)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (tr *tipsRouter_pg) ShowTips(c echo.Context) error {

	type_string := c.Request().URL.Query().Get("typeuser")
	type_int, _ := strconv.Atoi(type_string)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := ShowTips_Service(type_int)
	results := Response_Tips{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
