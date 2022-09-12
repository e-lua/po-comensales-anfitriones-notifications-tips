package models

type Mqtt_LegalIdentity struct {
	IdBusiness      int     `bson:"idbusiness" json:"idbusiness"`
	LegalIdentity   string  `bson:"legalidentity" json:"legalidentity"`
	TypeSuscription int     `bson:"typesuscription" json:"typesuscription"`
	IVA             float32 `bson:"iva" json:"iva"`
}
