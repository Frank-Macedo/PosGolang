package model

import "gorm.io/gorm"

type Cotacao struct {
	Usdbrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

type CotacaoDB struct {
	gorm.Model
	Code       string
	Codein     string
	Name       string
	High       string
	Low        string
	VarBid     string
	PctChange  string
	Bid        string
	Ask        string
	Timestamp  string
	CreateDate string
}

func (c *Cotacao) ToCotacaoDB() CotacaoDB {
	return CotacaoDB{
		Code:       c.Usdbrl.Code,
		Codein:     c.Usdbrl.Codein,
		Name:       c.Usdbrl.Name,
		High:       c.Usdbrl.High,
		Low:        c.Usdbrl.Low,
		VarBid:     c.Usdbrl.VarBid,
		PctChange:  c.Usdbrl.PctChange,
		Bid:        c.Usdbrl.Bid,
		Ask:        c.Usdbrl.Ask,
		Timestamp:  c.Usdbrl.Timestamp,
		CreateDate: c.Usdbrl.CreateDate,
	}
}
