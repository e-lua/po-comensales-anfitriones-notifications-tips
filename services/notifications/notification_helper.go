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
	Data      []models.Mo_NotificationShow `json:"data"`
}
