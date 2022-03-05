package tips

import "github.com/Aphofisis/po-comensales-anfitriones-notifications-tips/models"

type Response struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      string `json:"data"`
}

type Response_Tips struct {
	Error     bool                  `json:"error"`
	DataError string                `json:"dataError"`
	Data      []*models.Mo_TipsShow `json:"data"`
}
