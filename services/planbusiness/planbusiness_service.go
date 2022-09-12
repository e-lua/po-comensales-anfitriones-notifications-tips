package planbusiness

import (

	//REPOSITORIES

	models "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"
	plan_business "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/repositories/plan_business"
)

func AddBusiness_Service(anfitriones_planbusiness_all []models.Mqtt_LegalIdentity) (int, bool, string, string) {

	for _, anfitriones_planbusiness := range anfitriones_planbusiness_all {

		business, _ := plan_business.Mo_Find(anfitriones_planbusiness.IdBusiness)

		if len(business.LegalIdentity) > 0 {
			//Si existe se actualizara
			error_add := plan_business.Mo_Update(anfitriones_planbusiness)
			if error_add != nil {
				return 500, true, "Error interno en el servidor al intentar actualizar el Anfitrion, detalles: " + error_add.Error(), ""
			}
		} else {
			//Si no existe se agrega
			error_add := plan_business.Mo_Add(anfitriones_planbusiness)
			if error_add != nil {
				return 500, true, "Error interno en el servidor al intentar registrar el Anfitrion, detalles: " + error_add.Error(), ""
			}
		}

	}

	return 201, false, "", "Agregado o actualizado correctamente"
}
