package models

type Product struct {
	Detail *Detail `json:"detail"`
	Rating *Rating `json:"rating"`
	Reviews []*Review `json:"reviews"`
}
