package tips

import (

	//REPOSITORIES
	models "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"
	tips_repository "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/repositories/tips"
)

func AddTip_Service(tip_input models.Mo_Tips) (int, bool, string, string) {

	//Obtenemos las categorias
	error_update := tips_repository.Mo_Add(tip_input)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar agregar el tip, detalles: " + error_update.Error(), ""
	}

	return 201, false, "", "Tip agregado correctamente"
}

func ShowTips_Service(idtype_int int) (int, bool, string, []*models.Mo_TipsShow) {

	//Obtenemos las categorias
	tips, error_update := tips_repository.Mo_Find(idtype_int)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar mostrar los tips, detalles: " + error_update.Error(), tips
	}

	return 201, false, "", tips
}
