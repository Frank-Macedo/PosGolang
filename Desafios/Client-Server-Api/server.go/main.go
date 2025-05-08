package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
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
		w.Write([]byte("Resposta da API externa:\n"))
		w.Write([]byte(body))
	}
}

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
