package repository

import (
	"context"
	"errors"
	"moduloInicial/Desafios/Client-Server-Api/server.go/model"
	"time"

	"gorm.io/gorm"
)

type CotacaoRepository struct {
	db *gorm.DB
}

func NewCotacaoRepository(db *gorm.DB) *CotacaoRepository {
	return &CotacaoRepository{db: db}
}

func (r *CotacaoRepository) Save(cotacao model.CotacaoDB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	done := make(chan error, 1)

	go func() {
		err := r.db.WithContext(ctx).Create(&cotacao).Error
		done <- err
	}()

	select {
	case err := <-done:
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				return errors.New("timeout: salvar cotação demorou demais")
			}
			return err
		}
		return nil
	case <-ctx.Done():
		return errors.New("timeout: salvar cotação cancelada pelo contexto")
	}

}
