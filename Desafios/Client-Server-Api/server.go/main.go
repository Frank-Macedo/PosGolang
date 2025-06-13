package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"moduloInicial/Desafios/Client-Server-Api/server.go/db"
	"moduloInicial/Desafios/Client-Server-Api/server.go/model"
	"moduloInicial/Desafios/Client-Server-Api/server.go/repository"
	"moduloInicial/Desafios/Client-Server-Api/server.go/service"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/cotacao", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)

	conn := db.NewDB(db.DBConfig{
		Driver:     "sqlite",
		SQLitePath: "./meubanco.db",
	})

	repo := repository.NewCotacaoRepository(conn)

	defer cancel()

	log.Println("Request Iniciada")
	defer log.Println("Request finalizada")

	result := make(chan string)
	errChan := make(chan error)

	go func() {
		// Faz a requisição externa
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
		if err != nil {
			errChan <- fmt.Errorf("erro ao criar requisição: %w", err)
			return
		}

		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			errChan <- fmt.Errorf("erro na chamada externa: %w", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			errChan <- fmt.Errorf("erro ao ler resposta: %w", err)
			return
		}

		result <- string(body)
	}()

	select {
	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente")
		http.Error(w, "Request cancelada pelo cliente", http.StatusRequestTimeout)

	case <-time.After(20 * time.Second):
		log.Println("Tempo limite atingido")
		http.Error(w, "Tempo limite atingido", http.StatusGatewayTimeout)

	case err := <-errChan:
		log.Println("Erro interno:", err)
		http.Error(w, "Erro interno: "+err.Error(), http.StatusInternalServerError)

	case body := <-result:
		log.Println("Request processada com sucesso")

		var cotacao model.Cotacao
		if err := json.Unmarshal([]byte(body), &cotacao); err != nil {
			log.Println("Erro ao fazer unmarshal:", err)
		}

		svc := service.NewCotacaoService(repo)

		svc.SaveServiceData(cotacao)
		w.Write([]byte("Resposta da API externa:\n"))
		w.Write([]byte(body))
	}
}
