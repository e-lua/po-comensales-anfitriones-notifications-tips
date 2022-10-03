package notifications

import "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"

type Response struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      string `json:"data"`
}

type Response_Notifications struct {
	Error     bool                         `json:"error"`
	DataError string                       `json:"dataError"`
	Data      []models.Pg_NotificationShow `json:"data"`
}

type ResponseJWT struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      JWT    `json:"data"`
}

type JWT struct {
	IdBusiness int `json:"idBusiness"`
	IdWorker   int `json:"idWorker"`
	IdCountry  int `json:"country"`
	IdRol      int `json:"rol"`
}

type Income_IDDevice struct {
	Error     bool     `json:"error"`
	DataError string   `json:"dataError"`
	Data      []string `json:"data"`
}
