// service/cotacao_service.go

package service

import (
	"moduloInicial/Desafios/Client-Server-Api/server.go/model"
	"moduloInicial/Desafios/Client-Server-Api/server.go/repository"
)

type CotacaoService struct {
	repo *repository.CotacaoRepository
}

func NewCotacaoService(repo *repository.CotacaoRepository) *CotacaoService {
	return &CotacaoService{repo: repo}
}

func (s *CotacaoService) SaveServiceData(apiCotacao model.Cotacao) error {
	dbCotacao := model.CotacaoDB{
		Code:       apiCotacao.Usdbrl.Code,
		Codein:     apiCotacao.Usdbrl.Codein,
		Name:       apiCotacao.Usdbrl.Name,
		High:       apiCotacao.Usdbrl.High,
		Low:        apiCotacao.Usdbrl.Low,
		VarBid:     apiCotacao.Usdbrl.VarBid,
		PctChange:  apiCotacao.Usdbrl.PctChange,
		Bid:        apiCotacao.Usdbrl.Bid,
		Ask:        apiCotacao.Usdbrl.Ask,
		Timestamp:  apiCotacao.Usdbrl.Timestamp,
		CreateDate: apiCotacao.Usdbrl.CreateDate,
	}

	return s.repo.Save(dbCotacao)
}
