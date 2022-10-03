package automatic

import "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"

type Response struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      string `json:"data"`
}

type Income_IDDevice struct {
	Error     bool     `json:"error"`
	DataError string   `json:"dataError"`
	Data      []string `json:"data"`
}

//NOTIFY INSUMO
type Response_Notify_Insumo struct {
	Error     bool                      `json:"error"`
	DataError string                    `json:"dataError"`
	Data      []models.Pg_Notifications `json:"data"`
}
